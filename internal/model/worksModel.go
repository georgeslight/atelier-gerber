package model

type WorksImage struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type WorksData struct {
	WorksData []Work `json:"worksData"`
}

type Work struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Category    string       `json:"category"`
	Images      []WorksImage `json:"image"`
	Href        string       `json:"href"`
	Description string       `json:"description"`
}

func GetWorksData() WorksData {
	return WorksData{
		WorksData: []Work{
			{
				ID:       "die-zwei-jacos",
				Title:    "Die zwei Jacos",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-01.png", Alt: "Die zwei Jacos 01"},
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-02.png", Alt: "Die zwei Jacos 02"},
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-03.png", Alt: "Die zwei Jacos 03"},
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-04.png", Alt: "Die zwei Jacos 04"},
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-05.png", Alt: "Die zwei Jacos 05"},
					{Src: "/static/images/works/turnables/die-zwei-jacos/die-zwei-jacos-06.png", Alt: "Die zwei Jacos 06"},
				},
				Href:        "/works/die-zwei-jacos",
				Description: "A beautiful turnable item",
			},
			{
				ID:       "feuermuehle",
				Title:    "Feuermühle",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-01.png", Alt: "Feuermühle 01"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-02.png", Alt: "Feuermühle 02"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-03.png", Alt: "Feuermühle 03"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-04.png", Alt: "Feuermühle 04"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-05.png", Alt: "Feuermühle 05"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-06.png", Alt: "Feuermühle 06"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-07.png", Alt: "Feuermühle 07"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-08.png", Alt: "Feuermühle 08"},
					{Src: "/static/images/works/turnables/feuermuehle/feuermuehle-nb-09.png", Alt: "Feuermühle 09"},
				},
				Href:        "/works/feuermuehle",
				Description: "A beautiful turnable item",
			},
		},
	}
}
