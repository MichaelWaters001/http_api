package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Life struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

type Weight struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

type Attributes struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Life           Life   `json:"life"`
	MaleWeight     Weight `json:"male_weight"`
	FemaleWeight   Weight `json:"female_weight"`
	Hypoallergenic bool   `json:"hypoallergenic"`
}

type GroupData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Relationships struct {
	Group struct {
		Data GroupData `json:"data"`
	} `json:"group"`
}

type Breed struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type ApiResponse struct {
	Data  []Breed `json:"data"`
	Links struct {
		Self    string `json:"self"`
		Current string `json:"current"`
		Next    string `json:"next"`
		Last    string `json:"last"`
	} `json:"links"`
}

var Breeds []Breed

// init Breeds from dogapi.dog
func initDogBreeds() {
	resp, err := http.Get("https://dogapi.dog/api/v2/breeds")
	if err != nil {
		fmt.Printf("Failed to fetch dog Breeds")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to retrieve dog Breeds")
		return
	}

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Printf("Failed to unmarshal dog Breeds")
		return
	}
	Breeds = apiResponse.Data
}

func getDogBreeds(context *gin.Context) {
	var breedNames []string
	for _, breed := range Breeds {
		breedNames = append(breedNames, breed.Attributes.Name)
	}
	context.IndentedJSON(http.StatusOK, breedNames)
}
