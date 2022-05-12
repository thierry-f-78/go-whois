package whois

import "context"
import "fmt"
import "io/ioutil"
import "net"
import "strings"
import "time"

var UnidentifiedTLD error = fmt.Errorf("UnidentifiedTLD")

func Whois(domain string)(string, error) {
	return WhoisContext(context.Background(), domain)
}

func WhoisContext(ctx context.Context, domain string)(string, error) {
	return WhoisContextDialer(ctx, &net.Dialer{}, domain)
}

func WhoisContextDialer(ctx context.Context, dialer *net.Dialer, domain string)(string, error) {
	var s *suffix_elt
	var l int
	var srv *server_elt
	var ok bool
	var query string
	var conn net.Conn
	var err error
	var d []byte
	var deadline time.Time

	domain = strings.ToLower(domain)

	/* Choose server according with request */
	l = len(domain)
	for _, s = range suffix_db {
		if l >= s.length {
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
	switch srv.kind {
	case 1:  query = "=" + domain
	//case "whois.denic.de":         query = "-T dn,ace " + domain
	default: query = domain
	}

	/* connect to server */
	conn, err = dialer.DialContext(ctx, "tcp", srv.name + ":43")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	/* Set deadline */
	deadline, ok = ctx.Deadline()
	if ok {
		err = conn.SetDeadline(deadline)
		if err != nil {
			return "", err
		}
	}

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

	/* Check not found cases */
	if srv.pattern != nil {
		if srv.pattern.Match(d) {
			return "", nil
		}
	}

	return string(d), nil
}
