package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var (
	data = TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl = template.Must(template.ParseFiles("layout.html"))
)

func handler(w http.ResponseWriter, r *http.Request){
	
	if data.Todos[2].Done {
		data.Todos[2].Done = false
	} else {
		data.Todos[2].Done = true
	}
	
	fmt.Println(data.Todos[2].Done)
	
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}