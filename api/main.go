package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/myapp")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// if _, err := conn.Exec(context.Background(), "INSERT INTO TODOS(TITLE) VALUES($1)", "Hello db"); err != nil {
	// 	// Handling error, if occur
	// 	fmt.Println("Unable to insert due to: ", err)
	// 	return
	// }
	// fmt.Println("Insertion Succesfull")
}
