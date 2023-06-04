package main

import (
	"IM/models"
	"IM/router"
)

func main() {
	e := router.Router()

	e.Run(models.OptionsConfig.Server.Host + ":" + models.OptionsConfig.Server.Port)
}
