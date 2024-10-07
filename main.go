package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {

	var userName string

	// Regular expression pattern
	pattern := `^[\p{L}]+$`
	re := regexp.MustCompile(pattern)

	// Prompt user for input
	fmt.Print("github-activity: ")
	_, err := fmt.Fscan(os.Stdin, &userName)
	if err != nil {
		fmt.Println(err)
	}
	if re.MatchString(userName) {
		ExecuteQuery(userName)
	}

}

func ExecuteQuery(userName string) {

	var resp *http.Response

	resp, err := http.Get("https://api.github.com/users/" + userName + "/events")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	FormatStringAnswer(body)
}

func FormatStringAnswer(body []byte) {

	BodyString := string(body)
	PushedCount := strings.Count(BodyString, "PushEvent")
	CreateCount := strings.Count(BodyString, "CreateEvent")
	DeleteCount := strings.Count(BodyString, "DeleteEvent")
	WatchCount := strings.Count(BodyString, "WatchEvent")

	finalString := fmt.Sprintf("Pushes on GitHub: %d \n Creates on GitHub: %d \n Deletes on GitHub: %d \n Watches on GitHub: %d",
		PushedCount, CreateCount, DeleteCount, WatchCount)
	println(finalString)
}
