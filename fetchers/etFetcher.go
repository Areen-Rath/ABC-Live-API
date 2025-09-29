package fetchers;

import (
	"slices"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
);

var etLinks []string;
var etTitles []string;
var etDescs []string;
var etImgs []string;

var etNews []News;

func ETFetcher() (news []News) {
	etLinks = []string{};
	etTitles = []string{};
	etDescs = []string{};
	etImgs = []string{};

	etNews = []News{};

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://www.economictimes.com/markets"},
		ParseFunc: etFetch,
		RobotsTxtDisabled: true,
	}).Start();
	
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: etLinks,
		ParseFunc: func (g *geziyor.Geziyor, r *client.Response) {
			etScrapeMore(r);
		},
		RobotsTxtDisabled: true,
	}).Start();

	for i, link := range etLinks {
		if etDescs[i] != "" {
			etNews = append(etNews, News{etTitles[i], etDescs[i], link, etImgs[i]});
		}
	}

	return etNews;
}

func etFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;

	parsed.Find("div.FirstFoldWidget_topStories__tBR9H a").Each(func(_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		if(!slices.Contains(etLinks, link) &&
			len(link) >= 45 &&
			link[36:45] == "/markets/" &&
			link[44:57] != "/commodities/" &&
			link[44:57] != "/expert-view/") {
				etLinks = append(etLinks, link);
				etTitles = append(etTitles, strings.TrimSpace(a.Text()));
				etDescs = append(etDescs, "");
				etImgs = append(etImgs, "");
		}
	});
}

func etScrapeMore(r *client.Response) {
	parsed := r.HTMLDoc;

	link := r.Request.URL.String();

	desc := parsed.Find("p.summary").First().Text();
	etDescs[slices.Index(etLinks, link)] = strings.TrimSpace(desc);

	src, exists := parsed.Find("figure.artImg img").First().Attr("src");
	if exists {
		etImgs[slices.Index(etLinks, link)] = src;
	} else {
		etImgs[slices.Index(etLinks, link)] = "https://raw.githubusercontent.com/Areen-Rath/ABC-Live/refs/heads/main/assets/logo.png";
	}
}