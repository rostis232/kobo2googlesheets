package config

import "time"

type Config struct {
	Credentials        string
	UpdatesPeriod      time.Duration
	SleepTimeStart     int
	SleepTimeEnd       int
	StandartSheetName  string
	StandartCellsRange string
}
