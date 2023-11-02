package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
)

type SteamData struct {
	SteamKey string `yaml:"SteamKey"`
}

func GetNumbersId(steamID string) {
	client := &http.Client{}
	configData := &SteamData{}
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, configData)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	link := "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + configData.SteamKey + "&steamids=" + steamID
	req, err := http.NewRequest("GET", link, nil)

	resp, err := client.Do(req)

	ff, _ := io.ReadAll(resp.Body)
	//err := json.Unmarshal(ff, &iot)
	fmt.Println(string(ff))
}
