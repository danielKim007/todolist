package main

import (
	"log"
	"https://github.com/danielKim007/todolist/server"
)

// docker run --name tododb -p 5432:5432  -e POSTGRES_PASSWORD=1234Qwer -d postgres
//DATABASE_URL=postgres://my_app:my_password@localhost/my_database
func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8080",
		DatabaseURL: "postgresql://postgresql:1234Qwer@localhost:5432/todolist",
	}); err != nil {
		log.Fatalln(err)
	}
}
