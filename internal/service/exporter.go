package service

import (
	"encoding/csv"
	"io"
	"net/http"
	"strings"
)

func GetDataFromKobo(CSVlink, userLogin, userPass string) ([][]string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", CSVlink, nil)
	if err != nil {
		return [][]string{{}}, err
	}

	if userPass != "" && userLogin != "" {
		request.SetBasicAuth(userLogin, userPass)
	}

	response, err := client.Do(request)
	if err != nil {
		return [][]string{{}}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return [][]string{{}}, err
	}

	defer response.Body.Close()

	r := csv.NewReader(strings.NewReader(string(body)))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{{}}, err
	}

	return records, nil

}
