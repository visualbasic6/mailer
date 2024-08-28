package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gosuri/uilive"
	"github.com/mailgun/mailgun-go/v4"
	"golang.org/x/net/context"
)

func main() {
	// Load subject and body from files
	subject, err := ioutil.ReadFile("subject.txt")
	if err != nil {
		log.Fatalf("Failed to read subject.txt: %v", err)
	}

	body, err := ioutil.ReadFile("body.html")
	if err != nil {
		log.Fatalf("Failed to read body.html: %v", err)
	}

	// Open the recipients list
	file, err := os.Open("recipients.txt")
	if err != nil {
		log.Fatalf("Failed to open recipients.txt: %v", err)
	}
	defer file.Close()

	// Initialize Mailgun
	domain := "your-mailgun-domain.com"
	apiKey := "your-api-key-here"
	mg := mailgun.NewMailgun(domain, apiKey)

	// Set up a live writer to update the terminal UI
	writer := uilive.New()
	writer.Start()

	scanner := bufio.NewScanner(file)
	sentCount := 0
	startTime := time.Now()

	for scanner.Scan() {
		email := strings.TrimSpace(scanner.Text())
		message := mg.NewMessage(
			"you@yourdomain.com",                // Sender
			string(subject),                     // Subject
			"",                                  // Plaintext body
			email,                               // Recipient
		)
		message.SetHtml(string(body))

		// Send the email
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		_, _, err := mg.Send(ctx, message)
		if err != nil {
			fmt.Fprintf(writer, "failed to send to %s: %v\n", email, err)
		} else {
			sentCount++
			timestamp := time.Now().Format("15:04:05")
			fmt.Fprintf(writer, "[%s] sent to: %s | total sent: %d\n", timestamp, email, sentCount)
		}
		time.Sleep(2100 * time.Millisecond) // Sleeps for 2.1 seconds
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read recipients list: %v", err)
	}

	writer.Stop() // Stop the live writer
	elapsedTime := time.Since(startTime)
	fmt.Printf("Finished sending emails. Total sent: %d. Time elapsed: %s\n", sentCount, elapsedTime)
}
