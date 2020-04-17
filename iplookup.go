package main

import (
	"io/ioutil"
	"os"
)

func saveIP(ip string) {
	f, err := os.Create("ip.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(ip)
	if err != nil {
		panic(err)
	}
	f.Sync()
}

func getOldIP() string {
	ip, err := ioutil.ReadFile("ip.txt")
	if err != nil {
		return ""
	}
	return string(ip)
}
