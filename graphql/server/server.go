package main

import (
	"github.com/coding-latte/gqlgen-grapgql-demo/graphql"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/dghubble/oauth1"
)

const defaultPort = "8080"

func graphqlHandler() http.Handler {

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return handler.GraphQL(
		graphql.NewExecutableSchema(
			graphql.NewRootResolver(client),
		),
	)
}

// Defining the Playground handler
func playgroundHandler() http.Handler {
	return handler.Playground("GraphQL playground", "/query")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playgroundHandler())
	http.Handle("/query", graphqlHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
