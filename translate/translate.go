package translate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// GasResponse is 'Google App Script' response.
type GasResponse struct {
	TranslatedText string `json:"translated_text"`
}

// Translate function
func Translate(t string) string {
	endpoint := GenerateEndpoint(t)

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(GasResponse)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Fatal(err)
	}

	return data.TranslatedText
}

// GenerateEndpoint function
func GenerateEndpoint(words string) string {
	var sourceAndTarget string = "&source=ja&target=en"
	// en => ja
	r := regexp.MustCompile(`^#en`)
	if r.MatchString(words) {
		words = r.ReplaceAllString(words, "")
		sourceAndTarget = "&source=en&target=ja"
	}
	words = url.QueryEscape(words)
	bassURL := os.Getenv("API_BASSURL")
	return bassURL + "/exec?text=" + words + sourceAndTarget
}
