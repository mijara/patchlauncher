package adapter

import (
	"github.com/gocolly/colly/v2"
	"smwlauncher/port"
	"strings"
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

func (s *ScrapperService) ScrapHackList() ([]port.ROMHack, error) {
	hacks := make([]port.ROMHack, 0)

	c := colly.NewCollector()

	c.OnHTML("#list-content .content tbody tr", func(element *colly.HTMLElement) {
		hack := port.ROMHack{}

		element.ForEach("td", func(i int, element *colly.HTMLElement) {
			switch i {
			case 0:
				hack.Title = element.ChildText("a")
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
