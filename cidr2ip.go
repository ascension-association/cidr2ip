package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/adedayo/cidr"
)

const (
	Version = "2.0.0"
	IPRegex = `\b(?:\d{1,3}\.){3}\d{1,3}\b$`
)

var (
	cidrFilePtr = flag.String("f", "",
		"[Optional] Name of file with CIDR blocks")
	printRangesPtrPtr = flag.Bool("r", false,
		"[Optional] Print IP ranges instead of all IPs")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]

	if *cidrFilePtr != "" {
		file, err := os.Open(*cidrFilePtr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			displayIPs(scanner.Text())
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	} else if info.Mode()&os.ModeNamedPipe != 0 { // data is piped in
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			displayIPs(scanner.Text())
		}
	} else if len(args) > 0 { // look for CIDRs on cmd line
		var cidrs []string
		if *printRangesPtrPtr == true {
			cidrs = args[1:]
		} else {
			cidrs = args
		}

		for _, cs := range cidrs {
			displayIPs(cs)
		}
	} else { // no piped input, no file provide and no args, display usage
		flag.Usage()
	}
}

func isIPAddr(cs string) bool {
	match, _ := regexp.MatchString(IPRegex, cs)
	return match
}

func displayIPs(cs string) {
	ips := cidr.Expand(cs)

	for _, ip := range ips {
		println(ip)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "CIDR to IPs version %s\n", Version)
	fmt.Fprintf(os.Stderr, "Usage:   $ cidr2ip [-r] [-f <filename>] <list of cidrs> \n")
	fmt.Fprintf(os.Stderr, "Example: $ cidr2ip -f cidrs.txt\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip 10.0.0.0/24\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip -r 10.0.0.0/24\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip -r -f cidrs.txt\n")
	fmt.Fprintf(os.Stderr, "         $ cat cidrs.txt | cidr2ip \n")
	fmt.Fprintf(os.Stderr, "--------------------------\nFlags:\n")
	flag.PrintDefaults()
}
