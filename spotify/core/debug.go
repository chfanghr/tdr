package core

import (
	"log"
	"strings"
)

var Version = "DEBUG"
var BuildID = "0"

var PrintDebugMessage = func(msg string) {
	if Version == "DEBUG" || strings.ContainsAny(Version, "beta") || strings.ContainsAny(Version, "alpha") {
		log.Println(msg)
		//TODO
	}
}
