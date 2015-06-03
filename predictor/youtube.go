package predictor

import (
	"encoding/json"
	models "github.com/AlienStream/Shared-Go/models"
	"net/http"
	"strings"
)

type YoutubeRoot struct {
	Kind  string         `json:"kind"`
	Items []YoutubeVideo `json:"items"`
}

type YoutubeVideo struct {
	Id      string         `json:"id"`
	Snippet YoutubeSnippet `json:"snippet"`
	Stats   YoutubeStats   `json:"statistics"`
	Status  YoutubeStatus  `json:"status"`
}

type YoutubeSnippet struct {
	Channel       string               `json:"channelId"`
	Channel_title string               `json:"channelTitle"`
	Description   string               `json:"description"`
	Title         string               `json:"title"`
	Published_at  string               `json:"publishedAt"`
	Thumbnail     YoutubeThumbnailList `json:"thumbnails"`
}

type YoutubeThumbnailList struct {
	Default  YoutubeThumbnail `json:"default"`
	Medium   YoutubeThumbnail `json:"medium"`
	High     YoutubeThumbnail `json:"high"`
	Standard YoutubeThumbnail `json:"standard"`
	MaxRes   YoutubeThumbnail `json:"maxres"`
}

type YoutubeThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type YoutubeStatus struct {
	Privacy    string `json:"privacyStatus"`
	Embeddable bool   `json:"embeddable"`
}

type YoutubeStats struct {
	viewCount     int `json:"viewCount"`
	likeCount     int `json:"likeCount"`
	dislikeCount  int `json:"dislikeCount"`
	favoriteCount int `json:"favoriteCount"`
	commentCount  int `json:"commentCount"`
}

func ExtractInfoFromYoutube(p models.Post) Info {
	video := YoutubeVideoLookup(p.Embed_url)

	artist := models.Artist{
		Name:      video.Snippet.Channel_title,
		Thumbnail: video.Snippet.Thumbnail.MaxRes.URL,
	}

	channel := models.Channel{
		Url: "https://www.youtube.com/channel/" + video.Snippet.Channel,
	}

	var track_rank float64
	s, _ := models.Source{}.FromId(p.Source_id)
	track_rank = (float64) (p.Likes * s.Importance)
	
	track := models.Track{
		Rank:      track_rank,
		Title:     video.Snippet.Title,
		Thumbnail: video.Snippet.Thumbnail.MaxRes.URL,
	}

	embeddable := models.Embeddable{
		Url:  p.Embed_url,
		Type: "youtube",
	}

	return Info{
		artist,
		channel,
		track,
		embeddable,
	}
}

func YoutubeVideoLookup(url string) YoutubeVideo {
	var root YoutubeRoot
	var video YoutubeVideo
	if !strings.Contains(url, "v=") {
		return video
	}

	// get the video ID
	var v_id string = strings.Split(url, "v=")[1]
	v_id = strings.Split(v_id, "&")[0]

	// setup the request
	var base_url string = "https://www.googleapis.com/youtube/v3/videos"
	var params = []string{"part=snippet%2C+statistics%2C+status%2C+player", "id=" + v_id, "key=AIzaSyD5nqLKBD8adUKk9UVT1X9KfCcqcwcZuP8"}
	var curl_url = base_url + "?" + strings.Join(params, "&")

	// get the data
	client := &http.Client{}
	req, _ := http.NewRequest("GET", curl_url, nil)
	req.Header.Set("User-Agent", "AlienStream Master Server v. 1.0")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&root)

	if err == nil && len(root.Items) != 0 {
		video = root.Items[0]
	} else {
		//panic(err)
	}

	//fmt.Printf("decoding "+curl_url+" \n ")

	return video
}
