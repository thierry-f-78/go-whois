package main

import "encoding/xml"
import "fmt"
import "net/http"
import "os"

import "golang.org/x/net/idna"

const max_size = 64 /* arbitrary */

type Whois_server struct {
	Host string `xml:"host,attr"`
}

type Domain struct {
	Name string `xml:"name,attr"`
	Whois_server []Whois_server `xml:"whoisServer"`
	Domain []Domain `xml:"domain"`
}

type Domain_list struct {
	Domain []Domain `xml:"domain"`
}

type Base struct {
	Len int `json:"length"`
	Pfx map[string][]string `json:""`
}

func dec_domain(d Domain) {
	var dd Domain
	var f string
	var list []string
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
			list = append(list, ws.Host)
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

	base = make([]Base, max_size)
	for i = 0; i < max_size; i++ {
		base[i].Len = i
		base[i].Pfx = make(map[string][]string)
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

/* File genearted with cmd/genlist. Do not edit */

type suffix_elt struct {
	length int
	suffix map[string][]string
}
var suffix_db []*suffix_elt = []*suffix_elt{
`)

	for i = max_size - 1; i >= 0; i-- {
		if len(base[i].Pfx) == 0 {
			continue
		}
		fmt.Printf(`	{
		length: %d,
		suffix: %#v,
	},
`, i, base[i].Pfx)
	}
	fmt.Printf(`}
`)
}
