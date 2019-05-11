package main

import (
	"html/template"
	"net/http"
	"regexp"
)

const templLocation = "tmpl/"

var templates = template.Must(template.ParseFiles(templLocation+"edit.html", templLocation+"view.html"))

func renderHTMLTemplate(w http.ResponseWriter, htmlPageName string, page Page) {
	errorCode := templates.ExecuteTemplate(w, htmlPageName+".html", &page)
	if errorCode != nil {
		http.Error(w, errorCode.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validatePath := regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")
		substrings := validatePath.FindStringSubmatch(r.URL.Path)
		if substrings == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, substrings[2])
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, errorCode := loadPage(title)

	if errorCode != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderHTMLTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, _ := loadPage(title)

	p.Title = title
	renderHTMLTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := Page{title, []byte(body)}
	errorCode := page.save()
	if errorCode != nil {
		http.Error(w, errorCode.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/frontPage", http.StatusFound)
}
