package elite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Status represents the current state of the player and ship.
type Nav struct {
	Timestamp string  `json:"timestamp"`
	Event     string  `json:"event"`
	Route     []Route `json:"Route"`
}

type Route struct {
	StarSystem string `json:"StarSystem"`
	StarClass  string `json:"StarClass"`
}

// GetStatus reads the current player and ship status from Status.json.
// It will read them from the default log path, which is the Saved Games
// folder. The full path is:
//
//	C:/Users/<Username>/Saved Games/Frontier Developments/Elite Dangerous
//
// If that path is not suitable, use GetStatusFromPath.
func GetNavRoute() (*Nav, error) {
	return GetNavRouteFromPath(defaultLogPath)
}

func GetNavRouteFromPath(logPath string) (*Nav, error) {
	navsFilePath := filepath.FromSlash(logPath + "/NavRoute.json")
	retries := 5

	var nav *Nav
	for retries > 0 {
		navFile, err := os.Open(navsFilePath)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}
		defer navFile.Close()

		navBytes, err := ioutil.ReadAll(navFile)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}

		return GetNavFromBytes(navBytes)
	}
	return nav, nil
}

// GetStatusFromBytes reads the current player and ship status from the string contained in the byte array.
func GetNavFromBytes(content []byte) (*Nav, error) {
	status := &Nav{}
	if err := json.Unmarshal(content, status); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	return status, nil
}
