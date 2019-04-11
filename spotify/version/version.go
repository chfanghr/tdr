package version

import (
	"log"
	"os"
	"regexp"
	. "github.com/chfanghr/tdr/spotify/utils"
)

const versionRegExp = "(^v[0-9].[0-9].[0-9](.(beta|alpha))*$|^DEBUG$)"
const debugVersionRegExp = "(^v[0-9].[0-9].[0-9].(beta|alpha)$|^DEBUG$)"

var Version = "DEBUG" // Version should match versionRegExp above,filled when compile
var BuildID = "0"     //BuildID is git commit id,short version,filled when compile

var Debug *log.Logger

func IsDebugBuild() bool { return regexp.MustCompile(debugVersionRegExp).Match([]byte(Version)) }

func setupDebugLogger() {
	if regexp.MustCompile(debugVersionRegExp).Match([]byte(Version)) {
		Debug = log.New(os.Stdout, "spotify_debug", log.LstdFlags)
	} else {
		nullDevFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0600)
		CrashProgramIfError("open null device", err)
		Debug = log.New(nullDevFile, "", 0)
	}
}

func init() {
	if !regexp.MustCompile(versionRegExp).Match([]byte(Version)) {
		CrashProgram("invalid version")
	}
	setupDebugLogger()
}
