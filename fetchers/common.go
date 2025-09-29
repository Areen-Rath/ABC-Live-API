package fetchers

import (
	"regexp";
)

var regex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`);

type News struct {
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
	Link	string	`json:"link"`
	Img		string	`json:"img"`
}