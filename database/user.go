package database

import (
    "../types"

    "log"
)

func InsertUser(user *types.User) error {
	var id int
	err := DBCon.QueryRow(`
		INSERT INTO public."user"(firstname, lastname, email)
		VALUES ($1, $2, $3)
		RETURNING id
    `, user.Firstname, user.Lastname, user.Email).Scan(&id)
    
	if err != nil {
        log.Fatal(err)
		return err
	}
    
    log.Print("user inserted in database")
    
    user.ID = id
	return nil
}

func DeleteUser(email string) int64 {
    res, err := DBCon.Exec(`
        DELETE FROM public."user"
        WHERE email = $1;
    `, email)

    if err != nil {
        panic(err)
    }
    
    count, err := res.RowsAffected()
    if err != nil {
        panic(err)
    }

    if count > 0 {
        log.Print("user deleted in database")
    }

    return count
}

func UpdateUser(id int, firstname string, lastname string) int64 {
    res, err := DBCon.Exec(`
        UPDATE public."user"
        SET firstname = $1, lastname = $2
        WHERE id = $3
    `, firstname, lastname, id)

     if err != nil {
        panic(err)
    }
    
    count, err := res.RowsAffected()
    if err != nil {
        panic(err)
    }

    if count > 0 {
        log.Print("user updated in database")
    }

    return count
}
