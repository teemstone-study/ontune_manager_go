package app

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ErrorJson(err error) {
	if err != nil {
		log.Println("JSON Data Conversion error")
	}
}

func ErrorTx(err error, tx *sqlx.Tx) {
	if err != nil {
		//log.Println(err)
		tx.Rollback()
		return
	}
}

func ErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
