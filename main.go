package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nats-io/nats"
)

var nc NatsConnection

func main() {
	var err error
	nc, err = nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats server")
	}

	consumerToken := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerToken, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	// User stream
	userParams := &twitter.StreamUserParams{
		StallWarnings: twitter.Bool(true),
		Language:      []string{"en"},
	}

	stream, err := client.Streams.User(userParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	stream.Stop()
}

func handleTweet(tweet *twitter.Tweet) {
	if strings.Contains(tweet.Text, "@sheriff_bot") {
		sendMessage(tweet.Text)
	}
}

func sendMessage(m string) {
	nc.Publish("tweet", []byte(m))
}
