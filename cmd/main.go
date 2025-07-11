package main

import (
	"fmt"
	//"github.com/a-h/templ"
	"kristin-gerber/views/layouts"
	"kristin-gerber/views/pages"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//component := layouts.BaseLayout()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/exhibitions", exhibitionsHandler)
	http.HandleFunc("/works", worksHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homePage := pages.HomePage()
	html := layouts.BaseLayout(homePage)
	html.Render(r.Context(), w)
}

func exhibitionsHandler(w http.ResponseWriter, r *http.Request) {
	exhibitionsPage := pages.ExhibitionPage()
	html := layouts.BaseLayout(exhibitionsPage)
	html.Render(r.Context(), w)
}

func worksHandler(w http.ResponseWriter, r *http.Request) {
	worksPage := pages.WorksPage()
	html := layouts.BaseLayout(worksPage)
	html.Render(r.Context(), w)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutPage := pages.AboutPage()
	html := layouts.BaseLayout(aboutPage)
	html.Render(r.Context(), w)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contactPage := pages.ContactPage()
	html := layouts.BaseLayout(contactPage)
	html.Render(r.Context(), w)
}
