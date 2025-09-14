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
				ID:       "international-studio-program",
				Title:    "International Studio Program (Winter 1978-1979)",
				Location: "MoMA PS1, Long Island City, New York",
				Date:     "1978 - 1979",
				Href:     "/exhibitions/international-studio-program",
				BaseImage: ExhibitionImage{
					Src: "/static/images/exhibitions/moma/moma-01.jpg",
					Alt: "MoMA PS1",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/exhibitions/moma/moma-01.jpg",
							Alt: "MoMA detail 1",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-02.jpg",
							Alt: "MoMA detail 2",
						},

						{
							Src: "/static/images/exhibitions/moma/moma-03.jpg",
							Alt: "MoMA detail 3",
						},

						{
							Src: "/static/images/exhibitions/moma/moma-04.jpg",
							Alt: "MoMA detail 4",
						},
					},
					Description: "December 3, 1978 - Januray 1979. MoMA PS1.\nInternational Studio Program at The Museum of Modern Art, New Tork",
				},
			},
			// {
			// 	ID:       "galerie-metamorphose",
			// 	Title:    "Kristin Gerber",
			// 	Location: "Galerie Métamorphose, Paris",
			// 	Date:     "1963",
			// 	Href:     "/exhibitions/galerie-metamorphose",
			// 	BaseImage: ExhibitionImage{
			// 		Src: "",
			// 		Alt: "Galerie Métamorphose",
			// 	},
			// 	Details: ExhibitionDetails{
			// 		Images: []ExhibitionImage{
			// 			{
			// 				Src: "/static/images/exhibitions/galerie-metamorphose/galerie-metamorphose-01.png",
			// 				Alt: "Galerie Métamorphose 01",
			// 			},
			// 			{
			// 				Src: "/static/images/exhibitions/galerie-metamorphose/galerie-metamorphose-02.png",
			// 				Alt: "Galerie Métamorphose 02",
			// 			},
			// 			{
			// 				Src: "/static/images/exhibitions/galerie-metamorphose/galerie-metamorphose-03.png",
			// 				Alt: "Galerie Métamorphose 03",
			// 			},
			// 		},
			// 		Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.\n\nDuis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.\n\nUt wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.",
			// 	},
			// },
		},
		Group: []Exhibition{
			{
				ID:       "group-exhibition-example",
				Title:    "Group Exhibition Example",
				Location: "Test Art Museum",
				Date:     "2023",
				Href:     "/exhibitions/group-exhibition-example",
				BaseImage: ExhibitionImage{
					Src: "/static/images/exhibitions/moma/moma-01.jpg",
					Alt: "Test Image 01",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/exhibitions/moma/moma-01.jpg",
							Alt: "Test Image 01",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-02.jpg",
							Alt: "Test Image 02",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-03.jpg",
							Alt: "Test Image 03",
						},
					},
					Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.\n\nDuis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.\n\nUt wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.",
				},
			},
		},
	}
}
