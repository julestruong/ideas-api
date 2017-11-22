package database

import (
	"log"
)

type VoteQueryParams struct {
	Id    int
	Email string
}

func Vote(params VoteQueryParams) (bool, error) {
	sql := `UPDATE public.idea
    SET votes = votes || '["` + params.Email + `"]'::jsonb
    WHERE id = $1`

	res, err := DBCon.Exec(sql, params.Id)

	log.Printf("%s voting for idea %d ...", params.Email, params.Id)

	if err != nil {
		panic(err)
	}

	var count int64
	count, err = res.RowsAffected()

	if err != nil {
		panic(err)
	}

	log.Printf("oui %d", count)

	return true, nil
}
