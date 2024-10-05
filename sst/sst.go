package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type TimeResponse struct {
	DateTime string `json:"dateTime"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Day      int    `json:"day"`
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
}

func getTimeZone() (string, error) {
	// Read the timezone from /etc/timezone or use the 'date' command
	data, err := os.ReadFile("/etc/timezone")
	if err == nil {
		return strings.TrimSpace(string(data)), nil
	}

	// Fallback to using the 'date' command
	cmd := exec.Command("date", "+%Z")
	timezone, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("unable to get timezone: %v", err)
	}

	return strings.TrimSpace(string(timezone)), nil
}

func main() {
	timeZone, err := getTimeZone()
	if err != nil {
		fmt.Println("Error fetching timezone:", err)
		return
	}

	apiURL := fmt.Sprintf("https://timeapi.io/api/time/current/zone?timeZone=%s", timeZone)

	resp, err := http.Get(apiURL)
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

	// Set the hardware clock
	cmd = exec.Command("sudo", "hwclock", "--systohc")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error setting hardware clock:", err)
		return
	}

	fmt.Println("Time set successfully")
	period := "AM"
	if timeResponse.Hour >= 12 {
		period = "PM"
	}

	// convert 24-hour time to 12-hour time
	hour := timeResponse.Hour

	if hour > 12 {
		hour -= 12
	}

	fmt.Printf("%d %d %d %d:%02d %s\n", timeResponse.Day, timeResponse.Month, timeResponse.Year, hour, timeResponse.Minute, period)

}
