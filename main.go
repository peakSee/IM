package main

import "IM/router"

func main() {
	e := router.Router()

	e.Run(":8080")
}
