package main

import (
	"database/sql"
	"gobooks/internal/cli"
	service "gobooks/internal/services"
	"gobooks/internal/web"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:root@tcp(localhost:3306)/gobook"
	// db, err := sql.Open("sqlite3", "./books.db")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	BookHandlers := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate"){
		bookCLI := cli.NewBookCLI(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", BookHandlers.GetBooks)
	router.HandleFunc("POST /books", BookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", BookHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", BookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", BookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}