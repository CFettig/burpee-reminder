//deployed as cron job on AWS Lambda
package main

import (
  "encoding/json"
  "fmt"
  "math"
  "os"
  "strings"
  "time"
  "github.com/aws/aws-lambda-go/lambda"
//   "github.com/joho/godotenv"
  "github.com/twilio/twilio-go"
  twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// loadConfig("./.env")
	// Handler()
	lambda.Start(Handler)
}

func Handler() {
	recipients := strings.Split(os.Getenv("RECIPIENTS"), ",")
	sender := os.Getenv("SENDER")
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	
	for _, recipient := range recipients {
		message := buildMessage()
			sendMessage(client, sender, recipient, message)
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

func buildMessage() string {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timeElapsed := time.Since(time.Date(2023, 1, 1, 0, 0, 0, 0, location))
	num_burpees := int(math.Ceil(timeElapsed.Hours() / 24))

	var msg string

	if num_burpees == 100 {
		msg = fmt.Sprintf("Today is day 100!\nYou have done 5050 burpees since Jan 1\n%s\n[512/5050]", strings.Repeat("🦍", 511))
	} else {
		total_done := 0
		for i := 1; i <= num_burpees; i++ {
			total_done += i
		}

		msg =  fmt.Sprintf("We'll do %v burpees today\nMaking it %v burpees since Jan 1\nOnly %v left!\n", 
			num_burpees, total_done, 5050-total_done)
	}

	return msg;
}

// func loadConfig(path string) {
	// err := godotenv.Load(path)
	// if (err != nil) {
		// panic(err)
	// }
// }
