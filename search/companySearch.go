package search


import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sort"
	"net/http"
)

type SearchResults struct {
	Companies []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		IntelCount int    `json:"intelCount"`
	} `json:"companies"`
}


type Company struct {
	Name                     string `json:"name"`
	States            []interface{} `json:"states"`
	ProductAttributes struct {
		Muni struct {
			MunicipalsSectors []interface{} `json:"municipalsSectors"`
		} `json:"muni"`
		Debtwire struct {
			DominantSector struct {
				Code  string `json:"code"`
				Value string `json:"value"`
			} `json:"dominantSector"`
			DominantCountry struct {
				Code  string `json:"code"`
				Value string `json:"value"`
			} `json:"dominantCountry"`
		} `json:"debtwire"`
	} `json:"productAttributes"`
	Headquarters struct {
	} `json:"headquarters"`
	Identifiers      struct {
	} `json:"identifiers"`
	Sectors                  []struct {
		Code  string `json:"code"`
		Value string `json:"value"`
	} `json:"sectors"`
	Subsectors []struct {
		Code  string `json:"code"`
		Value string `json:"value"`
	} `json:"subsectors"`
	Countries []struct {
		Code  string `json:"code"`
		Value string `json:"value"`
	} `json:"countries"`
	Description    string        `json:"description"`
	Aliases        []interface{} `json:"aliases"`
	Mmgid          string        `json:"mmgid"`
	PublishingName string        `json:"publishingName"`
}


func SearchForCompanies(companyName string) {
	url:= fmt.Sprintf("https://aslive-intel-search-service.dev.mmgapi.net/search/issuer?q=%s&e=8_1,8_2,8_8&startFrom=0&pageSize=10", companyName)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var results SearchResults
	if err := json.Unmarshal(body, &results); err != nil {
		panic(err) // panic
	}

	//fmt.Printf("Got %+v", results)
	fmt.Println("set of results")
	sort.Slice(results.Companies, func(i, j int) bool { return results.Companies[i].IntelCount > results.Companies[j].IntelCount })
	for _, company := range results.Companies {
		fmt.Println()
		GetDetailsForCompany(company.ID)
		break
	}
}

func GetDetailsForCompany(id string) {

	url:= fmt.Sprintf("https://aslive-company-store.dev.mmgapi.net/company?mmgid=prime-%s", id)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var company Company
	if err := json.Unmarshal(body, &company); err != nil {
		panic(err)
	}

	fmt.Printf("Company was %s with Sector: %s, Country: %s", company.Name, company.ProductAttributes.Debtwire.DominantSector.Value, company.ProductAttributes.Debtwire.DominantCountry.Value)
	fmt.Println()
	fmt.Println(company.Description)

}
