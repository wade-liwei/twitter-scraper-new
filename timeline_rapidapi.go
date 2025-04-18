package twitterscraper

// timeline v2 JSON object
type TimelineV2RapidAPI struct {
	Result struct {
			Timeline struct {
				Instructions []struct {
					Entries []entry `json:"entries"`
					Entry   entry   `json:"entry"`
					Type    string  `json:"type"`
				} `json:"instructions"`
			} `json:"timeline"`
	} `json:"result"`
}


func (timeline *TimelineV2RapidAPI) ParseTweets() ([]*Tweet, string) {
	var cursor string
	var tweets []*Tweet
	for _, instruction := range timeline.Result.Timeline.Instructions {

		for _, entry := range instruction.Entries {
			if entry.Content.CursorType == "Bottom" {
				cursor = entry.Content.Value
				continue
			}
			if entry.Content.ItemContent.TweetResults.Result.Typename == "Tweet" || entry.Content.ItemContent.TweetResults.Result.Typename == "TweetWithVisibilityResults" {
				if tweet := entry.Content.ItemContent.TweetResults.Result.parse(); tweet != nil {
					tweets = append(tweets, tweet)
				}
			}
			if len(entry.Content.Items) > 0 {
				for _, item := range entry.Content.Items {
					if tweet := item.Item.ItemContent.TweetResults.Result.parse(); tweet != nil {
						tweets = append(tweets, tweet)
					}
				}
			}
		}
	}

	return tweets, cursor
}
