package english

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"profile/db"
	"profile/model"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetWork(c echo.Context) error {
	var works []model.Work
	db := db.Connect()
	db.Table("works_en").Find(&works)
	data, _ := json.Marshal(works)
	fmt.Print(works)

	return c.String(http.StatusOK, string(data))
}

func GetNews(c echo.Context) error {
	var news []model.News
	db := db.Connect()
	db.Table("news_en").Find(&news)
	data, _ := json.Marshal(news)

	return c.String(http.StatusOK, string(data))
}

/* Tackle set */
func GetTackle(c echo.Context) error {

	var tackles []model.Tackle
	var ids = getTackleRss()
	for i := 0; i < len(ids); i++ {
		nextTacke := model.Tackle{}
		nextTacke.Id = ids[i]
		tackles = append(tackles, nextTacke)
	}
	data, _ := json.Marshal(tackles)

	return c.String(http.StatusOK, string(data))
}

func getTackleRss() (ids []string) {
	url := "https://note.com/y_u_k_i_open/rss"

	re, err := regexp.Compile("y_u_k_i_open/n/([a-zA-Z0-9]*)")
	if err != nil {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	links := re.FindAllString(string(byteArray), -1)
	var writtenids []string

	for i := 0; i < len(links); i++ {
		writtenids = append(writtenids, strings.Replace(links[i], "y_u_k_i_open/n/", "", 1))
	}

	ids = toUnique(writtenids)

	return
}

func toUnique(target []string) (result []string) {
	m := map[string]bool{}
	for _, v := range target {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}

	return
}

/* up here */
