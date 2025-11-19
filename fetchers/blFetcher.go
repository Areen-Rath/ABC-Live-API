package fetchers;

import (
	"slices"
	"strings"
	"unicode/utf8"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/PuerkitoBio/goquery"
);

var blLinks []string;
var blTitles []string;
var blDescs []string;
var blImgs []string;

var blNews []News;

func BLFetcher() (news []News) {
	blLinks = []string{};
	blTitles = []string{};
	blDescs = []string{};
	blImgs = []string{};

	blNews = []News{};

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://www.thehindubusinessline.com/"},
		ParseFunc: blFetch,
		RobotsTxtDisabled: true,
	}).Start();

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: blLinks,
		ParseFunc: func (g *geziyor.Geziyor, r *client.Response) {
			blScrapeMore(r);
		},
		RobotsTxtDisabled: true,
	}).Start();

	for i, link := range blLinks {
		if blDescs[i] != "" {
			blNews = append(blNews, News{blTitles[i], blDescs[i], link, blImgs[i]});
		}
	}

	return blNews;
}

func blFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;

	parsed.Find("div.after-border-right a").Each(func (_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		title := a.Text();
		if(!slices.Contains(blLinks, link) &&
			title != "" &&
			title != "LIVE" &&
			link[36:42] != "/news/" &&
			link[len(link) - 4:] == ".ece") {
			blLinks = append(blLinks, link);
			blTitles = append(blTitles, strings.TrimSpace(title));
			blDescs = append(blDescs, "");
			blImgs = append(blImgs, "");
		}
	});
}

func blScrapeMore(r *client.Response) {
	parsed := r.HTMLDoc;

	link := r.Request.URL.String();

	desc := parsed.Find("h2.sub-title").First().Text();
	if !strings.Contains(desc, "LIVE") {
		if (utf8.RuneCountInString(strings.TrimSpace(desc)) > 220) {
			blDescs[slices.Index(blLinks, link)] = strings.TrimSpace(desc)[:220] + "... Read More";
		} else {
			blDescs[slices.Index(blLinks, link)] = strings.TrimSpace(desc);
		}
	}

	src, exists := parsed.Find("source").First().Attr("srcset");
	if exists {
		blImgs[slices.Index(blLinks, link)] = src;
	} else {
		blImgs[slices.Index(blLinks, link)] = "https://raw.githubusercontent.com/Areen-Rath/ABC-Live/refs/heads/main/assets/logo.png";
	}
}