package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Result []int `json:"result"`
}

func orderPackSize(x int, arr []int) ([]int, error) {
	// arr := []int{250, 500, 1000, 2000, 5000}
	result := []int{}

	remaining := x
	for i := len(arr) - 1; i >= 0; i-- {
		packSize := arr[i]
		for remaining >= packSize {
			result = append(result, packSize)
			remaining -= packSize
		}
	}

	if remaining > 0 {
		result = append(result, arr[0])
	}

	if len(result) >= 2 && result[len(result)-1] == result[len(result)-2] && (result[len(result)-1]+result[len(result)-2]) == arr[1] {
		result = result[:len(result)-2]
		result = append(result, arr[1])
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	// wastage := sum - x

	return result, nil

	// return map[string]interface{}{
	// 	"input":   x,
	// 	"result":  result,
	// 	"wastage": wastage,
	// }, nil
}

func orderPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp Response

	var result, err2 = orderPackSize(req.Target, req.Numbers)
	if err2 != nil {
		// Do something ere
	} else {
		resp.Result = result
	}

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
