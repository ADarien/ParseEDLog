package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// type Config struct {
	// 	Server   string
	// 	Username string
	// 	Key      string
	// }

	// configs := make([]Config, 0)

	configFile, err := ioutil.ReadFile("./Journal.200815091611.01.log")

	if err != nil {
		log.Fatal(err)
	}

	configLines := strings.Split(string(configFile), "\n")
	fmt.Println(configLines[1])
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
