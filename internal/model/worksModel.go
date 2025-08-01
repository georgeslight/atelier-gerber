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
	Category    string       `json:"category"`
	Images      []WorksImage `json:"image"`
	Href        string       `json:"href"`
	Description string       `json:"description"`
}

func GetWorksData() WorksData {
	return WorksData{
		WorksData: []Work{
			{
				ID:       "die-zwei-jacobs",
				Category: "turnables",
				Images: []WorksImage{
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-01.png", Alt: "Die Zwei Jacobs 01"},
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-02.png", Alt: "Die Zwei Jacobs 02"},
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-03.png", Alt: "Die Zwei Jacobs 03"},
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-04.png", Alt: "Die Zwei Jacobs 04"},
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-05.png", Alt: "Die Zwei Jacobs 05"},
					{Src: "/static/images/works/turnables/die-zwei-jacobs/die-zwei-jacobs-06.png", Alt: "Die Zwei Jacobs 06"},
				},
				Href:        "/works/die-zwei-jacobs",
				Description: "A beautiful turnable item",
			},
		},
	}
}
