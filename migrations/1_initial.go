package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-pg/migrations/v8"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
fmt.Println(path) 
	file, err := os.Open(path+"/migrations/up.sql")
	if err != nil {
		panic(err)
	}

	byte, err := ioutil.ReadAll(file)
	if err  != nil {
		panic(err)
	}




	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table my_table...")
		_, err := db.Exec(string(byte))
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table my_table...")
		_, err := db.Exec(`DROP TABLE blogs`)
		return err
	})
}