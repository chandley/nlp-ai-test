package main

import (
	"fmt"
	"log"
	// Imports the Google Cloud Natural Language API client package.
	language "cloud.google.com/go/language/apiv1"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"github.com/chandley/nlp-ai-test/search"
)

const appleStory = "Apple has been linked with a shock £1.5bn deal to buy McLaren Technology Group, the Formula One team owner and supercar maker. A deal" +
	" between Apple and the British company would dramatically shake up the technology and automotive industries. The California-based company’s " +
		"interest in McLaren Technology Group highlights its ambition to develop technology that could be used in an electric and driverless car."

const ryanairStory = "Ryanair is facing enforcement action from the Civil Aviation Authority for persistently misleading passengers" +
" about their rights, piling more woe on the no-frills carrier as it announced a second wave of flight cancellations that will " +
"affect 400,000 people. In a letter to the Dublin-based airline, the CAA said chief executive Michael O’Leary was wrong to " +
"tell passengers last week that it did not have to arrange new flights for them after an initial batch of cancellations " +
"were announced. The airline regulator said Ryanair had further transgressed when it announced fresh disruption on Wednesday, " +
"by failing to tell passengers that they could be rerouted with other airlines if there was no suitable alternative on one of its own planes."

const boeingStory = "The government has warned aircraft manufacturer Boeing it could lose UK defence contracts over its part in a US decision to " +
"slap punitive tariffs of 219% on rival Bombardier, in a dispute that threatens to sour trade relations between London and Washington. " +
"Theresa May said she was “bitterly disappointed” by the move to impose a tariff on sales of Bombardier’s C-Series passenger jet, which threatens " +
"at least 1,000 manufacturing jobs in Northern Ireland. Michael Fallon, the UK defence secretary, stepped up the government’s rhetoric, warning " +
"that Boeing’s assault on Bombardier “could jeopardise” its chances of securing government contracts. The business secretary, Greg Clark, joined " +
"the chorus of disapproval, branding the ruling “unjustified” and vowing to work with Canada – where Bombardier is based – to get it overturned."

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := boeingStory

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
	//fmt.Printf("Entites detected: %v", response)
	const SALIENCE_THRESHOLD = 0.05
	fmt.Println()
	for i, entity := range response.Entities {
		if entity.Type == languagepb.Entity_ORGANIZATION && entity.Salience > SALIENCE_THRESHOLD{
			fmt.Printf("Entity %s: %+v", i, entity.Name)
			fmt.Println(" ")
			fmt.Println(" ")
			search.SearchForCompanies(entity.Name)
			//break
			//fmt.Println()
		}
	}
}






