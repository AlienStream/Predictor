package predictor

import (
	"fmt"
	"strings"
	db "github.com/AlienStream/Shared-Go/database"
	models "github.com/AlienStream/Shared-Go/models"
)

func ImportFromPost(p models.Post) {
	//broken edge cases
	if strings.Contains(p.Embed_url, "/groups/") && IsSoundcloud(p) {
		return
	}
	
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

	p.Is_new = false;
	p.Save();
}
