package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	var serviceID string
	var tokenID string
	var taskSelect int
	var shortURL string
	var APIURLStr string
	var Version int
	var CloneURL string
	var ActivateURL string
	fmt.Println("Please, input your service ID")
	fmt.Scanln(&serviceID)
	fmt.Println("Please, input your token ID")
	fmt.Scanln(&tokenID)

	fmt.Println("Please, select the task you want to do.\n",
		"1. Create a new version\n 2. Clone an existing version\n 3. Activate a version")
	fmt.Scanln(&taskSelect)

	shortURL = "https://api.fastly.com/service/"
	APIURLStr = fmt.Sprintf("%s%s/version", shortURL, serviceID)

	switch taskSelect {
	case 1:
		req, err := http.NewRequest("POST", APIURLStr, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Fastly-Key", tokenID)
		req.Header.Set("Accept", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))

		//reponse status
		fmt.Println(string(resp.Status))
	case 2:
		fmt.Println("Please, input the Version you want to clone")
		fmt.Scanln(&Version)
		CloneURL = fmt.Sprintf("%s/%d/clone", APIURLStr, Version)
		req, err := http.NewRequest("PUT", CloneURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Fastly-Key", tokenID)
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))

		//reponse status
		fmt.Println(string(resp.Status))
	case 3:
		fmt.Println("Please, input the Version you want to Activate")
		fmt.Scanln(&Version)
		ActivateURL = fmt.Sprintf("%s/%d/activate", APIURLStr, Version)
		req, err := http.NewRequest("PUT", ActivateURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Fastly-Key", tokenID)
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))

		//reponse status
		fmt.Println(string(resp.Status))
	}
}
