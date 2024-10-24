package ex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MaxPathHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("files/hard.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file: %v", err), http.StatusInternalServerError)
		return
	}

	var pyramid [][]int
	err = json.Unmarshal(data, &pyramid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal JSON: %v", err), http.StatusInternalServerError)
		return
	}

	result := maxPathSum(pyramid)

	response := map[string]interface{}{
		"max_path_sum": result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func maxPathSum(pyramid [][]int) int {
	for row := len(pyramid) - 2; row >= 0; row-- {
		for col := 0; col < len(pyramid[row]); col++ {
			pyramid[row][col] += max(pyramid[row+1][col], pyramid[row+1][col+1])
		}
	}
	return pyramid[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
