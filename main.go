package main

import (
	"encoding/json"
	"fmt"
	"github.com/rostis232/kobo2googlesheets/internal/config"
	"github.com/rostis232/kobo2googlesheets/internal/models"
	"github.com/rostis232/kobo2googlesheets/internal/service"
	"io"
	"os"
	"time"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Printf("Error while opening data file: %s. Check if file data.json exists.", err)
		return
	}

	var data []models.Form

	parsedData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading data file:", err)
		return
	}

	if file.Close() != nil {
		fmt.Println("Error while closing file data.json: ", err)
	}

	err = json.Unmarshal(parsedData, &data)
	if err != nil {
		fmt.Println("Error while unmarshalling data from file:", err)
		return
	}

	file, err = os.Open("settings.json")
	if err != nil {
		fmt.Println("Error while opening settings file:", err)
		return
	}

	var settings config.Config

	parsedData, err = io.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading settings file:", err)
		return
	}

	if file.Close() != nil {
		fmt.Println("Error while closing file config.json: ", err)
	}

	err = json.Unmarshal(parsedData, &settings)
	if err != nil {
		fmt.Println("Error while unmarshalling settings from file:", err)
		return
	}

	for {
		iterationStartTime := time.Now()

		if iterationStartTime.Hour() >= 1 && iterationStartTime.Hour() < 7 {
			fmt.Println("Sleep time....")
			time.Sleep(time.Hour)
			continue
		}

		formsCount := 0
		sheetsCount := 0
		errorsCount := 0

		fmt.Println("New iteration started.")

		for _, f := range data {

			if f.Sheets == nil || len(f.Sheets) == 0 {
				fmt.Printf("Form <<%s>> hasn`t any sheets to import. Exporting skiped.", f.Title)
				continue
			}
			formsCount++
			startTime := time.Now()

			var (
				userLogin string
				userPass  string
			)

			if f.UserLogin != "" && f.UserPass != "" && &f.UserLogin != nil && &f.UserPass != nil {
				userLogin = f.UserLogin
				userPass = f.UserPass
			}

			fmt.Printf("Started exporting data from form <<%s>>.\n", f.Title)

			records, err := service.GetDataFromKobo(f.CSVlink, userLogin, userPass)
			if err != nil {
				errorsCount++
				fmt.Printf("Error while getting data from Kobo-form <<%s>>: %s\n", f.Title, err)
				continue
			}
			values := service.StringSlices2Interfaces(records)

			for _, s := range f.Sheets {
				sheetsCount++

				var credentials string

				if s.Credentials == "" || &s.Credentials == nil {
					credentials = settings.Credentials
				} else {
					credentials = s.Credentials
				}

				var sheetName string

				if s.SheetName == "" || &s.SheetName == nil {
					sheetName = settings.StandartSheetName
				} else {
					sheetName = s.SheetName
				}

				var cellsRange string

				if s.CellsRange == "" || &s.CellsRange == nil {
					cellsRange = settings.StandartCellsRange
				} else {
					cellsRange = s.CellsRange
				}

				fmt.Printf("Started importing to Spreadsheet <<%s>>.\n", s.Title)

				err = service.ImportToSheet(credentials, s.SpreadsheetId, sheetName, cellsRange, values)

				if err != nil {
					errorsCount++
					fmt.Printf("Error while importing data to sheet <<%s>>: %s\n", f.Title, err)
					break
				}

				fmt.Printf("Successful importing to Spreadsheet <<%s>>.\n", s.Title)
			}
			endTime := time.Now()
			totalTime := endTime.Sub(startTime)
			fmt.Printf("Process of importing/exporting from form <<%s>> completed. Total time spent: %g seconds.\n", f.Title, totalTime.Seconds())

		}
		iterationEndTime := time.Now()
		iterationTime := iterationEndTime.Sub(iterationStartTime)
		fmt.Printf("Current iteration compleated. Forms: %d. Sheets: %d. Errors: %d. Total time spent at this itteration: %g minutes. Waiting for the next ...\n", formsCount, sheetsCount, errorsCount, iterationTime.Minutes())

		updateDuration, err := time.ParseDuration(settings.UpdatesPeriod)
		if err != nil {
			fmt.Println("Error Updates Period Parsing: ", err)
			return
		}

		time.Sleep(updateDuration)
	}

}
