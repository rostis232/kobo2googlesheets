package models

type Sheet struct {
	Title         string
	Credentials   string
	SpreadsheetId string
	SheetName     string
	CellsRange    string
}

type Form struct {
	Title     string
	CSVlink   string
	UserLogin string
	UserPass  string
	Sheets    []Sheet
}
