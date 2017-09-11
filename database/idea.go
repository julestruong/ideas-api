package database

import (
	"../types"

	"log"
)

type Params struct {
	Email string
	Id    int
}

func Select(params Params) []types.Idea {
	var ideas []types.Idea

	statement, err := DBCon.Prepare(`
    SELECT id, body, email, created_at FROM public.idea
    WHERE 1 = 1
    AND email = ?
    AND id = ? 
    `)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	rows, err := statement.Query(params.Email, params.Id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var body string

		err := rows.Scan(&id, &body)
		if err != nil {
			log.Fatal(err)
		}

		ideas = append(ideas, types.Idea{ID: id, Body: body})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ideas
}

func InsertIdea(idea *types.Idea) error {
	var id int
	err := DBCon.QueryRow(`
		INSERT INTO public.idea(body, email)
		VALUES ($1, $2)
		RETURNING id
    `, idea.Body, idea.Email).Scan(&id)

	if err != nil {
		log.Fatal(err)
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
