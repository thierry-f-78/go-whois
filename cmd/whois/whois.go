package main

import "context"
import "flag"
import "net"
import "os"
import "time"

import "github.com/thierry-f-78/go-whois"

func main() {
	var domain string
	var res string
	var err error
	var ctx context.Context
	var dialer *net.Dialer

	flag.StringVar(&domain, "d", "", "domain")
	flag.Parse()

	if domain == "" {
		println("expect domain with -d")
		os.Exit(1)
	}

	ctx, _ = context.WithTimeout(context.Background(), 1000 * time.Millisecond)

	dialer = &net.Dialer{}

	res, err = whois.WhoisContextDialer(ctx, dialer, domain)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if res == "" {
		os.Stdout.Write([]byte("not found\n"))
	} else {
		os.Stdout.Write([]byte(res))
	}
}
