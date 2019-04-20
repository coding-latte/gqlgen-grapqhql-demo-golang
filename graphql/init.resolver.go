package graphql

import (
	"github.com/dghubble/go-twitter/twitter"
)

// NewRootResolver - return the root resolver with twitter client
// field set. The field is private
func NewRootResolver(client *twitter.Client) Config {
	c := Config{
		Resolvers: &Resolver{
			client: client,
		},
	}

	return c
}
