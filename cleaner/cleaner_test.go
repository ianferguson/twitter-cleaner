package cleaner

import (
	"testing"
	"fmt"
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

	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
	}
	t.Error("not yet implemented")
}
