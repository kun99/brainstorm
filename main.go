package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/rs/cors"
)

type ResponseData struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Word []struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
		License   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"license,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string        `json:"definition"`
			Synonyms   []interface{} `json:"synonyms"`
			Antonyms   []interface{} `json:"antonyms"`
			Example    string        `json:"example,omitempty"`
		} `json:"definitions"`
		Synonyms []interface{} `json:"synonyms"`
		Antonyms []interface{} `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}

func brainstorm() string {
	cmd := exec.Command("python", "brainstorm.py")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(output)
}

func getDefinition(randomWord string) string {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + randomWord
	resp, err := http.Get(strings.TrimSpace(url))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var word Word
	json.NewDecoder(resp.Body).Decode(&word)
	if len(word) > 0 {
		definition := word[0].Meanings[0].Definitions[0].Definition
		return definition
	} else {
		return "---"
	}
}

func main() {
	//comments are for me learning purposes
	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		// check for request method. if it aint a GET request we dont want it
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		randomWord := brainstorm()
		definition := getDefinition(randomWord)
		responseData := ResponseData{
			Word:       randomWord,
			Definition: definition,
		}

		// attempting to turn our custom struct into json
		jsonResponse, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// writing json to the http response writer
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	c := cors.Default()

	handler := c.Handler(http.DefaultServeMux)

	http.ListenAndServe(":8080", handler)
}
