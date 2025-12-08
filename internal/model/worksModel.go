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
	Sold        bool         `json:"sold"`
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
				Href:        "/works/turnables/die-zwei-jacos",
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
				Href:        "/works/turnables/feuermuehle",
				Description: "A beautiful turnable item",
			},
			{
				ID:       "jacinta",
				Title:    "Jacinta",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/jacinta/jacinta-01.png", Alt: "Jacinta 01"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-02.png", Alt: "Jacinta 02"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-03.png", Alt: "Jacinta 03"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-04.png", Alt: "Jacinta 04"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-05.png", Alt: "Jacinta 05"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-06.png", Alt: "Jacinta 06"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-07.png", Alt: "Jacinta 07"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-08.png", Alt: "Jacinta 08"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-09.png", Alt: "Jacinta 09"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-10.png", Alt: "Jacinta 10"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-11.png", Alt: "Jacinta 11"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-12.png", Alt: "Jacinta 12"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-13.png", Alt: "Jacinta 13"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-14.png", Alt: "Jacinta 14"},
					{Src: "/static/images/works/turnables/jacinta/jacinta-15.png", Alt: "Jacinta 15"},
				},
				Href:        "/works/turnables/jacinta",
				Description: "\"Jacinta\" rotation loop area    0,76 - 1995",
			},
			{
				ID:          "b001",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b001.jpg", Alt: ""}},
				Href:        "/works/pictures/b001",
				Description: "",
			},
			{
				ID:          "b002",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b002.jpg", Alt: ""}},
				Href:        "/works/pictures/b002",
				Description: "",
			},
			{
				ID:          "b003",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b003.jpg", Alt: ""}},
				Href:        "/works/pictures/b003",
				Description: "",
			},
			{
				ID:          "b004",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b004.jpg", Alt: ""}},
				Href:        "/works/pictures/b004",
				Description: "",
			},

			{
				ID:          "b005",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b005.jpg", Alt: ""}},
				Href:        "/works/pictures/b005",
				Description: "",
			},
			{
				ID:          "b006",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b006.jpg", Alt: ""}},
				Href:        "/works/pictures/b006",
				Description: "",
			},
			{
				ID:          "b007",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b007.jpg", Alt: ""}},
				Href:        "/works/pictures/b007",
				Description: "",
			},
			{
				ID:          "b008",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b008.jpg", Alt: ""}},
				Href:        "/works/pictures/b008",
				Description: "",
			},
			{
				ID:          "b009",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b009.jpg", Alt: ""}},
				Href:        "/works/pictures/b009",
				Description: "",
			},
			{
				ID:          "b010",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b010.jpg", Alt: ""}},
				Href:        "/works/pictures/b010",
				Description: "",
			},
			{
				ID:          "b011",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b011.jpg", Alt: ""}},
				Href:        "/works/pictures/b011",
				Description: "",
			},

			{
				ID:          "b012",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b012.jpg", Alt: ""}},
				Href:        "/works/pictures/b012",
				Description: "",
			},

			{
				ID:          "b013",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b013.jpg", Alt: ""}},
				Href:        "/works/pictures/b013",
				Description: "",
			},
			{
				ID:          "b014",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b014.jpg", Alt: ""}},
				Href:        "/works/pictures/b014",
				Description: "",
			},
			{
				ID:          "b015",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b015.jpg", Alt: ""}},
				Href:        "/works/pictures/b015",
				Description: "",
			},
			{
				ID:          "b016",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b016.jpg", Alt: ""}},
				Href:        "/works/pictures/b016",
				Description: "",
			},
			{
				ID:          "b017",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b017.jpg", Alt: ""}},
				Href:        "/works/pictures/b017",
				Description: "",
			},
			{
				ID:          "b018",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b018.jpg", Alt: ""}},
				Href:        "/works/pictures/b018",
				Description: "",
			},
			{
				ID:          "b019",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b019.jpg", Alt: ""}},
				Href:        "/works/pictures/b019",
				Description: "",
			},
			{
				ID:          "b020",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b020.jpg", Alt: ""}},
				Href:        "/works/pictures/b020",
				Description: "",
			},
			{
				ID:          "b021",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b021.jpg", Alt: ""}},
				Href:        "/works/pictures/b021",
				Description: "",
			},
			// {
			// 	ID:          "b022",
			// 	Title:       "",
			// 	Category:    "pictures",
			// 	Images:      []WorksImage{{Src: "/static/images/works/pictures/b022.jpg", Alt: ""}},
			// 	Href:        "/works/pictures/b022",
			// 	Description: "",
			// },
			{
				ID:          "b023",
				Title:       "",
				Category:    "pictures",
				Images:      []WorksImage{{Src: "/static/images/works/pictures/b023.jpg", Alt: ""}},
				Href:        "/works/pictures/b023",
				Description: "",
			},
			{
				ID:       "NoName 01",
				Title:    "NoName 01",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/1/IMG_5353.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5354.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5355.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5356.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5357.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5358.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5359.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5360.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5361.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5362.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5363.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5364.jpg", Alt: "NoName 01"},
					{Src: "/static/images/works/turnables/1/IMG_5365.jpg", Alt: "NoName 01"},
				},
				Href:        "/works/turnables/noname01",
				Description: "\"NoName 01\" - 1990\nRotation loop area 1,70m",
				Sold:        true,
			},
			{
				ID:       "NoName 02",
				Title:    "NoName 02",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/2/IMG_5366.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5367.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5368.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5369.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5370.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5371.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5372.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5373.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5374.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5375.jpg", Alt: "NoName 02"},
					{Src: "/static/images/works/turnables/2/IMG_5376.jpg", Alt: "NoName 02"},
				},
				Href:        "/works/turnables/noname02",
				Description: "\"NoName 02\" - 1990\nRotation loop area 2,60m\nFrame 1,54m x 1,57m",
				Sold:        false,
			}},
	}
}
