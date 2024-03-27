package main

import (
	"github.com/Cerebrovinny/login-app/migration"
	"log"
)

func main() {
	err := migration.CreateAdminUser()
	if err != nil {
		log.Fatal(err)
	}
}
