package main

import (
	"fmt"
	"strings"

	//"github.com/a-h/templ"
	"kristin-gerber/internal/model"
	"kristin-gerber/views/layouts"
	"kristin-gerber/views/pages"
	"kristin-gerber/views/pages/exhibitions"
	"kristin-gerber/views/pages/works"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//component := layouts.BaseLayout()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/exhibitions", exhibitionsHandler)
	http.HandleFunc("/exhibitions/{id}", exhibitionDetailHandler)
	http.HandleFunc("/works", worksHandler)
	http.HandleFunc("/works/{category}/{id}", workDetailHandler)
	// http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homePage := pages.HomePage(model.GetHomeData())
	html := layouts.BaseLayout(homePage)
	html.Render(r.Context(), w)
}

func exhibitionsHandler(w http.ResponseWriter, r *http.Request) {
	exhibitionsPage := pages.ExhibitionPage(model.GetExhibitionsData())
	html := layouts.BaseLayout(exhibitionsPage)
	html.Render(r.Context(), w)
}

func exhibitionDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/exhibitions/")

	// Check for invalid ID first
	if id == "" || id == r.URL.Path {
		http.NotFound(w, r)
		return
	}

	exhibition := findExhibitionByID(id)
	if exhibition == nil {
		http.NotFound(w, r)
		return
	}

	// Render the exhibition detail page
	exhibitionDetailPage := exhibitions.ExhibitionDetailPage(*exhibition)
	html := layouts.BaseLayout(exhibitionDetailPage)
	html.Render(r.Context(), w)
}

func worksHandler(w http.ResponseWriter, r *http.Request) {
	worksPage := pages.WorksPage(model.GetWorksData())
	html := layouts.BaseLayout(worksPage)
	html.Render(r.Context(), w)
}

func workDetailHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/works/")

	// Check for invalid ID first
	if path == "" || path == r.URL.Path {
		http.NotFound(w, r)
		return
	}

	// Split the path to get category and id
	parts := strings.Split(path, "/")
	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}

	category := parts[0]
	id := parts[1]

	// Check for empty category or id
	if category == "" || id == "" {
		http.NotFound(w, r)
		return
	}

	work := findWorkByID(id)
	if work == nil {
		http.NotFound(w, r)
		return
	}

	// Render the exhibition detail page
	workDetailPage := works.WorksDetailPage(*work)
	html := layouts.BaseLayout(workDetailPage)
	html.Render(r.Context(), w)
}

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	aboutPage := pages.AboutPage()
// 	html := layouts.BaseLayout(aboutPage)
// 	html.Render(r.Context(), w)
// }

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contactPage := pages.ContactPage()
	html := layouts.BaseLayout(contactPage)
	html.Render(r.Context(), w)
}

func findExhibitionByID(id string) *model.Exhibition {
	data := model.GetExhibitionsData()

	for _, exhibition := range data.Individual {
		if exhibition.ID == id {
			return &exhibition
		}
	}
	for _, exhibition := range data.Group {
		if exhibition.ID == id {
			return &exhibition
		}
	}
	return nil
}

func findWorkByID(id string) *model.Work {
	data := model.GetWorksData()

	for _, work := range data.WorksData {
		if work.ID == id {
			return &work
		}
	}
	return nil
}
