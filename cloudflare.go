package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var cfAPI *cloudflare.API

func initCF() {
	var err error
	cfAPI, err = cloudflare.NewWithAPIToken(os.Getenv("CF_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
}

func getCFRecords(zoneName string, ip string) []cloudflare.DNSRecord {
	zoneID, err := cfAPI.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	recs, err := cfAPI.DNSRecords(ctx, zoneID, cloudflare.DNSRecord{Content: ip})
	if err != nil {
		log.Fatal(err)
	}
	return recs
}

func updateCFRecords(recs []cloudflare.DNSRecord, ip string) {
	for _, rec := range recs {
		rec.Content = ip
		ctx := context.Background()
		err := cfAPI.UpdateDNSRecord(ctx, rec.ZoneID, rec.ID, rec)
		if err != nil {
			log.Fatal(err)
		}
	}
}
