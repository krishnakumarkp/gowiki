package controller

import (
	"gowiki2/domain"
	"net/http"
	"text/template"
)

type WikiController struct {
	Store domain.WikiStore
}

// func (w WikiController) Add(p *domain.Page) error {
// 	err := w.Store.Save(p)
// 	return err
// }

func (wc WikiController) View(w http.ResponseWriter, r *http.Request, title string) {
	p, err := wc.Store.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func (wc WikiController) Edit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := wc.Store.LoadPage(title)
	if err != nil {
		p = &domain.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func (wc WikiController) Save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &domain.Page{Title: title, Body: []byte(body)}
	err := wc.Store.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *domain.Page) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
