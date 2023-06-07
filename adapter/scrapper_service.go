package adapter

import (
	"github.com/gocolly/colly/v2"
	"github.com/mijara/patchlauncher/model"
	"net/url"
	"strings"
	"time"
)

const (
	_defaultWebBase   = "https://www.smwcentral.net"
	_defaultWebSuffix = "/?p=section&s=smwhacks"
)

type ScrapperService struct {
	webBase   string
	webSuffix string
}

func NewScrapperService(webBase, webSuffix string) *ScrapperService {
	return &ScrapperService{
		webBase:   webBase,
		webSuffix: webSuffix,
	}
}

func (s *ScrapperService) ScrapHackList() ([]model.Hack, error) {
	hacks := make([]model.Hack, 0)

	c := colly.NewCollector(colly.DetectCharset())

	c.OnHTML("#list-content .content tbody tr", func(element *colly.HTMLElement) {
		hack := model.Hack{}

		element.ForEach("td", func(i int, element *colly.HTMLElement) {
			switch i {
			case 0:
				hack.Title = element.ChildText("a")

				hackURL, err := url.QueryUnescape(s.webBase + element.ChildAttr("a", "href"))
				if err != nil {
					return
				}
				hack.URL = hackURL

				uploadedAt, err := time.Parse("2006-01-02T15:04:05", element.ChildAttr("span time", "datetime"))
				if err != nil {
					return
				}
				hack.UploadedAt = uploadedAt
			case 4:
				hack.Type = element.Text
			case 5:
				hack.Authors = element.Text
			case 6:
				hack.Rating = element.Text
			case 8:
				hack.DownloadURL = "https:" + element.ChildAttr("a", "href")
				hack.Downloads = strings.Split(element.ChildText(".secondary-info"), " ")[0]
			}
		})

		hacks = append(hacks, hack)
	})

	if err := c.Visit(s.webBase + s.webSuffix); err != nil {
		return nil, err
	}

	return hacks, nil
}
