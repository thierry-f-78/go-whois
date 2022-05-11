package main

import "flag"
import "os"

import "github.com/thierry-f-78/go-whois"

func main() {
	var domain string
	var res string
	var err error

	flag.StringVar(&domain, "d", "", "domain")
	flag.Parse()

	if domain == "" {
		println("expect domain with -d")
		os.Exit(1)
	}

	res, err = whois.Whois(domain)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Stdout.Write([]byte(res))
}
