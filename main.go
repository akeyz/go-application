package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zserge/lorca"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func httpServe() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8888", nil)
}

func main() {
	os.Remove("./screen.db")

	db, err := sql.Open("sqlite3", "./screen.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		create table foo (id integer not null primary key, name text);
		delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	go httpServe()

	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}

	bounds := new(lorca.Bounds)
	bounds.WindowState = "fullscreen"

	ui.SetBounds(*bounds)
	ui.Load("http://localhost:8888")

	defer ui.Close()

	// Wait until UI window is closed
	<-ui.Done()
}
