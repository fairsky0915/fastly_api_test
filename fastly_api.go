package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var defaultURLforAPI string
var serviceID string
var tokenID string
var serviceSelect int
var taskSelect int
var apiURL string
var versionNo int
var serviceName string
var domainName string
var updatedDomainName string

func main() {
	fmt.Println("Please, select the task of service you want to do.\n",
		"1. Version\n 2. Service\n 3. Domain")
	fmt.Scanln(&serviceSelect)
	switch serviceSelect {
	case 1:
		versionAPI()
	case 2:
		serviceAPI()
	case 3:
		domainAPI()
	}

}

func domainAPI() {

	fmt.Println("Please, input your token ID")
	fmt.Scanln(&tokenID)

	fmt.Println("Please, select the task you want to do.\n",
		"1. Add a domain name to a serivce\n 2. Remove a domain from a service\n 3. Update a domain")
	fmt.Scanln(&taskSelect)

	defaultURLforAPI = "https://api.fastly.com/service"

	switch taskSelect {
	case 1: /*Add a domain name to a serivce
		https://api.fastly.com/service/[service ID]/version/13/domain?name=www.songtest2.com */

		fmt.Println("Please, input your service ID you want to work")
		fmt.Scanln(&serviceID)
		fmt.Println("Please, input the version number you want to work")
		fmt.Scanln(&versionNo)
		fmt.Println("Please, input the domain name you want to make")
		fmt.Scanln(&domainName)

		apiURL = fmt.Sprintf("%s/%s/version/%d/domain?name=%s", defaultURLforAPI, serviceID, versionNo, domainName)
		req, err := http.NewRequest("POST", apiURL, nil)
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

		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(body))
		fmt.Println("request:", req)

	case 2: /* Remove a domain from a service
		'https://api.fastly.com/service/[Service_ID]/version/13/domain/www.songtest2.com */
		fmt.Println("Please, input your service ID you want to work")
		fmt.Scanln(&serviceID)
		fmt.Println("Please, input the version number you want to work")
		fmt.Scanln(&versionNo)
		fmt.Println("Please, input the domain name you want to delete")
		fmt.Scanln(&domainName)

		apiURL = fmt.Sprintf("%s/%s/version/%d/domain/%s", defaultURLforAPI, serviceID, versionNo, domainName)
		req, err := http.NewRequest("DELETE", apiURL, nil)
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

		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(body))
		fmt.Println("request:", req)
	case 3: /* update a domain
		'https://api.fastly.com/service/[Service_ID]/version/13/domain/old_domain_name?name=newdomain_name */
		fmt.Println("Please, input your service ID you want to work")
		fmt.Scanln(&serviceID)
		fmt.Println("Please, input the version number you want to work")
		fmt.Scanln(&versionNo)
		fmt.Println("Please, input the exist domain name you want to update")
		fmt.Scanln(&domainName)
		fmt.Println("Please, input the new domain name")
		fmt.Scanln(&updatedDomainName)

		apiURL = fmt.Sprintf("%s/%s/version/%d/domain/%s?name=%s", defaultURLforAPI, serviceID, versionNo, domainName, updatedDomainName)
		req, err := http.NewRequest("PUT", apiURL, nil)
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

		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(body))
		fmt.Println("request:", req)
	}
}

func serviceAPI() {

	fmt.Println("Please, input your token ID")
	fmt.Scanln(&tokenID)

	fmt.Println("Please, select the task you want to do.\n",
		"1. Create a new Service\n 2. Update a service\n 3. Delete a service")
	fmt.Scanln(&taskSelect)

	defaultURLforAPI = "https://api.fastly.com/service"

	switch taskSelect {
	case 1: //Create a new service https://api.fastly.com/service?name=new-song-apitest

		fmt.Println("Please, input the serviceName you want to make")
		fmt.Scanln(&serviceName)
		apiURL = fmt.Sprintf("%s?name=%s", defaultURLforAPI, serviceName)
		fmt.Println(apiURL)
		req, err := http.NewRequest("POST", apiURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Fastly-Key", tokenID)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(body))
		fmt.Println("request:", req)

	case 2: //update a service https://api.fastly.com/service/[Service_ID]?name=updated-service-name"
		fmt.Println("Please, input your service ID you want to update")
		fmt.Scanln(&serviceID)
		fmt.Println("Please, input the serviceName you want to be updated")
		fmt.Scanln(&serviceName)
		apiURL = fmt.Sprintf("%s/%s?name=%s", defaultURLforAPI, serviceID, serviceName)
		fmt.Println(apiURL)
		req, err := http.NewRequest("PUT", apiURL, nil)
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
	case 3: //delete a serivce https://api.fastly.com/service/[Service_ID]
		fmt.Println("Please, input your service ID you want to delete")
		fmt.Scanln(&serviceID)

		apiURL = fmt.Sprintf("%s/%s", defaultURLforAPI, serviceID)
		// Create client
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", apiURL, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Fastly-Key", tokenID)
		req.Header.Set("Accept", "application/json")

		//resp, err := http.DefaultClient.Do(req)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//response status
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(body))

	}
}

func versionAPI() {

	fmt.Println("Please, input your service ID")
	fmt.Scanln(&serviceID)
	fmt.Println("Please, input your token ID")
	fmt.Scanln(&tokenID)

	fmt.Println("Please, select the task you want to do.\n",
		"1. Create a new version\n 2. Clone an existing version\n 3. Activate a version")
	fmt.Scanln(&taskSelect)

	defaultURLforAPI = "https://api.fastly.com/service"

	switch taskSelect {
	case 1: //Create a new version https://api.fastly.com/service/[service ID]/version
		apiURL = fmt.Sprintf("%s/%s/version", defaultURLforAPI, serviceID)
		req, err := http.NewRequest("POST", apiURL, nil)
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
	case 2: //Clone an existing version https://api.fastly.com/service/[Service_ID]/version/2/clone
		fmt.Println("Please, input the Version Number you want to clone")
		fmt.Scanln(&versionNo)
		apiURL = fmt.Sprintf("%s/%s/version/%d/clone", defaultURLforAPI, serviceID, versionNo)
		fmt.Println(apiURL)

		req, err := http.NewRequest("PUT", apiURL, nil)
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
	case 3: //Activate a version https://api.fastly.com/service/[Service ID]/version/2/activate
		fmt.Println("Please, input the Version Number you want to Activate")
		fmt.Scanln(&versionNo)
		apiURL = fmt.Sprintf("%s/%s/version/%d/activate", defaultURLforAPI, serviceID, versionNo)
		fmt.Println(apiURL)
		req, err := http.NewRequest("PUT", apiURL, nil)
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
