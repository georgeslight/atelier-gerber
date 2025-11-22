package model

type HomeImage struct {
	ID  string `json:"id"`
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type HomeData struct {
	HomeImages []HomeImage `json:"homeImages"`
}

func GetHomeData() HomeData {
	return HomeData{
		HomeImages: []HomeImage{
			{
				ID:  "home1",
				Src: "/static/images/home/kristin-01.jpg",
				Alt: "Home Image 01",
			},
			{
				ID:  "home2",
				Src: "/static/images/home/kristin-02.jpg",
				Alt: "Home Image 02",
			},
			{
				ID:  "home3",
				Src: "/static/images/home/kristin-03.jpg",
				Alt: "Home Image 03",
			},
			{
				ID:  "home4",
				Src: "/static/images/home/kristin-04.jpg",
				Alt: "Home Image 04",
			},
			{
				ID:  "home5",
				Src: "/static/images/home/kristin-05.jpg",
				Alt: "Home Image 05",
			},
			{
				ID:  "home6",
				Src: "/static/images/home/kristin-06.jpg",
				Alt: "Home Image 06",
			},
			{
				ID:  "home7",
				Src: "/static/images/home/kristin-07.jpg",
				Alt: "Home Image 07",
			},
			{
				ID:  "home8",
				Src: "/static/images/home/kristin-08.jpg",
				Alt: "Home Image 08",
			},
			{
				ID:  "home09",
				Src: "/static/images/home/kristin-test-01.jpg",
				Alt: "Home Image 09",
			},
			{
				ID:  "home10",
				Src: "/static/images/home/kristin-test-02.jpg",
				Alt: "Home Image 10",
			},
			// {
			// 	ID:  "home12",
			// 	Src: "/static/images/home/b012.jpg",
			// 	Alt: "Home Image 12",
			// },
			// {
			// 	ID:  "home13",
			// 	Src: "/static/images/home/b013.jpg",
			// 	Alt: "Home Image 13",
			// },
		},
	}
}
