package main

import (
	"fmt"
	"os"

	"github.com/askft/newsapi"
)

func main() {
	apiKey := os.Getenv("NEWS_API_KEY")

	c := newsapi.NewClient(apiKey)

	res, err := c.GetEverything(newsapi.EverythingRequestParams{
		Language: newsapi.LanguageEnglish,
		Keywords: "+millennials",
		PageSize: 100,
	})

	if err != nil {
		panic(err)
	}

	for _, src := range res.Articles {
		fmt.Println(src.Title)
		fmt.Println(src.Description)
		fmt.Println()
	}
}
