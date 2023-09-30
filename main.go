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

// converts m/s to km/h
func msToKph(ms float64) (kmh float64) {
	kmh = ms * 3.6
	return kmh
}

func formatActivities(activities []connect.Activity) (transformedActivites []connect.Activity) {
	for _, activity := range activities {
		activity.MaxSpeed = msToKph(activity.MaxSpeed)
		activity.AverageSpeed = msToKph(activity.AverageSpeed)
		transformedActivites = append(transformedActivites, activity)

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

	activities, err := c.Activities("", 1, 1000)
	if err != nil {
		log.Fatal(err)
	}

	activitiesJSON, err := json.Marshal(formatActivities(activities))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(activitiesJSON))
}
