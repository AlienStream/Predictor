package predictor

import (
	"encoding/json"
	models "github.com/AlienStream/Shared-Go/models"
	"net/http"
	"net/url"
	"strings"
)

type SoundCloudoEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail_url"`
	Author_Name string `json:"author_name"`
	Author_Url  string `json:"url"`
}


func SoundCloudoEmbedLookup(uri string) SoundCloudoEmbed {
	var oEmbed SoundCloudoEmbed
	// setup the request
	var base_url string = "http://soundcloud.com/oembed"
	var params = []string{"format=json","url="+ url.QueryEscape(uri)}
	var curl_url = base_url +"?"+ strings.Join(params, "&")

	// get the data
	client := &http.Client{}
	req,_ := http.NewRequest("GET",curl_url,nil)
    req.Header.Set("User-Agent","AlienStream Master Server v. 1.0")
    resp,_ := client.Do(req)
    defer resp.Body.Close();
    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&oEmbed)

    //fmt.Printf("decoding " + curl_url + " \n ")
    return oEmbed
}

func ExtractInfoFromSoundCloud(p models.Post) Info {
	track_info := SoundCloudoEmbedLookup(p.Embed_url)

	artist := models.Artist{
		Name:      track_info.Author_Name,
		Thumbnail: track_info.Thumbnail,
	}

	channel := models.Channel{
		Url: track_info.Author_Url,
	}

	var track_rank float64
	s, _ := models.Source{}.FromId(p.Source_id)
	track_rank = (float64) (p.Likes * s.Importance)
	
	track := models.Track{
		Rank:      track_rank,
		Title:     track_info.Title,
		Thumbnail: track_info.Thumbnail,
		Created_at: p.Posted_at,
	}

	embeddable := models.Embeddable{
		Url:  p.Embed_url,
		Type: "soundcloud",
	}

	return Info{
		artist,
		channel,
		track,
		embeddable,
	}
}


