package database

import (
    "../types"

    "time"
    "log"
    "strings"
    "database/sql"
)

type QueryParams struct {
	Email string
	Week  string
}

//TODO REWORK
func Select(params QueryParams) []types.Idea {
	var ideas []types.Idea
    var query string

    query = "SELECT id, body, email, week, created_at FROM public.idea WHERE 1=1"
    
    if params.Email != "" {
        query += " AND email = $1 "
    }
    
    if params.Week != "" {
        query += " AND week = $2 "
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
        var week string

        log.Printf("Reading row %v", rows)

		err := rows.Scan(&id, &body, &email, &week, &created_at)
		if err != nil {
			log.Fatal(err)
        }
        
        log.Printf("Row value : (%d, %s, %s, %s, %v)", id, body, email, week, created_at)

        idea := types.Idea{
            ID: id, 
            Body: body,
            Email: email,
            Week: week,
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

func UpdateIdea(id int, body string) int64 {
	res, err := DBCon.Exec(`
        UPDATE public.idea
        SET body = $1
        WHERE id = $2
    `, body, id)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count > 0 {
		log.Print("idea updated in database")
	}

	return count
}
