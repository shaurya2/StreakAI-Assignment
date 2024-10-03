package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// album represents data about a record album.
type FindPairsRequest struct {
	Number []int `json:"number"`
	Target int   `json:"target"`
}

type FindTargetResponse struct {
	Solutions [][]int `json:"solutions"`
}

func findPairs(w http.ResponseWriter, r *http.Request) {
	var targets FindPairsRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(reqBody, &targets)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := GetTarget(&targets)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetTarget(reqBody *FindPairsRequest) (resp *FindTargetResponse, err error) {

	nums := reqBody.Number
	target := reqBody.Target
	m := make(map[int]int)
	var result [][]int

	for i, num := range nums {
		needed := target - num

		if index, found := m[needed]; found {
			result = append(result, []int{index, i})
		}

		m[num] = i
	}
	resp.Solutions = result
	return
}
