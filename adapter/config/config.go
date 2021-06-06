package config

import (
	"fmt"
	"os"
)

type config struct {
	gcpProjectId    string
	sheetsKey       string
	sheetsSheetName string
}

func NewConfig() config {
	return config{}
}

func (c config) Load() error {
	gcpProjectId := os.Getenv("APPENGINEDEMO_GCP_PROJECT_ID")
	if len(gcpProjectId) <= 0 {
		return fmt.Errorf(`gcp project id must be set`)
	}
	sheetsKey := os.Getenv("APPENGINEDEMO_SHEETS_KEY")
	if len(sheetsKey) <= 0 {
		return fmt.Errorf(`google sheets key must be set`)
	}
	sheetsSheetName := os.Getenv("APPENGINEDEMO_SHEETS_NAME")
	if len(sheetsSheetName) <= 0 {
		return fmt.Errorf(`google sheets sheet name must be set`)
	}
	c.gcpProjectId = gcpProjectId
	c.sheetsKey = sheetsKey
	c.sheetsSheetName = sheetsSheetName
	return nil
}
