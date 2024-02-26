package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "integrations"

var (
	ErrListAgendas    = errors.New("error list your agendas")
	ErrAgendaNotFound = errors.New("agenda not found")
	ErrEventsList     = errors.New("error list events")
	ErrAddAgenda      = errors.New("error add agenda to service account")
)

type Calendar struct {
	Service *gCalendar.Service
}

func NewClient() *Calendar {
	ctx := context.Background()
	credentialsPath := os.Getenv("CREDENTIALS_PATH")

	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(b))
	if err != nil {
		log.Fatal(err)
	}
	return &Calendar{
		Service: service,
	}
}

func (c *Calendar) AddAgenda(id string) error {
	ag := &gCalendar.CalendarListEntry{
		Id: id,
	}
	add, err := c.Service.CalendarList.Insert(ag).Do()
	if err != nil {
		log.Fatal(err)
		return ErrAddAgenda
	}
	fmt.Println("==>", add.HTTPStatusCode)

	return nil
}

func (c *Calendar) GetAgendaID() (string, error) {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		return "", ErrListAgendas
	}
	fmt.Println("==>", list.HTTPStatusCode)

	for _, v := range list.Items {
		if v.Summary == AGENDA {
			return v.Id, nil
		}
	}

	return "", ErrAgendaNotFound
}

// ListWeekEvents returns all events in a week of a calendar
func (c *Calendar) ListWeekEvents() ([]string, error) {
	allEvents := []string{}
	id, err := c.GetAgendaID()
	if err != nil {
		return []string{}, err
	}

	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)
	events, err := c.Service.Events.List(id).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return []string{}, err
	}

	for _, v := range events.Items {
		allEvents = append(allEvents, fmt.Sprintf("%s | %s | at %s\n", v.Summary, v.Status, v.Start.DateTime))
	}
	return allEvents, nil
}

func (c *Calendar) ListTodayEvents() ([]string, error) {
	allEvents := []string{}
	id, err := c.GetAgendaID()
	if err != nil {
		return []string{}, err
	}

	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)

	events, err := c.Service.Events.List(id).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return []string{}, err
	}

	for _, v := range events.Items {
		allEvents = append(allEvents, fmt.Sprintf("Nome do evento: %s | status: %s | quando: %s | timezone: %s\n", v.Summary, v.Status, v.Start.DateTime, events.TimeZone))
	}
	return allEvents, nil
}
