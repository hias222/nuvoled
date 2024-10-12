package sendclock

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

var hiit_data Hiit
var config_filename = "hiitparameter.json"

type Hiit struct {
	Lane []Lane `json:"hiit"`
}

type Lane struct {
	Runtime int   `json:"runtime"`
	Times   []int `json:"times"`
}

func InitParam() {

	exist := checkFileExists(config_filename)

	if !exist {
		logger.Warn("File " + config_filename + " not exist")
	}

	plan, _ := os.ReadFile(config_filename)
	//var data map[string]interface{}
	var data Hiit

	err := json.Unmarshal(plan, &data)

	if err != nil {

		logger.Warn("Json File read failed:  " + err.Error())

		myLane := Lane{Runtime: 55, Times: []int{35, 35, 35}}
		lane_array := []Lane{myLane}
		hiit := Hiit{lane_array}
		hiit_data = hiit

		printHiitData(hiit)

		return
	}

	hiit_data = data

	printHiitData(data)

}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func printHiitData(data Hiit) {
	for i := 0; i < len(data.Lane); i++ {
		logger.Info(strconv.Itoa(i) + ": Runtime: " + strconv.Itoa(data.Lane[i].Runtime))
		for g := 0; g < len(data.Lane[i].Times); g++ {
			logger.Info("- " + strconv.Itoa(data.Lane[i].Times[g]))
		}
	}
}
