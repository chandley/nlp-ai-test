// Sample language-quickstart uses the Google Cloud Natural API to analyze the
// sentiment of "Hello, world!".
package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud Natural Language API client package.
	language "cloud.google.com/go/language/apiv1"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := "We all love James Sabben-Clare"

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Printf("Sentiment: positive, score: %v", sentiment.DocumentSentiment.Score)
	} else {
		fmt.Printf("Sentiment: negative, score: %v", sentiment.DocumentSentiment.Score)
	}
	fmt.Println()
}
