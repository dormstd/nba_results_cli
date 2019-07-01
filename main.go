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

func main() {

	now := time.Now().Format("01/02/2006")
	date := flag.String("date", now, "set the date in MM/DD/YYYY format")

	flag.Parse()
	fmt.Printf("Starting the application with date: %v...\n", *date)
	var netClient = &http.Client{
		Timeout: time.Second * 5,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", "https://stats.nba.com/stats/scoreboard/?GameDate="+*date+"&LeagueID=00&DayOffset=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	/** This headers are mandatory for the request to work, stats.nba thing...*/
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	request.Header.Set("origin", "http://stats.nba.com")
	request.Header.Set("Dnt", "1")
	request.Header.Set("Accept-Encoding", "'gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "en")

	response, err := netClient.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("response not OK from webservice")
	}
	data, _ := ioutil.ReadAll(response.Body)

	resStructure := scoreBoardResult{}
	if err := json.Unmarshal(data, &resStructure); err != nil {
		panic(err)
	}
	games, err := newGames(resStructure)

	if err != nil {
		panic(err)
	}

	printResults(games)
}
