package ex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func MeatSummaryHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	text := string(body)
	beefCount := countBeefTypes(text)

	response := map[string]interface{}{
		"beef": beefCount,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func countBeefTypes(text string) map[string]int {
	beefTypes := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}
	beefCount := make(map[string]int)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	for _, word := range words {
		for _, beef := range beefTypes {
			if strings.ToLower(word) == beef {
				beefCount[beef]++
			}
		}
	}
	return beefCount
}
