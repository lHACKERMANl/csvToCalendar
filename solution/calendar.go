package solution

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func readCredentials(filePath string) []byte {
	credentials, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}
	return credentials
}

func getToken(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func csvParser(csvPath string) ([]*calendar.Event, error) {
	var events []*calendar.Event

	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(record) == 8 {
			dateLayout := "02/01/2006"
			timeLayout := "15:04"

			date, err := time.Parse(dateLayout, record[0])
			if err != nil {
				// log.Println("Unable to parse", err)
				continue
			}

			startTime, err := time.Parse(timeLayout, record[3])
			if err != nil {
				// log.Println("Unable to parse", err)
				continue
			}

			endTime, err := time.Parse(timeLayout, record[5])
			if err != nil {
				// log.Println("Unable to parse", err)
				continue
			}

			startDateTime := date.Add(time.Hour*time.Duration(startTime.Hour()) + time.Minute*time.Duration(startTime.Minute()))
			endDateTime := date.Add(time.Hour*time.Duration(endTime.Hour()) + time.Minute*time.Duration(endTime.Minute()))

			summary := record[1]
			location := record[2]
			description := record[7]
			colorId := "11"

			event := &calendar.Event{
				Summary:  summary,
				Location: location,
				Start: &calendar.EventDateTime{
					DateTime: startDateTime.Add(-3 * time.Hour).Format(time.RFC3339),
					TimeZone: "UTC",
				},
				End: &calendar.EventDateTime{
					DateTime: endDateTime.Add(-3 * time.Hour).Format(time.RFC3339),
					TimeZone: "UTC",
				},
				Description: description,
				ColorId:     colorId,
			}

			events = append(events, event)
		}
	}

	return events, nil
}

func CsvToCalendar(credPath string, csvPath string) {
	tokenPath := "solution/token.json"

	config, err := google.ConfigFromJSON(readCredentials(credPath), calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	var token *oauth2.Token
	token, err = tokenFromFile(tokenPath)
	if err != nil {
		fmt.Println("Token not found, obtaining a new one...")
		token = getToken(config)
		saveToken(tokenPath, token)
	}

	ctx := context.Background()
	client, err := calendar.NewService(ctx, option.WithHTTPClient(config.Client(ctx, token)))
	if err != nil {
		log.Fatalf("Unable to create Google Calendar API client: %v", err)
	}

	events, err := csvParser(csvPath)
	if err != nil {
		fmt.Println("Unable to parse", err)
	}

	for _, v := range events {
		v, err = client.Events.Insert("primary", v).Do()
		if err != nil {
			log.Fatalf("Unable to insert event: %v", err)
		}
		fmt.Printf("Event created: %s\n", v.HtmlLink)
	}
}
