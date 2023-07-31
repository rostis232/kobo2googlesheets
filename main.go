package main

import (
	"encoding/json"
	"fmt"
	"github.com/rostis232/kobo2googlesheets/internal/models"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error while opening data file:", err)
		return
	}

	var data []models.Form

	parsedData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading data file:", err)
		return
	}

	file.Close()

	err = json.Unmarshal(parsedData, &data)
	if err != nil {
		fmt.Println("Error while unmarshalling data from file:", err)
		return
	}
	fmt.Println(data)

}
