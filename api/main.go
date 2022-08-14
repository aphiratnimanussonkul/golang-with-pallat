package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	//graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{Addr: ":8080"}
	go func() {
		log.Fatal((srv.ListenAndServe()))
	}()

	<-ctx.Done()
	stop()
	timeooutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(timeooutCtx); err != nil {
		log.Println(err)
	}
	fmt.Println("Server stopped")
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

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(map[string]string{"message": "success"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
