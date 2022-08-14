package main

import (
	"context"
	"encoding/json"
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
	store := NewStore(conn)
	todoHandler := &todoHandler{store: store}
	http.Handle("/todos", todoHandler)

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// if _, err := conn.Exec(context.Background(), "INSERT INTO TODOS(TITLE) VALUES($1)", "Hello db"); err != nil {
	// 	// Handling error, if occur
	// 	fmt.Println("Unable to insert due to: ", err)
	// 	return
	// }
	// fmt.Println("Insertion Succesfull")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

type store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *store {
	return &store{conn: conn}
}

func (s *store) NewTask(title string) error {
	if _, err := s.conn.Exec(context.Background(), "INSERT INTO TODOS(TITLE) VALUES($1)", title); err != nil {
		return err
	}
	return nil
}

type todoHandler struct {
	store interface {
		NewTask(title string) error
	}
}

type TodoNewTask struct {
	Title string `json:"title"`
}

func (h *todoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t TodoNewTask
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.store.NewTask(t.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(map[string]string{"message": " success"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
