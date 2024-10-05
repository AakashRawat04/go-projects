package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type TimeResponse struct {
	DateTime string `json:"dateTime"`
}

func main() {
	resp, err := http.Get("https://timeapi.io/api/time/current/zone?timeZone=Asia%2FKolkata")
	if err != nil {
		fmt.Println("Error fetching time:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		return
	}

	var timeResponse TimeResponse
	if err := json.NewDecoder(resp.Body).Decode(&timeResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Set the system time (Unix timestamp is used here)
	cmd := exec.Command("sudo", "date", "-s", timeResponse.DateTime)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error setting time:", err)
		return
	}

	fmt.Println("System time set to:", timeResponse.DateTime)
}
