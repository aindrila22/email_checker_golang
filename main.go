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
	scanner := bufio.NewScanner(os.Stdin)//call the scan function and accept the input
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord \n")//fields to check in every mail

	for scanner.Scan(){
		checkDomain(scanner.Text())//passing to check
	}

	if err := scanner.Err(); err!=nil{
		log.Fatal("Error: could not read from input: %v\n", err)//printing err
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool//define variables
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)//check mxfield with predefined .net library of goLang

	if err!=nil{
		log.Printf("Error %v\n", err)
	}
	if len(mxRecords) > 0{
		hasMX = true
	}
    
	txtRecords, err := net.LookupTXT(domain)
	if err!=nil{
		log.Printf("Error %v\n", err)
	}
    for _, record := range txtRecords{//range over the function and check v=spf1 field
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc."+ domain)
	if err!=nil{
		log.Printf("Error %v\n", err)
	}
	for _, record := range dmarcRecords{//range over the function and check v=DMARC1 field
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v",domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)//print the output
}