package main

import (
	"time"
	"github.com/AlienStream/Predictor/predictor"
	models "github.com/AlienStream/Shared-Go/models"
)

func main() {
	longinterval := time.NewTicker(time.Minute * 15).C;

	go func() {
            for {
                select {
                case <- longinterval:
                    importLatestTracks()
                    break
              }
            }
        }()
}

func importLatestTracks() {
	posts := models.NewPosts()
	for _, post := range posts {
		predictor.ImportFromPost(post)
	}

}

func refreshAllTracks() {
	posts := models.AllPosts()
	for _, post := range posts {
		go predictor.ImportFromPost(post)
	}
}

func refreshExpiredTracks() {
	posts := models.AllPosts()
	for _, post := range posts {
		go predictor.ImportFromPost(post)
	}
}
