package database

import (
    "../types"

    "log"
)

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
