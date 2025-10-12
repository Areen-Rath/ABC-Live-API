package fetchers

type News struct {
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
	Link	string	`json:"link"`
	Img		string	`json:"img"`
}