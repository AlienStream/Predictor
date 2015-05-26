package predictor

import (
	"fmt"
	db "github.com/AlienStream/Shared-Go/database"
	models "github.com/AlienStream/Shared-Go/models"
)

func ImportFromPost(p models.Post) {
	if IsYoutube(p) {
		fmt.Printf("Importing Post %s \n", p.Title)
		info, e_err := extractInfo(p)
		if e_err != nil {
			panic(e_err)
		}
		rows, _, q_err := db.Con.Query("select * from sources where `id`=%d", p.Source_id)
		if q_err != nil {
			panic(q_err)
		}
		source := models.RowsToSources(rows)[0]
		SaveInfo(info, source)
	}
}
