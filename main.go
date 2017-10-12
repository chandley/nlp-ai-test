package main

import (
	"fmt"
	"log"
	// Imports the Google Cloud Natural Language API client package.
	language "cloud.google.com/go/language/apiv1"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"github.com/chandley/nlp-ai-test/search"
  "net/http"
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
  http.HandleFunc("/", inputHandler)
  http.HandleFunc("/save", saveHandler)
  http.ListenAndServe(":8080", nil)
}

func inputHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Paste your story</h1>"+
        "<form action=\"/save\" method=\"POST\">"+
        "<textarea style=\"height: 200px; width: 500px;\" name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        "Input text here")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    body := r.FormValue("body")
    fmt.Printf(analyseStory(body))
    http.Redirect(w, r, "/", http.StatusFound)
}

func analyseStory(story string) string {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Detects the sentiment of the story.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: story,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze story: %v", err)
	}

  returnString := ""

	returnString += fmt.Sprintf("text: %v\n", story)
	if sentiment.DocumentSentiment.Score >= 0 {
		returnString += fmt.Sprintf("Sentiment: positive, score: %v", sentiment.DocumentSentiment.Score)
	} else {
		returnString += fmt.Sprintf("Sentiment: negative, score: %v", sentiment.DocumentSentiment.Score)
	}
	returnString += "\n"

	response, err := client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: story,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze entities: %v", err)
	}
	const SALIENCE_THRESHOLD = 0.05
	returnString += "\n"

	for i, entity := range response.Entities {
		if entity.Type == languagepb.Entity_ORGANIZATION && entity.Salience > SALIENCE_THRESHOLD{
			returnString += fmt.Sprintf("Entity %s: %+v", i, entity.Name)
      returnString += "\n"
      returnString += "\n"
			returnString += fmt.Sprintf(search.SearchForCompanies(entity.Name))
		}
	}

  return returnString
}




