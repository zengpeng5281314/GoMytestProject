package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	//相对路径
	filePrefix, _ := filepath.Abs("./templates")
	tmpl := template.Must(template.ParseFiles(filePrefix + "/layout.html"))
	//绝对路径
	//tmpl := template.Must(template.ParseFiles("/Users/edy/GolandProjects/mytestProject/templates/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todo list", Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
			},
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
