package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/f3ar87/go-sample-webapp/viewmodel"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates[requestedFile+".html"]
		var context interface{}
		switch requestedFile {
		case "shop":
			context = viewmodel.NewShop()
		case "home":
			context = viewmodel.NewHome()
		case "stand_locator":
			context = viewmodel.NewStandLocator()
		default:
			context = viewmodel.NewBase()
		}
		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(404)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmp1 := template.Must(layout.Clone())
		_, err = tmp1.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmp1
	}
	return result
}
