package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	

	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		

		checkDomain(scanner.Text())
		fmt.Println("")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error : could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMx,hasDMARC,hasSPF bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Fatal(err)
	}
	if len(mxRecords) > 0{
		hasMx = true
	}

	txtRecord, err := net.LookupTXT(domain)

	if err != nil {
		log.Fatal(err)
	}
	for _ , record := range txtRecord{
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF= true
			spfRecord = record
			break
		}

	}

	dmarcRecords , err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Fatal(err)
	}
	
	for _ , record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC= true
			dmarcRecord = record
			break
		}

	}

	fmt.Println(domain,hasMx,hasSPF,spfRecord,hasDMARC,dmarcRecord)
}
