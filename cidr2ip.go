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
	Version = "3.0.0"
	IPRegex = `\b(?:\d{1,3}\.){3}\d{1,3}\b$`
)

var (
	cidrFilePtr = flag.String("i", "",
		"[Optional] Name of input file with CIDR blocks")
	outputFilePtr = flag.String("o", "",
		"[Optional] Name of output file for storing result")
	outputWriterPtr *bufio.Writer
)

func main() {
	flag.Usage = usage
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]

	if *outputFilePtr != "" {
		f, err := os.Create(*outputFilePtr)
		if err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			outputWriterPtr = bufio.NewWriter(f)
		}
	}

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
		cidrs = args

		for _, cs := range cidrs {
			displayIPs(cs)
		}
	} else { // no piped input, no file provide and no args, display usage
		flag.Usage()
	}

	if outputWriterPtr != nil {
		outputWriterPtr.Flush()
	}
}

func isIPAddr(cs string) bool {
	match, _ := regexp.MatchString(IPRegex, cs)
	return match
}

func displayIPs(cs string) {
	ips := cidr.Expand(cs)

	for _, ip := range ips {
		if *outputFilePtr != "" {
			outputWriterPtr.WriteString(ip + "\n")
		} else {
			println(ip)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "CIDR to IPs version %s\n", Version)
	fmt.Fprintf(os.Stderr, "Usage:   $ cidr2ip [-i <filename>] [-o <filename>] [<list of cidrs>] \n")
	fmt.Fprintf(os.Stderr, "Example: $ cidr2ip 10.0.0.0/27\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip -i cidrs.txt\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip -o results.txt 10.0.0.0/27\n")
	fmt.Fprintf(os.Stderr, "         $ cidr2ip -i cidrs.txt -o results.txt\n")
	fmt.Fprintf(os.Stderr, "         $ cat cidrs.txt | cidr2ip \n")
	fmt.Fprintf(os.Stderr, "--------------------------\nFlags:\n")
	flag.PrintDefaults()
}
