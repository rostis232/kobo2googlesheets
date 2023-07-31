package service

import (
	"context"
	b64 "encoding/base64"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ImportToSheet(credentials string, spreadsheetId string, sheetName string, cellsRange string, values [][]interface{}) error {
	ctx := context.Background()

	credBytes, err := b64.StdEncoding.DecodeString(credentials)
	if err != nil {
		//log.Println(err)
		return err
	}

	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		//log.Println(err)
		return err
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		//log.Println(err)
		return err
	}

	row := &sheets.ValueRange{
		Values: values,
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, sheetName+"!"+cellsRange, row).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		//log.Println(err)
		return err
	}

	return nil
}
