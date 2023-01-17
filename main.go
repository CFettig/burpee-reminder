package main

import (
  "encoding/json"
  "fmt"
  "math"
  "os"
  "strings"
  "time"
  "github.com/joho/godotenv"
  "github.com/twilio/twilio-go"
  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	loadConfig("./.env")

	//list of numbers to send reminder to
	sender := os.Getenv("SENDER")
	recipients := strings.Split(os.Getenv("RECIPIENTS"), ",")

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// c := make(chan bool)
	// go messageTimer(c)

	for _, recipient := range recipients {   
		sendMessage(client, sender, recipient, createMessage())
	}
}

func messageTimer(c chan bool) {
	for true {
		time.Sleep(5 * time.Second)
	}
}

func sendMessage(client *twilio.RestClient, sender, recipient, message string) {
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(sender)
	params.SetTo(recipient)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func createMessage() string {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timeElapsed := time.Since(time.Date(2023, 1, 1, 0, 0, 0, 0, location))
	num_burpees := int(math.Ceil(timeElapsed.Hours() / 24))

	total_done := 0
	for i := 1; i <= num_burpees; i++ {
		total_done += i
	}

	return fmt.Sprintf("Today is day %v :)\nToday will be %v burpees\nOnly %v burpees left!\n", 
		num_burpees, total_done, 5050-total_done)
}

func loadConfig(path string) {
	err := godotenv.Load(path)
	if (err != nil) {
		panic(err)
	}
}
