package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// type FSDJumpKeys struct {
	// 	timestamp  string  `key: "timestamp"`
	// 	StarSystem string  `key: "StarSystem"`
	// 	JumpDist   float32 `key: "JumpDist"`
	// 	FuelUsed   float32 `key: "FuelUsed"`
	// 	FuelLevel  string  `key: "FuelLevel"`
	// 	BoostUsed  string  `key: "BoostUsed"`
	// 	event      string  `key: "event"`
	// }

	// type FSDJumpInfo map[string]FSDJumpKeys

	configFile, err := ioutil.ReadFile("./Journal.200815091611.01.log")

	if err != nil {
		log.Fatal(err)
	}

	configLines := strings.Split(string(configFile), "\n")
	// fmt.Println(configLines[24])

	var parsedData map[string]interface{}
	json.Unmarshal([]byte(configLines[24]), &parsedData)

	for key, value := range parsedData {
		//fmt.Println("key:", key, "value:", value)
		fmt.Println(key, ": ", value)
	}

	// for i := 0; i < len(configLines); i++ {
	// 	fmt.Println(configLines[i])

	// 	// if configLines[i] != "" {

	// 	// 	configLine := strings.Split(string(configLines[i]), " ")

	// 	// 	newConfig := Config{Server: configLine[0], Username: configLine[1], Key: configLine[2]}
	// 	// 	configs = append(configs, newConfig)
	// 	// }
	// }

	// for _, config := range configs {
	// 	println(config.Server + " " + config.Username + " " + config.Key)
	// }
}
