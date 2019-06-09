package twittercleaner

import (
	"fmt"
	"testing"
)

func TestThatSearchTweetsCanBeRetrieved(t *testing.T) {
	cleaner, err := New()
	if err != nil {
		t.Fatal(err)
	}

	tweets, err := cleaner.getTweets()
	if err != nil {
		t.Fatal(err)
	}

	for _, tweet := range *tweets {
		fmt.Printf("%s: %s\n", tweet.User.Name, tweet.Text)
	}
	t.Error("not yet implemented")
}

func TestThatUserTimelineCanBeRetrieved(t *testing.T) {
	cleaner, err := New()
	if err != nil {
		 t.Fatal(err)
	}

	tweets, err := cleaner.GetMyTweets()
	if err != nil {
		t.Fatal(err)
	}

	for _, tweet := range *tweets {
		fmt.Printf("%s: %s\n", tweet.User.Name, tweet.Text)
	}

	fmt.Printf("Found %d total tweets", len(*tweets))
	t.Error("not yet implemented")
}

func TestTweetFiltering(t *testing.T) {
	cleaner, err := New()
	if err != nil {
		t.Fatal(err)
	}

	tweets, err := cleaner.GetMyTweets()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Found %d total tweets", len(*tweets))

	unengaged := filterUnengagedTweets(tweets)
	for _, tweet := range *unengaged {
		fmt.Printf("delete: %s\n", tweet.Text)
	}

	engaged := filterEngagedTweets(tweets)
	for _, tweet := range *engaged {
		fmt.Printf("keep: %s\n", tweet.Text)
	}

	t.Error("not yet implemented")
}


func TestThatTweetsOlderThan30DaysWithoutEngagementsAreRemoved(t *testing.T) {
	t.Error("not yet implemented")
}
