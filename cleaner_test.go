package twittercleaner

import (
	"fmt"
	"testing"
)

func TestThatTweetsCanBeRetrieved(t *testing.T) {
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

func TestThatTweetsOlderThan30DaysWithoutEngagementsAreRemoved(t *testing.T) {
	t.Error("not yet implemented")
}
