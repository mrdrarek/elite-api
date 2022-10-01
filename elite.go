// Package elite provides real-time data from Elite Dangerous through
// files written to disk by the game.
package elite

import (
	"fmt"
	"path/filepath"
	"regexp"
)

// JournalEntry is a minimal entry in the Journal file.
// It is primarily intended for embedding within event types,
// such as StarSystemEvent.
type JournalEntry struct {
	Timestamp string `json:"timestamp"`
	Event     string `json:"event"`
}

var defaultLogPath string
var journalFilePattern *regexp.Regexp

func init() {

	journalFilePattern = regexp.MustCompile(`^Journal\.\d{4}\-\d{2}\-\d{2}`)
	fmt.Println("init Elite")
}

func SetLogPath(logPath string) {
	defaultLogPath = filepath.FromSlash(logPath)
}
