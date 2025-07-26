package model

type ExhibitionImage struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type ExhibitionDetails struct {
	Images      []ExhibitionImage `json:"images"`
	Description string            `json:"description"`
}

type Exhibition struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Location  string            `json:"location"`
	Date      string            `json:"date"`
	Href      string            `json:"href"`
	BaseImage ExhibitionImage   `json:"baseImage"`
	Details   ExhibitionDetails `json:"details"`
}

type ExhibitionsData struct {
	Individual []Exhibition `json:"individual"`
	Group      []Exhibition `json:"group"`
}

func GetExhibitionsData() ExhibitionsData {
	return ExhibitionsData{
		Individual: []Exhibition{
			{
				ID:       "galerie-metamorphose",
				Title:    "Kristin Gerber",
				Location: "Galerie Métamorphose, Paris",
				Date:     "1963",
				Href:     "/exhibitions/galerie-metamorphose",
				BaseImage: ExhibitionImage{
					Src: "/static/images/b002.jpg",
					Alt: "Galerie Métamorphose",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/detail1.jpg",
							Alt: "Detail image 1",
						},
						{
							Src: "/static/images/detail2.jpg",
							Alt: "Detail image 2",
						},
					},
					Description: "Test Description",
				},
			},
			{
				ID:       "international-studio-program",
				Title:    "International Studio Program",
				Location: "MoMA PS1, Long Island City, New York",
				Date:     "1978 - 1979",
				Href:     "/exhibitions/international-studio-program",
				BaseImage: ExhibitionImage{
					Src: "/static/images/b002.jpg",
					Alt: "MoMA PS1",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/moma1.jpg",
							Alt: "MoMA detail 1",
						},
					},
					Description: "International Studio Program at MoMA PS1",
				},
			},
		},
		Group: []Exhibition{
			{
				ID:       "group-exhibition-example",
				Title:    "Group Exhibition Example",
				Location: "Contemporary Art Museum",
				Date:     "2023",
				Href:     "/exhibitions/group-exhibition-example",
				BaseImage: ExhibitionImage{
					Src: "/static/images/b002.jpg",
					Alt: "Group exhibition",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/group_detail1.jpg",
							Alt: "Group detail 1",
						},
					},
					Description: "A contemporary group exhibition",
				},
			},
		},
	}
}
