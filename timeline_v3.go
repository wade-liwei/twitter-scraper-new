package twitterscraper

type TimelineV3 struct {
	Data struct {
		User struct {
			Result struct {
				Timeline struct {
					Timeline struct {
						Instructions []struct {
							Entries []entry `json:"entries"`
							Entry   entry   `json:"entry"`
							Type    string  `json:"type"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

func (timeline *TimelineV3) ParseTweets() ([]*Tweet, string) {
	var cursor string
	var tweets []*Tweet
	for _, instruction := range timeline.Data.User.Result.Timeline.Timeline.Instructions {
		if instruction.Type == "TimelinePinEntry" {
			entry := instruction.Entry
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



func (timeline *TimelineV3) ParseUsers() ([]*Profile, string) {
	var cursor string
	var users []*Profile
	for _, instruction := range timeline.Data.User.Result.Timeline.Timeline.Instructions {
		for _, entry := range instruction.Entries {
			if entry.Content.CursorType == "Bottom" {
				cursor = entry.Content.Value
				continue
			}
			if entry.Content.ItemContent.UserResults.Result.Typename == "User" {
				user := entry.Content.ItemContent.UserResults.Result.parse()
				users = append(users, &user)
			}
		}
	}
	return users, cursor
}
