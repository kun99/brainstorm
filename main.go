package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"fmt"
)

type ResponseData struct {
	Word string `json:"word"`
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

func main() {
	//comments are for me learning purposes
	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		// check for request method. if it aint a GET request we dont want it
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		randomWord := brainstorm()		
		responseData := ResponseData{
			Word: randomWord,
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

	http.ListenAndServe(":8080", nil)
}
