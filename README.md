# kobo2googlesheets
The program allows you to regularly and automatically export and import data from KoboToolbox forms into Google Sheets.

## How to make settings

This program needs 2 json files: `data.json` and `settings.json`. Program will read these files during the start. Next is a description of both of them.

### settings.json

Sample file contents:

```
{
  "Credentials": "ewogICJ0eXB",
  "UpdatesPeriod": "2h",
  "SleepTimeStart": 1,
  "SleepTimeEnd": 7,
  "StandartSheetName": "kobo",
  "StandartCellsRange": "A1:XYZ"
}
```

In this file all fields are required. 

`Credentials` - it is a base64-encoded .json file with a credentials to use Google Sheet API. You can get it while activating Google Sheet API. This is the default value that will be used unless otherwise specified for a particular Spreadsheet.

`UpdatesPeriod` - it is a period of downtime between each iteration. You can use there `m` (like `20m`)for minutes and `h` for hours (like `2h`).

`SleepTimeStart` - it is an hour in 24-hours format when begins a period of program sleeping.

`SleepTimeEnd` - it is an hour in 24-hours format when ends a period of program sleeping.

In this example the program will not work from 1 to 7 hours in 24 hours format.

`StandartSheetName` - It is the name of the sheet that will be used to import data. This is the default value that will be used unless otherwise specified for a particular Spreadsheet.

`StandartCellsRange` - It is the name of the range of cells that will be used to import data. This is the default value that will be used unless otherwise specified for a particular Spreadsheet. I recommend to use "A1:XYZ" in a way to import from beginning of the sheet.

### data.json

Sample file contents:

```
[
  {
    "Title": "KoboFormTitle1",
    "CSVlink": "https://kobo.humanitarianresponse.info/api/v2/assets/AAAAAAAAA/export-settings/BBBBBBB/data.csv",
    "UserLogin": "YourKoboLogin1",
    "UserPass": "YourKoboPass1",
    "Sheets": [
      {
        "Title": "GoogleSpreadSheetTitle1",
        "SpreadsheetId": "SpreadsheetIdAAAAAA",
        "Credentials": "ewogICJ0eXB",
        "SheetName": "KOBO",
        "CellsRange": "A10:X"
      },
      {
        "Title": "NIN PROT CASE Івано-Франківськ",
        "SpreadsheetId": "SpreadsheetIdBBBBB"
      }
    ]
  },
  {
    "Title": "KoboFormTitle2",
    "CSVlink": "https://kobo.humanitarianresponse.info/api/v2/assets/CCCCCCCC/export-settings/DDDDDDD/data.csv",
    "Sheets": [
      {
        "Title": "GoogleSpreadSheetTitle2",
        "SpreadsheetId": "SpreadsheetIdCCCCCCC"
      }
    ]
  }
]
```

Fields `UserLogin` and `UserPass` are not required. If they are empty program will try get information without authentication. In this case access to assets in KoboToolbox form must be opened.

`Credentials`, `SheetName`, `CellsRange` - these fields are not required, if empty will be used default values from `settings.json`.

## How to run

This program is written in the Golang language, so you will need to install Golang to compile and run the program.

If you have Golang installed just type into terminal or command line to run it:

`go run main.go`
