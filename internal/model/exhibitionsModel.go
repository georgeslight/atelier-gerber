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
				ID:       "moma",
				Title:    "\"From Inside Out\", Lightroom, P.S.1 New York",
				Location: "International Studio Program (Winter 1978-1979) - MoMA PS1, New York",
				Date:     "1978 - 1979",
				Href:     "/exhibitions/moma",
				BaseImage: ExhibitionImage{
					Src: "/static/images/exhibitions/moma/moma-00.jpg",
					Alt: "MoMA PS1",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/exhibitions/moma/moma-01.jpg",
							Alt: "MoMA detail 01",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-02.jpg",
							Alt: "MoMA detail 02",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-03.jpg",
							Alt: "MoMA detail 03",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-04.jpg",
							Alt: "MoMA detail 04",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-05.jpg",
							Alt: "MoMA detail 05",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-06.jpg",
							Alt: "MoMA detail 06",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-07.jpg",
							Alt: "MoMA detail 07",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-08.jpg",
							Alt: "MoMA detail 08",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-09.jpg",
							Alt: "MoMA detail 09",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-10.jpg",
							Alt: "MoMA detail 10",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-11.jpg",
							Alt: "MoMA detail 11",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-12.jpg",
							Alt: "MoMA detail 12",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-13.jpg",
							Alt: "MoMA detail 13",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-14.jpg",
							Alt: "MoMA detail 14",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-15.jpg",
							Alt: "MoMA detail 15",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-16.jpg",
							Alt: "MoMA detail 16",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-17.jpg",
							Alt: "MoMA detail 17",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-18.jpg",
							Alt: "MoMA detail 18",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-19.jpg",
							Alt: "MoMA detail 19",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-20.jpg",
							Alt: "MoMA detail 20",
						},
						{
							Src: "/static/images/exhibitions/moma/moma-21.jpg",
							Alt: "MoMA detail 21",
						},
					},
					Description: "48.000 Coffe Cups, String 2,1 Miles Long.\n\nKristin Gerber's Open Studio Exhibition is sponsored by the Institude for Art and\nUrban Resources' International Comitee with the support of the City of West\nBerlin, DADD.\n\nP.S.1, a Center for the Experimental Arts, is a project of the Institude for Art\nand Urban Resources, Inc., and is supported in part by the New York State\nCouncil on the Arts and the National Endowment for the Arts.\n\nP.S.1 (Project Studios One)\n46-01 21st Street\nLong Island City, Queens, NY 11101",
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
					Src: "/static/images/exhibitions/group-01/new-01.jpg",
					Alt: "Test Image 01",
				},
				Details: ExhibitionDetails{
					Images: []ExhibitionImage{
						{
							Src: "/static/images/exhibitions/group-01/new-01.jpg",
							Alt: "Test Image 01",
						},
						{
							Src: "/static/images/exhibitions/group-01/new-02.jpg",
							Alt: "Test Image 02",
						},
						{
							Src: "/static/images/exhibitions/group-01/new-03.jpg",
							Alt: "Test Image 03",
						},
						{
							Src: "/static/images/exhibitions/group-01/new-04.jpg",
							Alt: "Test Image 04",
						},
						{
							Src: "/static/images/exhibitions/group-01/new-05.jpg",
							Alt: "Test Image 05",
						},
					},
					Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.\n\nDuis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.\n\nUt wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.",
				},
			},
		},
	}
}
