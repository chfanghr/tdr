package main

import (
	"github.com/chfanghr/tdr/spotify/utils"
	"log"
)

func main() {
	log.Println(utils.UnwrapResultFromJob(func() {
		utils.ThrowData(1)
	}))
}
