package whois

import "fmt"
import "io/ioutil"
import "net"

var UnidentifiedTLD error = fmt.Errorf("UnidentifiedTLD")

func Whois(domain string)(string, error) {
	var s *suffix_elt
	var l int
	var srv []string
	var ok bool
	var query string
	var conn net.Conn
	var err error
	var d []byte

	/* Choose server according with request */
	l = len(domain)
	for _, s = range suffix_db {
		if (l > s.length && domain[l - s.length - 1] == '.') || l == s.length {
			srv, ok = s.suffix[domain[l - s.length:]]
			if ok {
				break
			}
		}
	}

	/* whois server not found */
	if srv == nil {
		return "", UnidentifiedTLD
	}

	/* Build query according with some specific servers */
	switch srv[0] {
	case "whois.verisign-grs.com": query = "=" + domain
	case "whois.denic.de":         query = "-T dn,ace " + domain
	default:                       query = domain
	}

	/* connect to server */
	conn, err = net.Dial("tcp", srv[0] + ":43")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	/* perform request */
	_, err = conn.Write([]byte(query + "\r\n"))
	if err != nil {
		return "", err
	}

	/* read response */
	d, err = ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
