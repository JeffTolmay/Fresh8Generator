package main

import (
	"github.com/google/uuid"
)

// Data is needed (lint)
type Data struct {
	ViewID        string `json:"viewId"`
	EventDateTime string `json:"eventDateTime"`
}

// Event is needed (lint)
type Event struct {
	Type      string `json:"type"`
	InnerData Data   `json:"data"`
}

func createViewedData(dataTime string, viewedOnly float64, interactedOnly float64, clickOnly float64, clickAndInteracted float64, directoryString string, interval *int) {

	var allEvents []Event

	for i := 0; i < int(viewedOnly); i++ {
		id := uuid.New()
		uuid := id.String()
		var dataEvents Event
		dataEvents.Type = "Viewed"
		dataEvents.InnerData.EventDateTime = dataTime
		dataEvents.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEvents)
	}

	for i := 0; i < int(interactedOnly); i++ {
		id := uuid.New()
		uuid := id.String()
		var dataEvents Event
		dataEvents.Type = "Interacted"
		dataEvents.InnerData.EventDateTime = dataTime
		dataEvents.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEvents)
		var dataEventsNext Event
		dataEventsNext.Type = "Viewed"
		dataEventsNext.InnerData.EventDateTime = dataTime
		dataEventsNext.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEventsNext)
	}

	for i := 0; i < int(interactedOnly); i++ {
		id := uuid.New()
		uuid := id.String()
		var dataEvents Event
		dataEvents.Type = "Click-Through"
		dataEvents.InnerData.EventDateTime = dataTime
		dataEvents.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEvents)
		var dataEventsNext Event
		dataEventsNext.Type = "Viewed"
		dataEventsNext.InnerData.EventDateTime = dataTime
		dataEventsNext.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEventsNext)
	}

	for i := 0; i < int(interactedOnly); i++ {
		id := uuid.New()
		uuid := id.String()
		var dataEvents Event
		dataEvents.Type = "Click-Through"
		dataEvents.InnerData.EventDateTime = dataTime
		dataEvents.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEvents)
		var dataEventsNext Event
		dataEventsNext.Type = "Viewed"
		dataEventsNext.InnerData.EventDateTime = dataTime
		dataEventsNext.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEventsNext)
		var dataEventsAgain Event
		dataEventsAgain.Type = "Interacted"
		dataEventsAgain.InnerData.EventDateTime = dataTime
		dataEventsAgain.InnerData.ViewID = uuid
		allEvents = append(allEvents, dataEventsAgain)
	}

	writeAsBatch(allEvents, directoryString)

}
