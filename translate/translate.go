package translate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// GasResponse is 'Google App Script' response.
type GasResponse struct {
	TranslatedText string `json:"translated_text"`
}

// Translate function
func Translate(t string) string {
	text := url.QueryEscape(t)

	bassURL := os.Getenv("API_BASSURL")
	endpoint := bassURL + "/exec?text=" + text + "&source=ja&target=en"

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
