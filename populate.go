package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	templateFile, err := os.Open("template.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully opened template file")
	defer templateFile.Close()

	byteValue, _ := io.ReadAll(templateFile)

	var template map[string]interface{}
	json.Unmarshal([]byte(byteValue), &template)

	fmt.Println("Number of items to generate: ")
	var number int
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Invalid number of items.")
		return
	}

	outputArray := make([]map[string]interface{}, number)
	for i := 0; i < number; i++ {
		item := template
		for k := range template {
			item[k] = RandStringRunes(10)
		}
		outputArray[i] = item
	}

	jsonItem, err := json.MarshalIndent(outputArray, " ", " ")
	if err != nil {
		fmt.Println("Error while marshaling item")
		return
	}
	err = os.WriteFile("output.json", jsonItem, 0644)
	if err != nil {
		fmt.Println("Error while writing to file")
		return
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
