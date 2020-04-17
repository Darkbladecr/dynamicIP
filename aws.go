package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var svc *route53.Route53

func initAWS() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	svc = route53.New(sess)
}

func getAWSRecords(zoneID string, ip string) (dnsRecords []*route53.ResourceRecordSet) {
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}
	respList, err := svc.ListResourceRecordSets(listParams)
	if err != nil {
		log.Fatal(err)
	}
	for _, rec := range respList.ResourceRecordSets {
		if *rec.Type == "A" {
			match := false
			for _, resRec := range rec.ResourceRecords {
				if *resRec.Value == ip {
					match = true
				}
			}
			if match {
				dnsRecords = append(dnsRecords, rec)
			}
		}
	}
	return dnsRecords
}

func updateAWSRecords(zoneID string, dnsRecords []*route53.ResourceRecordSet, oldIP string, newIP string) {
	var changes []*route53.Change
	for _, recordSet := range dnsRecords {
		for _, record := range recordSet.ResourceRecords {
			if *record.Value == oldIP {
				record.SetValue(newIP)
			}
		}
		change := &route53.Change{
			Action:            aws.String("UPSERT"),
			ResourceRecordSet: recordSet,
		}
		changes = append(changes, change)
	}
	// fmt.Printf("%+v", changes)
	changeBatch := &route53.ChangeBatch{
		Changes: changes,
	}
	input := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
		ChangeBatch:  changeBatch,
	}
	out, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", out)
}
