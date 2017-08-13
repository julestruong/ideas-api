package database

import (
    "../types"

    "log"
)

func InsertUser(user *types.User) error {
	var id int
	err := DBCon.QueryRow(`
		INSERT INTO public."user"(id, firstname, lastname)
		VALUES ($1, $2, $3)
		RETURNING id
    `, user.ID, user.Firstname, user.Lastname).Scan(&id)
    
	if err != nil {
        log.Fatal(err);
		return err
	}
    
    log.Print("user inserted in database");
    
    user.ID = id
	return nil
}
