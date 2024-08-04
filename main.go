package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Result []int `json:"result"`
}

// Function to return array of packs
// Can remove tuple - error
func orderPackSize(x int, arr []int) []int {

	//Create empty array
	result := []int{}

	//Incase User doesn't input in ascending order
	//Improvements - check for duplicates in array
	sort.Ints(arr)

	remaining := x
	//Iterate through packs starting at the largest see if remaining figure is larger that the pack size in the array
	for i := len(arr) - 1; i >= 0; i-- {
		packSize := arr[i]
		// Remove multiples of one pack whilst remainding is greater than or equal too (most likely the biggest pack size)
		for remaining >= packSize {
			result = append(result, packSize)
			remaining -= packSize
		}
	}

	//Resolves for final iteration - if less than the smallest packsize it'll never trigger in the above loop
	if remaining > 0 {
		result = append(result, arr[0])
	}

	//If wastage is the same then only send one pack which is larger for efficiency
	if len(result) >= 2 && result[len(result)-1] == result[len(result)-2] && (result[len(result)-1]+result[len(result)-2]) == arr[1] {
		result = result[:len(result)-2]
		result = append(result, arr[1])
	}

	return result

}

// route handler function
func orderPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	var payload Request
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp Response

	var result = orderPackSize(payload.Target, payload.Numbers)
	resp.Result = result
	json.NewEncoder(w).Encode(resp)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		w.Write([]byte("Server Error"))
	} else {
		w.Write(data)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/packsize", orderPackSizesHandler)
	http.ListenAndServe(":8080", nil)
}
