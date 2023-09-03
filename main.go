package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	connect "github.com/abrander/garmin-connect"
)

func readCreds() (email string, password string) {
	var exists bool
	email, exists = os.LookupEnv("GARMIN_EMAIL")
	if !exists {
		log.Fatal("GARMIN_EMAIL environment variable not set")
	}
	password, exists = os.LookupEnv("GARMIN_PASSWORD")
	if !exists {
		log.Fatal("GARMIN_PASSWORD environment variable not set")
	}

	return email, password
}

func mphToKph(mph float64) (kph float64) {
	return mph * 1.609344
}

func formatActivities(activities []connect.Activity) (transformedActivites []connect.Activity) {
	for _, activity := range activities {
		activity := &activity
		activity.MaxSpeed = mphToKph(activity.MaxSpeed)
		activity.AverageSpeed = mphToKph(activity.AverageSpeed)

		transformedActivites = append(transformedActivites, *activity)

	}
	return transformedActivites
}

func main() {
	email, password := readCreds()
	creds := connect.Credentials(email, password)
	c := connect.NewClient(creds, connect.AutoRenewSession(true))
	err := c.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	activities, err := c.Activities("", 1, 20)
	if err != nil {
		log.Fatal(err)
	}

	activitiesJSON, err := json.Marshal(formatActivities(activities))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(activitiesJSON))
}
