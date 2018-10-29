package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type Asset struct {
	Jquery template.JS
	Bootst template.JS
	Custom template.JS
	Css    template.CSS
}

type TodoPageData struct {
	Empty     bool
	Artifact  string
	Region    string
	PageTitle string
	Assets    Asset
	Todos     []Todo
}

var (
	data = TodoPageData{
		Empty:     true,
		Artifact:  "samplelambda",
		Region:    "use1",
		PageTitle: "My TODO list",
		Assets: Asset{
			Jquery: template.JS(Jquery),
			Bootst: template.JS(Bootst),
			Custom: template.JS(Custom),
			Css:    template.CSS(Css),
		},
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl = template.Must(template.ParseFiles("layout.html"))
)

func imports(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {

	if data.Todos[2].Done {
		data.Todos[2].Done = false
	} else {
		data.Todos[2].Done = true
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	http.HandleFunc("/imports/", imports)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
