package parser

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mehmetkule/film-box/core"
)

var movieList = make(map[string][]string, 0)

func ParserWeb(categoryID int, count int, isFive bool) map[string][]string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Get("https://www.netflix.com/tr/browse/genre/34399")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		} else {

			section := doc.Find(".nm-collections-row")
			section.Each(func(i int, s *goquery.Selection) {
				if !isFive && (i < 3 && categoryID == core.NewNetFlixCategory().IkinciUc || i < 6 && categoryID == core.NewNetFlixCategory().UcuncuUc ||
					i < 9 && categoryID == core.NewNetFlixCategory().DorduncuUc || i < 12 && categoryID == core.NewNetFlixCategory().BesinciUc) {
					//skip
				} else if isFive && (i < 5 && categoryID == core.NewNetFlixCategory().IkinciBes || i < 10 && categoryID == core.NewNetFlixCategory().UcuncuBes) {
					//skip
				} else {
					if (!isFive && i >= 3*categoryID) || (isFive && i >= 5*categoryID) {
						return
					}
					name := strconv.Itoa(i+1) + "-" + s.Find(".nm-collections-row-name").Last().Text()
					liList := s.Find(".nm-content-horizontal-row ul li")

					var list = make([]string, 0)
					liList.Each(func(i2 int, s2 *goquery.Selection) {
						if i2 >= count {
							return
						}
						var maxlength = len(s2.Text())
						if len(s2.Text()) > 30 {
							maxlength = 30
						}
						list = append(list, strconv.Itoa(i2+1)+"."+s2.Text()[:maxlength])
					})
					movieList[name] = list
				}
			})
			return movieList
		}
	} else {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return make(map[string][]string, 0)
}
