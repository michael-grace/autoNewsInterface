package main

import (
	"encoding/json"
	"fmt"
	"github.com/UniversityRadioYork/myradio-go"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type configAutoNews struct {
	TimeslotID    int  `json:"timeslotID"`
	AutoNewsStart bool `json:"autoNewsStart"`
	AutoNewsEnd   bool `json:"autoNewsEnd"`
}

type thonkyConfigBoi struct {
	APIKey           string           `json:"apiKey"`
	NewsOnJukebox    bool             `json:"newsOnJukebox"`
	OBShows          []int            `json:"obShows"`
	AutonewsRequests []configAutoNews `json:"autonewsRequests"`
}

type interfaceConfig struct {
	SwitcherConfigFilePath string `json:"switcherConfigFilePath"`
	APIKey                 string `json:"apiKey"`
	Port                   int    `json:"port"`
	TemplatePath           string `json:"templatePath"`
}

const (
	interfaceConfigFilePath = "interface_config.json"
)

func autonewsCheck(timeslotid int, config thonkyConfigBoi) [2]bool {
	for _, value := range config.AutonewsRequests {
		if value.TimeslotID == timeslotid {
			return [2]bool{value.AutoNewsStart, value.AutoNewsEnd}
		}
	}
	return [2]bool{true, true}
}

func main() {
	interfaceConfigFile, err := os.Open(interfaceConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer interfaceConfigFile.Close()
	byteValue, _ := ioutil.ReadAll(interfaceConfigFile)
	var interfaceConfig interfaceConfig
	json.Unmarshal(byteValue, &interfaceConfig)

	session, err := myradio.NewSession(interfaceConfig.APIKey)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		configFile, err := os.Open(interfaceConfig.SwitcherConfigFilePath)
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
		byteValue, _ := ioutil.ReadAll(configFile)
		var config thonkyConfigBoi
		json.Unmarshal(byteValue, &config)

		urlparams := r.URL.Query()
		urltimeslot, ok := urlparams["timeslotid"]
		if ok {
			timeslotint, _ := strconv.Atoi(urltimeslot[0])
			var autonews [2]bool = [2]bool{true, true}
			_, okS := urlparams["S"]
			_, okE := urlparams["E"]
			if !okS {
				autonews[0] = false
			}
			if !okE {
				autonews[1] = false
			}
			var complete bool
			for key, val := range config.AutonewsRequests {
				if val.TimeslotID == timeslotint {
					val.AutoNewsStart = autonews[0]
					val.AutoNewsEnd = autonews[1]
					complete = true
					config.AutonewsRequests[key] = val
				}
			}
			if !complete {
				config.AutonewsRequests = append(config.AutonewsRequests, configAutoNews{TimeslotID: timeslotint, AutoNewsStart: autonews[0], AutoNewsEnd: autonews[1]})
			}
			fmt.Printf("Timeslot ID: %v, AutoNews Start: %v, AutoNews End: %v\n", timeslotint, autonews[0], autonews[1])
		}

		year, week := time.Now().ISOWeek()
		shows, err := session.GetWeekSchedule(year, week)
		if err != nil {
			panic(err)
		}

		type timeslot struct {
			Title         string
			Timeslotid    uint64
			AutoNewsStart bool
			AutoNewsEnd   bool
		}

		var timeslots []timeslot

		var weekday int = int(time.Now().Weekday())
		if weekday == 0 {
			weekday = 7
		}

		for _, show := range shows[weekday] {
			anews := autonewsCheck(int(show.TimeslotID), config)
			timeslots = append(timeslots, timeslot{
				Title:         show.Title,
				Timeslotid:    show.TimeslotID,
				AutoNewsStart: anews[0],
				AutoNewsEnd:   anews[1],
			})
		}

		data := struct {
			Timeslots  []timeslot
			DataDriven bool
		}{
			Timeslots:  timeslots,
			DataDriven: ok,
		}

		tmpl, err := template.ParseFiles(interfaceConfig.TemplatePath)
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, data)

		file, err := os.Create(interfaceConfig.SwitcherConfigFilePath)

		if err != nil {
			return
		}
		defer file.Close()

		newConfig, _ := json.Marshal(config)
		file.WriteString(string(newConfig))
	})
	http.ListenAndServe(fmt.Sprintf(":%v", interfaceConfig.Port), nil)
}
