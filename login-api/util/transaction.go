package util

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		SentPanicIfError(errRollBack)
		panic(err)
	} else {
		errCommit := tx.Commit()
		SentPanicIfError(errCommit)
	}
}
