package app

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ErrorJson(err error, code uint32) {
	if err != nil {
		log.Printf("JSON Data Conversion error - %v\n", code)
	}
}

func ErrorTx(err error, tx *sqlx.Tx, debug ...interface{}) {
	if err != nil {
		//log.Println(err)
		//log.Printf("%v\n", debug)
		tx.Rollback()
		return
	}
}

func ErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
