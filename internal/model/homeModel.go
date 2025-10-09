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
			// {
			// 	ID:  "home2",
			// 	Src: "/static/images/home/b002.jpg",
			// 	Alt: "Home Image 02",
			// },
			// {
			// 	ID:  "home3",
			// 	Src: "/static/images/home/b003.jpg",
			// 	Alt: "Home Image 03",
			// },
			// {
			// 	ID:  "home4",
			// 	Src: "/static/images/home/b004.jpg",
			// 	Alt: "Home Image 04",
			// },
			// {
			// 	ID:  "home6",
			// 	Src: "/static/images/home/b006.jpg",
			// 	Alt: "Home Image 06",
			// },
			// {
			// 	ID:  "home7",
			// 	Src: "/static/images/home/b007.jpg",
			// 	Alt: "Home Image 07",
			// },
			// {
			// 	ID:  "home8",
			// 	Src: "/static/images/home/b008.jpg",
			// 	Alt: "Home Image 08",
			// },
			// {
			// 	ID:  "home9",
			// 	Src: "/static/images/home/b009.jpg",
			// 	Alt: "Home Image 09",
			// },
			// {
			// 	ID:  "home10",
			// 	Src: "/static/images/home/b010.jpg",
			// 	Alt: "Home Image 10",
			// },
			// {
			// 	ID:  "home11",
			// 	Src: "/static/images/home/b011.jpg",
			// 	Alt: "Home Image 11",
			// },
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
