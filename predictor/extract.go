package predictor

import (
	"errors"
	models "github.com/AlienStream/Shared-Go/models"
	"strings"
)

type Info struct {
	Artist  models.Artist
	Channel models.Channel
	Track   models.Track
	Embed   models.Embeddable
}

func extractInfo(post models.Post) (Info, error) {

	if IsSoundcloud(post) {
		return ExtractInfoFromSoundCloud(post), nil
	}

	if IsYoutube(post) {
		return ExtractInfoFromYoutube(post), nil
	}

	return Info{}, errors.New("No Service Found For Post with URL: " + post.Embed_url)
}

func IsYoutube(p models.Post) bool {

	if strings.Contains(p.Embed_url, "youtube.com") {
		return true
	}

	if strings.Contains(p.Embed_url, "youtu.be") {
		return true
	}

	return false
}

func IsSoundcloud(p models.Post) bool {

	if strings.Contains(p.Embed_url, "soundcloud.com") {
		return true
	}

	return false
}
