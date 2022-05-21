package helper

import "fmt"

const (
	twitterBASEURL = "https://twitter.com/%s/status/%s"
)

func TwitterURLBuilder(handler, tweetID string) string {
	return fmt.Sprintf(twitterBASEURL, handler, tweetID)
}
