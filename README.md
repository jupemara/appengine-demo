## appengine-demo

demo project

## Components

- App Engine
- Google Sheets
- Cloud Trace
- OpenTelemetry
- Cloud Firestore

## How to start this demo

```sh
# create app engine application
$ gcloud app create

# deploy your app
$ gcloud app deploy
```

## Execution in local

```sh
$ go run main.go
```

## Prerequisites

### Environment variables

- APPENGINEDEMO_GCP_PROJECT_ID
- APPENGINEDEMO_SHEETS_KEY
- APPENGINEDEMO_SHEETS_NAME
