package inputs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetInput(year int, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	session := os.Getenv("AOC_SESSION_COOKIE")
	req.Header.Set("Cookie", "session="+session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get input: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
