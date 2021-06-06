package config

import (
	"fmt"
	"log"
	"os"
)

type config struct {
	gcpProjectId string
	sheetsKey    string
	sheetsName   string
	port         string
}

func NewConfig() config {
	return config{}
}

func (c *config) Load() error {
	gcpProjectId := os.Getenv("APPENGINEDEMO_GCP_PROJECT_ID")
	if len(gcpProjectId) <= 0 {
		return fmt.Errorf(`gcp project id must be set`)
	}
	sheetsKey := os.Getenv("APPENGINEDEMO_SHEETS_KEY")
	if len(sheetsKey) <= 0 {
		return fmt.Errorf(`google sheets key must be set`)
	}
	sheetsName := os.Getenv("APPENGINEDEMO_SHEETS_NAME")
	if len(sheetsName) <= 0 {
		return fmt.Errorf(`google sheets sheet name must be set`)
	}
	port := os.Getenv("APPENGINEDEMO_PORT")
	if len(port) <= 0 {
		log.Println(`set exported port 8080 as default`)
		port = "8080"
	}
	c.gcpProjectId = gcpProjectId
	c.sheetsKey = sheetsKey
	c.sheetsName = sheetsName
	c.port = port
	return nil
}

func (c config) GcpProjectId() string {
	return c.gcpProjectId
}

func (c config) SheetsKey() string {
	return c.sheetsKey
}

func (c config) SheetsName() string {
	return c.sheetsName
}

func (c config) Port() string {
	return c.port
}
