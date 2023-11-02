package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type browserData struct {
	UserAgent string `yaml:"UserAgent"`
	Cookie    string `yaml:"Cookie"`
	AccessKey string `yaml:"AccessKey"`
}

func friendsDownload(steamID string) {

	configData := &browserData{}
	client := &http.Client{}
	out, err := os.Create(steamID)
	if err != nil {
		fmt.Println("Error with create file: ", err)
	}
	defer out.Close()

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, configData)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	link := "https://steamid.uk/export.php?accesskey=" + configData.AccessKey + "&task=exportfriends&user=" + steamID

	req, err := http.NewRequest("GET", link, nil)
	req.Header.Set("User-Agent", configData.UserAgent)
	req.Header.Set("Cookie", configData.Cookie)
	resp, err := client.Do(req)

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Copy from browser: %v", err)
	}
	fmt.Println("Successful")
}

func arrayFriendsFunc(steamID string) []string {
	var arrayFriends []string
	file, _ := os.Open(steamID)
	temp, _ := io.ReadAll(file)
	readFile := string(temp)
	lines := strings.Split(readFile, "\n")
	for _, v := range lines {

		vv := strings.Split(v, ",")
		link := vv[len(vv)-1]
		if strings.HasPrefix(link, "https") {
			arrayFriends = append(arrayFriends, link)
		}

	}

	return arrayFriends

}
