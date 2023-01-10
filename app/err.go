package app

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func log_write(data string) {
	var file, err = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	logger := *log.New(file, "", 0)
	defer file.Close()

	logger.Output(1, data)
}

func ErrorJson(err error, code uint32) {
	if err != nil {
		log_write(fmt.Sprintf("json error %s %v\n", err.Error(), code))
		log.Printf("JSON Data Conversion error - %v\n", code)
	}
}

func ErrorTx(err error, tx *sqlx.Tx, debug ...interface{}) {
	if err != nil {
		//log.Println(err)
		//log.Printf("%v\n", debug)
		log_write(fmt.Sprintf("tx error %s\n", err.Error()))
		tx.Rollback()
		return
	}
}

func ErrorFatal(err error) {
	if err != nil {
		log_write(fmt.Sprintf("fatal error %s\n", err.Error()))
		log.Fatal(err)
	}
}
