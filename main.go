package main

import (
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/rdegges/go-ipify"
)

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	initCF()
	initAWS()
}

func main() {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Fatalln("Couldn't get my IP address:", err)
	}
	oldIP := getOldIP()
	if oldIP != ip {
		saveIP(ip)
		// Cloudflare
		recs := getCFRecords("mitrasinovic.co.uk", oldIP)
		recs2 := getCFRecords("exambuddy.co.uk", oldIP)
		recs = append(recs, recs2...)
		updateCFRecords(recs, ip)
		// AWS Route53
		zoneID := os.Getenv("AWS_ZONE_ID")
		dnsRecords := getAWSRecords(zoneID, oldIP)
		updateAWSRecords(zoneID, dnsRecords, oldIP, ip)
	}
}
