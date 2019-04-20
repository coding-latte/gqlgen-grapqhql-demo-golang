//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Tweet() TweetResolver {
	return &tweetResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostTweet(ctx context.Context, post Post) (*twitter.Tweet, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Tweets(ctx context.Context) ([]twitter.Tweet, error) {
	panic("not implemented")
}

type tweetResolver struct{ *Resolver }

func (r *tweetResolver) ID(ctx context.Context, obj *twitter.Tweet) (string, error) {
	panic("not implemented")
}
func (r *tweetResolver) CreatedAt(ctx context.Context, obj *twitter.Tweet) (*time.Time, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *twitter.User) (string, error) {
	panic("not implemented")
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *twitter.User) (*time.Time, error) {
	panic("not implemented")
}
