package database

import (
    "../types"

    "time"
    "log"
    "strings"
    "database/sql"
    "errors"
    "encoding/json"
)

type IdeaQueryParams struct {
    Body  string
	Email string
	Week  string
}

//TODO REWORK
func Select(params IdeaQueryParams) []types.Idea {
	var ideas []types.Idea
    var query string

    query = "SELECT id, body, email, week, jsonb_array_length(votes) votes, votes, created_at FROM public.idea WHERE 1=1"
    
    if params.Email != "" {
        query += " AND email = $1 "

        if params.Week != "" {
            query += " AND week = $2 "
        }
    } else {
        if params.Week != "" {
            query += " AND week = $1 "
        }
    }
    
    log.Printf("SQL : %s", query)
	statement, err := DBCon.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

    var rows *sql.Rows

    if params.Email != "" && params.Week != "" {
        rows, err = statement.Query(params.Email, params.Week)
    } else if (params.Email != "") {
        rows, err = statement.Query(params.Email)
    } else if (params.Week != "") {
        rows, err = statement.Query(params.Week)
    } else {        
        rows, err = statement.Query()
    }

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
        var body string
        var email string 
        var created_at time.Time
        var votes int
        var voters string
        var week string

        log.Printf("Reading row %v", rows)

		err := rows.Scan(&id, &body, &email, &week, &votes, &voters, &created_at)
		if err != nil {
			log.Fatal(err)
        }
        
        log.Printf("Row value : (%d, %s, %s, %s,%d, %v)", id, body, email, week, votes, voters, created_at)

        log.Printf("voters %v", voters)
        var votersArray []string
        err = json.Unmarshal([]byte(voters), &votersArray)
        if err != nil {
            log.Printf("Error voters json %v", err)
        }

        log.Printf("voters %v", votersArray)

        idea := types.Idea{
            ID: id, 
            Body: body,
            Email: email,
            Week: week,
            Votes: votes,
            Voters: votersArray,
            CreatedAt: created_at,
        }

		ideas = append(ideas, idea)
    }
    
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
    }
    
    log.Printf("Found %v rows", len(ideas))

	return ideas
}

func InsertIdea(idea *types.Idea) error {
	var id int
	err := DBCon.QueryRow(`
		INSERT INTO public.idea(body, email, week)
		VALUES ($1, $2, $3)
		RETURNING id
    `, idea.Body, idea.Email, idea.Week).Scan(&id)

	if err != nil {
        if strings.Contains(err.Error(), "idea_user_email_week_key") {
            log.Printf("Error, user already post an idea this week")
        } else {
            log.Fatal(err)
        }
		return err
	}

	log.Print("idea inserted in database")

	idea.ID = id
	return nil
}

func UpdateIdea(params IdeaQueryParams) (types.Idea, error) {

    sql := `UPDATE public.idea
    SET body = $1
    WHERE email = $2
    AND week = $3`

	res, err := DBCon.Exec(sql, params.Body, params.Email, params.Week)

    log.Printf("SQL %v (%s %s %s)", sql, params.Body, params.Email, params.Week)
    
    if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

    var idea types.Idea 
    
    if count > 1 {
        panic("should not update more than one idea")
    }

	if count > 0 {
        log.Printf("%d idea updated in database", count)
        
        idea = Select(params)[0]
        log.Printf("%v", idea)

        return idea, nil
    }
    

    error := errors.New("Nothing updated")

	return idea, error 
}
