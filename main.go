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

const ryanairStory = "Ryanair is facing enforcement action from the Civil Aviation Authority for persistently misleading passengers" +
" about their rights, piling more woe on the no-frills carrier as it announced a second wave of flight cancellations that will " +
"affect 400,000 people. In a letter to the Dublin-based airline, the CAA said chief executive Michael O’Leary was wrong to " +
"tell passengers last week that it did not have to arrange new flights for them after an initial batch of cancellations " +
"were announced. The airline regulator said Ryanair had further transgressed when it announced fresh disruption on Wednesday, " +
"by failing to tell passengers that they could be rerouted with other airlines if there was no suitable alternative on one of its own planes."

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := ryanairStory

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

	response, err := client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze entities: %v", err)
	}
	fmt.Printf("Entites detected: %v", response)
	fmt.Println()
	for i, entity := range response.Entities {
		fmt.Printf("Entity %v: %+v", i, entity)
		fmt.Println()
	}
}


