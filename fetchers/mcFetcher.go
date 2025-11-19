package fetchers;

import (
	"slices"
	"strings"
	"unicode/utf8"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/PuerkitoBio/goquery"
);

var mcLinks []string;
var mcTitles []string;
var mcDescs []string;
var mcImgs []string;

var mcNews []News;

func MCFetcher() (news []News) {
	mcLinks = []string{};
	mcTitles = []string{};
	mcDescs = []string{};
	mcImgs = []string{};

	mcNews = []News{};

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://www.moneycontrol.com/"},
		ParseFunc: mcFetch,
		RobotsTxtDisabled: true,
	}).Start();

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: mcLinks,
		ParseFunc: func (g *geziyor.Geziyor, r *client.Response) {
			mcScrapeMore(r);
		},
		RobotsTxtDisabled: true,
	}).Start();

	for i, link := range mcLinks {
		mcNews = append(mcNews, News{mcTitles[i], mcDescs[i], link, mcImgs[i]});
	}

	return mcNews;
}

func mcFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;

	parsed.Find("div.sub-col-left a").Each(func (_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		title := a.Text();
		if (!slices.Contains(mcLinks, link) &&
			title != "MC EXCLUSIVE" &&
			len(link) >= 55 &&
			link[28:43] == "/news/business/" &&
			link[42:55] != "/commodities/") {
			mcLinks = append(mcLinks, link);
			mcTitles = append(mcTitles, strings.TrimSpace(a.Text()));
			mcDescs = append(mcDescs, "");
			mcImgs = append(mcImgs, "");
		}
	});
	parsed.Find("div.sub-col-right a").Each(func (_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		title := a.Text()
		if (!slices.Contains(mcLinks, link) &&
			title != "MC EXCLUSIVE" &&
			len(link) >= 55 &&
			link[28:43] == "/news/business/" &&
			link[42:55] != "/commodities/") {
			mcLinks = append(mcLinks, link);
			mcTitles = append(mcTitles, strings.TrimSpace(a.Text()));
		}
	});
}

func mcScrapeMore(r *client.Response) {
	parsed := r.HTMLDoc;

	link := r.Request.URL.String();

	desc := parsed.Find("h2.article_desc").First().Text();
	if (utf8.RuneCountInString(strings.TrimSpace(desc)) > 220) {
		mcDescs[slices.Index(mcLinks, link)] = strings.TrimSpace(desc)[:220] + "... Read More";
	} else {
		mcDescs[slices.Index(mcLinks, link)] = strings.TrimSpace(desc);
	}

	src, exists := parsed.Find("div.article_image img").First().Attr("data-src");
	if exists {
		mcImgs[slices.Index(mcLinks, link)] = src;
	}  else {
		mcImgs[slices.Index(mcLinks, link)] = "https://raw.githubusercontent.com/Areen-Rath/ABC-Live/refs/heads/main/assets/logo.png";
	}
}