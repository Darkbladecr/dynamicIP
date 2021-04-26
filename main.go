package main

import (
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
	"github.com/rdegges/go-ipify"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initCF()
}

func main() {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Fatalln("Couldn't get my IP address:", err)
	}
	oldIP := getOldIP()
	if oldIP == "" {
		oldIP = ip
		saveIP(ip)
	}
	if oldIP != ip {
		// log.Printf("old ip: %s", oldIP)
		// log.Printf("new ip: %s", ip)
		saveIP(ip)
		// Cloudflare
		recs := getCFRecords("mitrasinovic.co.uk", oldIP)
		recs2 := getCFRecords("exambuddy.co.uk", oldIP)
		recs3 := getCFRecords("quesmed.com", oldIP)
		recs4 := getCFRecords("relaydr.com", oldIP)
		var totalLen int
		slices := [][]cloudflare.DNSRecord{
			recs, recs2, recs3, recs4,
		}
		for _, s := range slices {
			totalLen += len(s)
		}
		records := make([]cloudflare.DNSRecord, totalLen)
		var i int
		for _, s := range slices {
			i += copy(records[i:], s)
		}
		// s, _ := json.MarshalIndent(records, "", "\t")
		// log.Printf("%s", s)
		updateCFRecords(records, ip)
	}
}
