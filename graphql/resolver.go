//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	client *twitter.Client
}

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

	if post.Text == "" {
		return nil, fmt.Errorf("Post can not be empty")
	}

	tweet, _, err := r.client.Statuses.Update(post.Text, nil)

	if err != nil {
		// NOTE: In production you might not want to send this error to client
		return nil, fmt.Errorf("Ooops! Something went wrong: %v", err)
	}

	return tweet, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Search(ctx context.Context, query string) ([]twitter.Tweet, error) {

	if query == "" {
		return nil, fmt.Errorf("Please provide a valid search term")
	}

	search, _, err := r.client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
	})

	if err != nil {
		// NOTE: In production you might not want to send this error to client
		return nil, fmt.Errorf("Ooops! Something went wrong: %v", err)
	}

	// we are blatantly ignoring everything else apart from tweets
	return search.Statuses, nil
}

type tweetResolver struct{ *Resolver }

func (r *tweetResolver) ID(ctx context.Context, obj *twitter.Tweet) (string, error) {
	return obj.IDStr, nil
}
func (r *tweetResolver) CreatedAt(ctx context.Context, obj *twitter.Tweet) (*time.Time, error) {
	time, err := obj.CreatedAtTime()
	return &time, err
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *twitter.User) (string, error) {
	return obj.IDStr, nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *twitter.User) (*time.Time, error) {
	time, err := time.Parse(time.RubyDate, obj.CreatedAt)
	return &time, err
}
