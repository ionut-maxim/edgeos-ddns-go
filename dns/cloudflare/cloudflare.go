package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

// Create a new API object
func New(key string, email string) (*cloudflare.API, error) {

	api, err := cloudflare.New(key, email)
	if err != nil {
		return nil, fmt.Errorf("unable to create api")
	}
	return api, nil
}

// Get zone ID by zone name.
func zoneID(name string, api cloudflare.API) (string, error) {
	id, err := api.ZoneIDByName(name)

	if err != nil {
		return "", fmt.Errorf("unable to get zone id: %v", err)
	}

	return id, nil
}

func GetRecord(recordName string, zoneID string, api cloudflare.API) (id string, content string) {
	record := cloudflare.DNSRecord{Name: recordName}

	records, err := api.DNSRecords(context.Background(), zoneID, record)

	if err != nil {
		log.Fatal(err)
	}

	if len(records) > 0 && records[0].Content != "" {
		return records[0].ID, records[0].Content
	} else {
		return "", ""
	}
}

func UpdateRecord(zoneID string, recordID string, content string, api cloudflare.API) {
	record := cloudflare.DNSRecord{Content: content}

	err := api.UpdateDNSRecord(context.Background(), zoneID, recordID, record)

	if err != nil {
		log.Fatal(err)
	}
}
