package main

import "bufio"
import "fmt"
import "net/http"
import "os"
import "strings"

const max_size = 64 /* arbitrary */

type Srv struct {
	Name string
	Kind int /* 0: normal, 1: verisign */
	Regex string
}

type Base struct {
	Len int
	Pfx map[string]*Srv
}

var base []Base

func main() {
	var res *http.Response
	var err error
	var i int
	var s *Srv
	var k string
	var scanner *bufio.Scanner
	var line string
	var pos int
	var b4_len int
	var tokens []string
	var domain string

	base = make([]Base, max_size)
	for i = 0; i < max_size; i++ {
		base[i].Len = i
		base[i].Pfx = make(map[string]*Srv)
	}

	/* This control has no time control. Press ctrl+c if the time is too long */
	res, err = http.Get("https://raw.githubusercontent.com/rfc1036/whois/next/tld_serv_list")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	/* read line by line */
	scanner = bufio.NewScanner(res.Body)
	for scanner.Scan() {

		line = scanner.Text()

		/* Remove # .* $ */
		pos = strings.Index(line, "#")
		if pos != -1 {
			line = line[:pos]
		}

		/* trim spaces */
		line = strings.TrimSpace(line)

		/* ignore empty */
		if line == "" {
			continue
		}

		/* only one space as separator */
		line = strings.ReplaceAll(line, "\t", " ")
		for {
			b4_len = len(line)
			line = strings.ReplaceAll(line, "  ", " ")
			if b4_len == len(line) {
				break
			}
		}

		/* split by tokens */
		tokens = strings.Split(line, " ")

		/* decode token according with thier length */
		if len(tokens) == 2 {
			if tokens[1] == "WEB" {
				continue
			}
			if tokens[1] == "NONE" {
				continue
			}
			s = &Srv{}
			s.Name = tokens[1]
			s.Kind = 0
			domain = tokens[0]
		} else if len(tokens) == 3 {
			if tokens[1] == "WEB" {
				continue
			}
			if tokens[1] == "NONE" {
				continue
			}
			s = &Srv{}
			s.Name = tokens[2]
			s.Kind = 0
			if tokens[1] == "VERISIGN" {
				s.Kind = 1
			}
			domain = tokens[0]
		} else {
			continue
		}

		/* Specific pattern */
		switch domain {
		case ".fr":
			s.Regex = `\QNo entries found\E`
		case ".io", ".org":
			s.Regex = `\QNOT FOUND\E`
		case ".com", ".net":
			s.Regex = `\QNo match for\E`
		case ".biz":
			s.Regex = `\QNo Data Found\E`
		case ".web.com":
			s.Regex = `\QDOMAIN NOT FOUND\E`
		}

		/* index entry */
		if len(domain) > max_size - 1 {
			panic("Domain '" + domain + "' is too long, please increase const 'max_size'")
		}

		base[len(domain)].Pfx[domain] = s
	}

	fmt.Printf(`package whois

import "regexp"

/* File genearted with cmd/genlist. Do not edit */

type server_elt struct {
	name string
	kind int /* 0: normal, 1: verisign */
	pattern *regexp.Regexp
}

type suffix_elt struct {
	length int
	suffix map[string]*server_elt
}
var suffix_db []*suffix_elt = []*suffix_elt{
`)

	for i = max_size - 1; i >= 0; i-- {
		if len(base[i].Pfx) == 0 {
			continue
		}
		fmt.Printf("\t{\n")
		fmt.Printf("\t\tlength: %d,\n", i)
		fmt.Printf("\t\tsuffix: map[string]*server_elt{\n")
		for k, s = range base[i].Pfx {
			fmt.Printf("\t\t\t%#v: &server_elt{\n", k)
			fmt.Printf("\t\t\t\tname: %#v,\n", s.Name)
			fmt.Printf("\t\t\t\tkind: %d,\n", s.Kind)
			if s.Regex != "" {
				fmt.Printf("\t\t\t\tpattern: regexp.MustCompile(%#v),\n", s.Regex)
			} else {
				fmt.Printf("\t\t\t\tpattern: nil,\n")
			}
			fmt.Printf("\t\t\t},\n")
		}
		fmt.Printf("\t\t},\n")
		fmt.Printf("\t},\n")
	}
	fmt.Printf("}\n")
}
