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
type Cargo struct {
	Timestamp string      `json:"timestamp"`
	Event     string      `json:"event"`
	Count     int32       `json:"Count"`
	Inventory []Inventory `json:"Inventory"`
}

type Inventory struct {
	Name_Localised string `json:"Name_Localised"`
	MissionID      int32  `json:"MissionID"`
	Count          int32  `json:"Count"`
}

// GetStatus reads the current player and ship status from Status.json.
// It will read them from the default log path, which is the Saved Games
// folder. The full path is:
//
//	C:/Users/<Username>/Saved Games/Frontier Developments/Elite Dangerous
//
// If that path is not suitable, use GetStatusFromPath.
func GetCargo() (*Cargo, error) {
	return GetCargoFromPath(defaultLogPath)
}

func GetCargoFromPath(logPath string) (*Cargo, error) {
	filePath := filepath.FromSlash(logPath + "/Cargo.json")
	retries := 5

	var cargo *Cargo
	for retries > 0 {
		file, err := os.Open(filePath)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}

		return GetCargoFromBytes(bytes)
	}
	return cargo, nil
}

// GetStatusFromBytes reads the current player and ship status from the string contained in the byte array.
func GetCargoFromBytes(content []byte) (*Cargo, error) {
	cargo := &Cargo{}
	if err := json.Unmarshal(content, cargo); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	return cargo, nil
}
