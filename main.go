package main

import (
	"fmt"

	"github.xiadat.com/goddd/domain/models"
	"github.xiadat.com/goddd/domain/valueobjs"
)

func main() {
	user := models.NewUserModel(
		models.WithUserName("taoshumin"),
		models.WithUserExtra(
			valueobjs.NewUserExtra(
				valueobjs.WithUserCity("beijing"),
				valueobjs.WithUserQQ("174300000"),
			),
		),
	)
	fmt.Println(user)
}
