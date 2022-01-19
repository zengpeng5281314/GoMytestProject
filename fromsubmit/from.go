package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Content struct {
	email   string
	subject string
	message string
}

func main() {
	filePrefix, _ := filepath.Abs("./fromsubmit")
	tmpl := template.Must(template.ParseFiles(filePrefix + "/submitfrom.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
			return
		}

		content := Content{
			email:   request.FormValue("email"),
			subject: request.FormValue("subject"),
			message: request.FormValue("message"),
		}
		//do something with content
		_ = content
		log.Println(content)
		tmpl.Execute(writer, struct {
			Success bool
		}{true})
	})
	http.ListenAndServe(":80", nil)

}
