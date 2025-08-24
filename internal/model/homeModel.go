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
				ID:  "home",
				Src: "/static/images/home/test.jpg",
				Alt: "Home Image 01",
			},
			{
				ID:  "home",
				Src: "/static/images/howTo/b001.jpg",
				Alt: "Home Image 01",
			},
			//	{
			//		ID:  "home",
			//		Src: "/static/images/home/b002.jpg",
			//		Alt: "Home Image 02",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b003.jpg",
			//		Alt: "Home Image 03",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b004.jpg",
			//		Alt: "Home Image 04",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b005.jpg",
			//		Alt: "Home Image 05",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b006.jpg",
			//		Alt: "Home Image 06",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b007.jpg",
			//		Alt: "Home Image 07",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b008.jpg",
			//		Alt: "Home Image 08",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b009.jpg;",
			//		Alt: "Home Image 09",
			//	}, {
			//		ID:  "home",
			//		Src: "/static/images/home/b010.jpg",
			//		Alt: "Home Image 10",
			//	},
			//},
		},
	}
}
