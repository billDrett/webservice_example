package main

import "io/ioutil"

const filesLocation = "data/"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filesLocation+filename, p.Body, 0600)
}

func loadPage(title string) (Page, error) {
	filename := title + ".txt"

	body, errorCode := ioutil.ReadFile(filesLocation + filename)
	if errorCode != nil {
		return Page{}, errorCode
	}
	return Page{title, body}, nil
}
