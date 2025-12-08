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
		},
	}
}
