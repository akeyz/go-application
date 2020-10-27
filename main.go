package main

import (
	"html/template"
	"log"
	"net/http"

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
