package main

import "encoding/xml"
import "fmt"
import "net/http"
import "os"

import "golang.org/x/net/idna"

const max_size = 64 /* arbitrary */

type Whois_server struct {
	Host string `xml:"host,attr"`
	Pattern string `xml:"availablePattern"`
}

type Domain struct {
	Name string `xml:"name,attr"`
	Whois_server []Whois_server `xml:"whoisServer"`
	Domain []Domain `xml:"domain"`
}

type Domain_list struct {
	Domain []Domain `xml:"domain"`
}

type Srv struct {
	Name string
	Regex string
}

type Base struct {
	Len int
	Pfx map[string][]*Srv
}

func dec_domain(d Domain) {
	var dd Domain
	var f string
	var list []*Srv
	var ws Whois_server
	var err error

	for { /* not a loop, just a breakable statement */

		/* convert utf8 tu punny code, ignore errors */
		f, err = idna.ToASCII(d.Name)
		if err != nil {
			break
		}

		/* Ignore when no whois server */
		list = nil
		for _, ws = range d.Whois_server {
			if ws.Host == "" {
				continue
			}
			list = append(list, &Srv{Name: ws.Host, Regex: ws.Pattern})
		}
		if len(list) <= 0 {
			break
		}

		if len(f) > max_size - 1 {
			panic("Domain '" + f + "' is too long, please increase const 'max_size'")
		}

		base[len(f)].Pfx[f] = list
		break
	}

	/* process subdomains */
	for _, dd = range d.Domain {
		dec_domain(dd)
	}
}

var base []Base

func main() {
	var res *http.Response
	var xml_dec *xml.Decoder
	var err error
	var d Domain_list
	var dd Domain
	var i int
	var ss []*Srv
	var s *Srv
	var k string

	base = make([]Base, max_size)
	for i = 0; i < max_size; i++ {
		base[i].Len = i
		base[i].Pfx = make(map[string][]*Srv)
	}

	/* This control has no time control. Press ctrl+c if the time is too long */
	res, err = http.Get("https://raw.githubusercontent.com/whois-server-list/whois-server-list/master/whois-server-list.xml")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	xml_dec = xml.NewDecoder(res.Body)
	err = xml_dec.Decode(&d)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, dd = range d.Domain {
		dec_domain(dd)
	}

	fmt.Printf(`package whois

import "regexp"

/* File genearted with cmd/genlist. Do not edit */

type server_elt struct {
	name string
	pattern *regexp.Regexp
}

type suffix_elt struct {
	length int
	suffix map[string][]*server_elt
}
var suffix_db []*suffix_elt = []*suffix_elt{
`)

	for i = max_size - 1; i >= 0; i-- {
		if len(base[i].Pfx) == 0 {
			continue
		}
		fmt.Printf("\t{\n")
		fmt.Printf("\t\tlength: %d,\n", i)
		fmt.Printf("\t\tsuffix: map[string][]*server_elt{\n")
		for k, ss = range base[i].Pfx {
			fmt.Printf("\t\t\t%#v: []*server_elt{\n", k)
			for _, s = range ss {
				fmt.Printf("\t\t\t\t{\n")
				fmt.Printf("\t\t\t\t\tname: %#v,\n", s.Name)
				if s.Regex != "" {
					fmt.Printf("\t\t\t\t\tpattern: regexp.MustCompile(%#v),\n", s.Regex)
				} else {
					fmt.Printf("\t\t\t\t\tpattern: nil,\n")
				}
				fmt.Printf("\t\t\t\t},\n")
			}
			fmt.Printf("\t\t\t},\n")
		}
		fmt.Printf("\t\t},\n")
		fmt.Printf("\t},\n")
	}
	fmt.Printf("}\n")
}
