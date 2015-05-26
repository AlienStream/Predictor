package predictor

import (
	models "github.com/AlienStream/Shared-Go/models"
)

func SaveInfo(info Info, source models.Source) {
	if info.Embed.IsNew() {
		if info.Artist.IsNew() {
			info.Artist.Insert()
		}
		artist, _ := models.Artist{}.FromName(info.Artist.Name)

		if info.Channel.IsNew() {
			info.Channel.Artist_id = artist.Id
			info.Channel.Insert()
		}
		channel, _ := models.Channel{}.FromUrl(info.Channel.Url)

		if info.Track.IsNew() {
			info.Track.Channel_id = channel.Id
			info.Track.Insert()
		}
		track, _ := models.Track{}.FromTitle(info.Track.Title)
		info.Embed.Track_id = track.Id
		info.Embed.Insert()

		models.CreateTrackSourcePivot(source, track)
	} else {
		embed, _ := models.Embeddable{}.FromUrl(info.Embed.Url)
		track, _ := models.Track{}.FromId(embed.Track_id)
		models.CreateTrackSourcePivot(source, track)
	}

}
