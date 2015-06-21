package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type msgs struct {
	Status     string
	Body       string
	Created_on string
}

var (
	jsonOut         bool
	lastHumanComm   bool
	recentHumanComm bool
)

func init() {
	flag.BoolVar(&jsonOut, "j", false, "Output responses as JSON.")

	flag.BoolVar(&lastHumanComm, "last-human", false, "Returns the last human "+
		"communication, status, and timestamp.")

	flag.BoolVar(&recentHumanComm, "recent-human", false, "Returns the most "+
		"recent human communications with status and timestamp.")

	flag.Parse()
}

func main() {
	data, err := performRequest("https://status.github.com/api.json")
	checkError(err)
	j, err := extractJSONObj(data)
	checkError(err)

	last_msg_url := j["last_message_url"].(string)
	msgs_url := j["messages_url"].(string)
	status_url := j["status_url"].(string)

	if lastHumanComm {
		data, err = performRequest(last_msg_url)
		checkError(err)
		if jsonOut {
			fmt.Println(string(data))
		} else {
			j, err = extractJSONObj(data)
			checkError(err)
			fmt.Println("Status:", j["status"].(string))
			fmt.Println("Message:", j["body"].(string))
			t, err := time.Parse(time.RFC3339, j["created_on"].(string))
			checkError(err)
			fmt.Println("Created on:", t)
		}
	} else if recentHumanComm {
		data, err = performRequest(msgs_url)
		checkError(err)
		if jsonOut {
			fmt.Println(string(data))
		} else {
			m, err := extractJSONArray(data)
			checkError(err)
			if len(m) == 0 {
				fmt.Println("No recent human communications.")
			} else {
				for i := 0; i < len(m); i++ {
					fmt.Println("Status", m[i].Status)
					fmt.Println("Message:", m[i].Body)
					t, err := time.Parse(time.RFC3339, m[i].Created_on)
					checkError(err)
					fmt.Println("Created on:", t)
					fmt.Printf("\n")
				}
			}
		}
	} else {
		data, err = performRequest(status_url)
		checkError(err)
		if jsonOut {
			fmt.Println(string(data))
		} else {
			j, err = extractJSONObj(data)
			checkError(err)
			fmt.Println("Status:", j["status"].(string))
			t, err := time.Parse(time.RFC3339, j["last_updated"].(string))
			checkError(err)
			fmt.Println("Last updated:", t)
		}
	}
}

// performRequest performs a HTTP Get and returns the associated JSON data.
func performRequest(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil
}

// extractJSONObj extracts and returns a JSON structure.
func extractJSONObj(jsonBlob []byte) (map[string]interface{}, error) {
	var j map[string]interface{}

	err := json.Unmarshal(jsonBlob, &j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// extractJSONArray unpacks and returns array of msgs from a JSON structure.
func extractJSONArray(jsonBlob []byte) ([]msgs, error) {
	var m []msgs

	err := json.Unmarshal(jsonBlob, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
