package main

import (
	"fmt"
	"os"

	namecheap "github.com/billputer/go-namecheap"
)

var namecheapToken string
var user string

func initNC() {
	namecheapToken = os.Getenv("NAMECHEAP_TOKEN")
	user = os.Getenv("NAMECHEAP_USER")
}

func getNCDNSRecords() {
	client := namecheap.NewClient(user, namecheapToken, user)
	domains, _ := client.DomainsGetList()
	for _, domain := range domains {
		fmt.Printf("Domain: %+v\n\n", domain.Name)
	}
}
