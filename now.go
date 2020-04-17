package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type zeitRecord struct {
	ID      string `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Creator string `json:"creator"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

type zeitRecords struct {
	Records []zeitRecord `json:"records"`
}

var authHeader string

func initNow() {
	token := os.Getenv("NOW_TOKEN")
	authHeader = fmt.Sprintf("Bearer %s", token)
}

func getNowDNSRecords(domain string) []zeitRecord {
	url := fmt.Sprintf("https://api.zeit.co/v2/domains/%s/records", domain)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var records zeitRecords
	if err := json.NewDecoder(res.Body).Decode(&records); err != nil {
		panic(err)
	}
	return records.Records
}

func deleteNowRecord(domain, recID string) {
	url := fmt.Sprintf("https://api.zeit.co/v2/domains/%s/records/%s", domain, recID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}

type nowDNSRecordI struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func createNowRecord(domain, name, dnsType, value string) {
	postData := &nowDNSRecordI{Name: name, Type: dnsType, Value: value}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(postData)
	url := fmt.Sprintf("https://api.zeit.co/v2/domains/%s/records", domain)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}
