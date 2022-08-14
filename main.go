package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

var (
	supported_seasons = map[string]string{
		"2017.2018": "https://www.zerozero.pt/edition.php?id_edicao=111450",
		"2018.2019": "https://www.zerozero.pt/edition.php?id_edicao=125348",
		"2019.2020": "https://www.zerozero.pt/edition.php?id_edicao=136037",
		"2020.2021": "https://www.zerozero.pt/edition.php?id_edicao=148340",
		"2021.2022": "https://www.zerozero.pt/edition.php?id_edicao=157056",
	}
)

type table_line struct {
	Position	string	`json:"position"`
	Team		string	`json:"team"`
	Games		string	`json:"games"`
	Victories	string 	`json:"victories"`
	Draws		string 	`json:"draws"`
	Defeats		string 	`json:"defeats"`
	PlusMinus	string	`json:"plus_minus"`
	Points		string	`json:"points"`
} 

func get_available_seasons() []string {
	keys := make([]string, 0, len(supported_seasons))
	for key := range supported_seasons {
		keys = append(keys, key)
	}
	return keys
}

func perform_scraping_on_zerozero(url string) []table_line {
	classification := []table_line{}
	c := colly.NewCollector()
	c.OnHTML("div[id=edition_table].box_container tbody tr", func(h *colly.HTMLElement) {
		line := table_line{}
		h.ForEach("td", func(order int, el *colly.HTMLElement) {
			switch order {
			case 0:
				line.Position = el.Text
			case 2:
				line.Team = el.Text
			case 3:
				line.Points = el.Text
			case 4:
				line.Games = el.Text
			case 5:
				line.Victories = el.Text
			case 6:
				line.Draws = el.Text
			case 7:
				line.Defeats = el.Text
			case 10:
				line.PlusMinus = el.Text
			default:
				break
			}
		})
		classification = append(classification, line)
	})
	c.Visit(url)
	return classification
}

func retrieve_available_seasons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, get_available_seasons())
}

func retrieve_classification_table(c *gin.Context) {
	season := c.Param("season")
	value, ok := supported_seasons[season]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, "The request season is not supported")
		return
	}
	classification := perform_scraping_on_zerozero(value)
	c.IndentedJSON(http.StatusOK, classification)
}

func bootstrap_api() {
	router := gin.Default()
	router.GET("/available_seasons", retrieve_available_seasons)
	router.GET("/:season", retrieve_classification_table)
	router.Run("localhost:8080")
	log.Println("Started to listen on localhost:8080")
}

func main() {
	bootstrap_api()
}
