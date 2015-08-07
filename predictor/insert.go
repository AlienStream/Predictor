package predictor

import (
	models "github.com/AlienStream/Shared-Go/models"
)

func SaveInfo(info Info, source models.Source) {
	if info.Embed.IsNew() {
		if info.Artist.IsNew() {
			info.Artist.Insert()
		} else {
			info.Artist.Save()
		}
		artist, _ := models.Artist{}.FromName(info.Artist.Name)

		info.Channel.Artist_id = artist.Id
		if info.Channel.IsNew() {
			info.Channel.Insert()
		} else {
			info.Channel.Save()
		}
		channel, _ := models.Channel{}.FromUrl(info.Channel.Url)
		
		info.Track.Channel_id = channel.Id
		if info.Track.IsNew() {
			info.Track.Insert()
		} else {
			info.Track.Save()
		}
		track, _ := models.Track{}.FromTitle(info.Track.Title)

		info.Embed.Track_id = track.Id
		info.Embed.Insert()

		models.CreateTrackSourcePivot(source, track)
	} else {
		embed, _ := models.Embeddable{}.FromUrl(info.Embed.Url)
		track, _ := models.Track{}.FromId(embed.Track_id)
		track.Rank = info.Track.Rank
		models.CreateTrackSourcePivot(source, track)
	}
}
