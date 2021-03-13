package main

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

func cfClient(key string, email string) (result cloudflare.API) {
	// Construct a new API object
	api, err := cloudflare.New(key, email)
	if err != nil {
		log.Fatal(err)
	}

	return *api
}

func getZoneID(zoneName string, api cloudflare.API) (id string) {
	zoneId, err := api.ZoneIDByName(zoneName)

	if err != nil {
		log.Fatal(err)
	}

	return zoneId
}

func getDNSRecord(recordName string, zoneID string, api cloudflare.API) (id string, content string) {
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

func updateDNSRecord(zoneID string, recordID string, content string, api cloudflare.API) {
	record := cloudflare.DNSRecord{Content: content}

	err := api.UpdateDNSRecord(context.Background(), zoneID, recordID, record)

	if err != nil {
		log.Fatal(err)
	}
}
