package main

import (
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/matryer/is"
)

func setup(t *testing.T) *NatsConnectionMock {
	mockedNatsConnection := &NatsConnectionMock{
		PublishFunc: func(subj string, data []byte) error {
			return nil
		},
	}

	nc = mockedNatsConnection

	return mockedNatsConnection
}

func TestShouldSendMessgaeWhenContainsHandle(t *testing.T) {
	m := setup(t)
	is := is.New(t)

	tweetMessage := "Hello there @sheriff_bot"

	tw := &twitter.Tweet{
		Text: tweetMessage,
		User: &twitter.User{},
	}

	handleTweet(tw)

	is.Equal(len(m.PublishCalls()), 1) // should have a called publish once

	c := m.PublishCalls()[0]
	is.Equal(c.Subj, "tweet")              // should have set the message name to tweet
	is.Equal(c.Data, []byte(tweetMessage)) // should have passed the message on
}

func TestShouldNotSendMessgaeWhenNotContainsHandle(t *testing.T) {
	m := setup(t)
	is := is.New(t)

	tw := &twitter.Tweet{
		Text: "hello",
		User: &twitter.User{},
	}

	handleTweet(tw)

	is.Equal(len(m.PublishCalls()), 0) // should not have a called publish
}
