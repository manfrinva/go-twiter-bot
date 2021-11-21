package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRET_KEY"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Client do twitter
	client := twitter.NewClient(httpClient)

	// Publicar um tweet
	_, _, err := client.Statuses.Update("uma mensagem qualquer", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Procurar por tweets
	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#desenvolvimento",
		Count: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Verificar todos os tweets listados
	for _, val := range tweets.Statuses {
		log.Print("Nome do usuario :", val.User.Name)
		log.Print("Tweet :", val.Text)

		// Retweetar os tweets listados
		_, _, err := client.Statuses.Retweet(val.ID, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
