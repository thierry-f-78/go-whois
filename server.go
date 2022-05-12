package whois

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
	{
		length: 23,
		suffix: map[string]*server_elt{
			".xn--clchc0ea0b2g2a9gcd": &server_elt{
				name: "whois.sgnic.sg",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 18,
		suffix: map[string]*server_elt{
			".xn--mgbah1a3hjkrd": &server_elt{
				name: "whois.nic.mr",
				kind: 0,
				pattern: nil,
			},
			".xn--mgberp4a5d4ar": &server_elt{
				name: "whois.nic.net.sa",
				kind: 0,
				pattern: nil,
			},
			".xn--xkc2dl3a5ee0h": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 17,
		suffix: map[string]*server_elt{
			".xn--xkc2al3hye2a": &server_elt{
				name: "whois.nic.lk",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 16,
		suffix: map[string]*server_elt{
			".xn--lgbbat1ad8j": &server_elt{
				name: "whois.nic.dz",
				kind: 0,
				pattern: nil,
			},
			".xn--mgba3a4f16a": &server_elt{
				name: "whois.nic.ir",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 15,
		suffix: map[string]*server_elt{
			".xn--h2breg3eve": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbaam7a8h": &server_elt{
				name: "whois.aeda.net.ae",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbbh1a71e": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbx4cd0ab": &server_elt{
				name: "whois.mynic.my",
				kind: 0,
				pattern: nil,
			},
			".xn--rvc1e0am3e": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 14,
		suffix: map[string]*server_elt{
			".xn--fpcrj9c3d": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--fzc2c9e2c": &server_elt{
				name: "whois.nic.lk",
				kind: 0,
				pattern: nil,
			},
			".xn--h2brj9c8c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--yfro4i67o": &server_elt{
				name: "whois.sgnic.sg",
				kind: 0,
				pattern: nil,
			},
			".xn--ygbi2ammx": &server_elt{
				name: "whois.pnina.ps",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 13,
		suffix: map[string]*server_elt{
			".in-addr.arpa": &server_elt{
				name: "ARPA",
				kind: 0,
				pattern: nil,
			},
			".xn--3e0b707e": &server_elt{
				name: "whois.kr",
				kind: 0,
				pattern: nil,
			},
			".xn--45br5cyl": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--4dbrk0ce": &server_elt{
				name: "whois.isoc.org.il",
				kind: 0,
				pattern: nil,
			},
			".xn--mgb9awbf": &server_elt{
				name: "whois.registry.om",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbgu82a": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--ogbpf8fl": &server_elt{
				name: "whois.tld.sy",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 12,
		suffix: map[string]*server_elt{
			".xn--kprw13d": &server_elt{
				name: "whois.twnic.net.tw",
				kind: 0,
				pattern: nil,
			},
			".xn--kpry57d": &server_elt{
				name: "whois.twnic.net.tw",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbbh1a": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--mgbtx2b": &server_elt{
				name: "whois.cmc.iq",
				kind: 0,
				pattern: nil,
			},
			".xn--pgbs0dh": &server_elt{
				name: "whois.ati.tn",
				kind: 0,
				pattern: nil,
			},
			".xn--3hcrj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--gecrj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--h2brj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--s9brj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--j6w193g": &server_elt{
				name: "whois.hkirc.hk",
				kind: 0,
				pattern: nil,
			},
			".xn--2scrj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--45brj9c": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".xn--80ao21a": &server_elt{
				name: "whois.nic.kz",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 11,
		suffix: map[string]*server_elt{
			".xn--fiqz9s": &server_elt{
				name: "cwhois.cnnic.cn",
				kind: 0,
				pattern: nil,
			},
			".xn--o3cw4h": &server_elt{
				name: "whois.thnic.co.th",
				kind: 0,
				pattern: nil,
			},
			".xn--q7ce6a": &server_elt{
				name: "whois.nic.la",
				kind: 0,
				pattern: nil,
			},
			".xn--wgbh1c": &server_elt{
				name: "whois.dotmasr.eg",
				kind: 0,
				pattern: nil,
			},
			".xn--wgbl6a": &server_elt{
				name: "whois.registry.qa",
				kind: 0,
				pattern: nil,
			},
			".xn--y9a3aq": &server_elt{
				name: "whois.amnic.net",
				kind: 0,
				pattern: nil,
			},
			".xn--90a3ac": &server_elt{
				name: "whois.rnids.rs",
				kind: 0,
				pattern: nil,
			},
			".xn--fiqs8s": &server_elt{
				name: "cwhois.cnnic.cn",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 10,
		suffix: map[string]*server_elt{
			".xn--e1a4c": &server_elt{
				name: "whois.eu",
				kind: 0,
				pattern: nil,
			},
			".xn--j1amh": &server_elt{
				name: "whois.dotukr.com",
				kind: 0,
				pattern: nil,
			},
			".xn--qxa6a": &server_elt{
				name: "whois.eu",
				kind: 0,
				pattern: nil,
			},
			".solutions": &server_elt{
				name: "whois.nic.solutions",
				kind: 0,
				pattern: regexp.MustCompile("\\QDomain not found\\E"),
			},
			".e164.arpa": &server_elt{
				name: "whois.ripe.net",
				kind: 0,
				pattern: nil,
			},
			".xn--90ais": &server_elt{
				name: "whois.cctld.by",
				kind: 0,
				pattern: nil,
			},
			".xn--d1alf": &server_elt{
				name: "whois.marnet.mk",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 9,
		suffix: map[string]*server_elt{
			".ip6.arpa": &server_elt{
				name: "IP6",
				kind: 0,
				pattern: nil,
			},
			".xn--90ae": &server_elt{
				name: "whois.imena.bg",
				kind: 0,
				pattern: nil,
			},
			".xn--node": &server_elt{
				name: "whois.itdc.ge",
				kind: 0,
				pattern: nil,
			},
			".xn--p1ai": &server_elt{
				name: "whois.tcinet.ru",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 8,
		suffix: map[string]*server_elt{
			".priv.at": &server_elt{
				name: "whois.nic.priv.at",
				kind: 0,
				pattern: nil,
			},
			".jpn.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".web.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: regexp.MustCompile("\\QDOMAIN NOT FOUND\\E"),
			},
		},
	},
	{
		length: 7,
		suffix: map[string]*server_elt{
			".org.za": &server_elt{
				name: "org-whois.registry.net.za",
				kind: 0,
				pattern: nil,
			},
			".qc.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".us.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".za.net": &server_elt{
				name: "whois.za.net",
				kind: 0,
				pattern: nil,
			},
			".eu.org": &server_elt{
				name: "whois.eu.org",
				kind: 0,
				pattern: nil,
			},
			".za.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".com.ru": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".msk.ru": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".msk.su": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".in.net": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".se.net": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".uk.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".uy.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".web.za": &server_elt{
				name: "web-whois.registry.net.za",
				kind: 0,
				pattern: nil,
			},
			".gb.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".gb.net": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".gr.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".spb.ru": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".fed.us": &server_elt{
				name: "whois.nic.gov",
				kind: 0,
				pattern: nil,
			},
			".net.za": &server_elt{
				name: "net-whois.registry.net.za",
				kind: 0,
				pattern: nil,
			},
			".eu.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".museum": &server_elt{
				name: "whois.nic.museum",
				kind: 0,
				pattern: nil,
			},
			".net.ru": &server_elt{
				name: "whois.nic.net.ru",
				kind: 0,
				pattern: nil,
			},
			".nov.ru": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".gov.uk": &server_elt{
				name: "whois.ja.net",
				kind: 0,
				pattern: nil,
			},
			".alt.za": &server_elt{
				name: "whois.alt.za",
				kind: 0,
				pattern: nil,
			},
			".cn.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".de.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".edu.ru": &server_elt{
				name: "whois.informika.ru",
				kind: 0,
				pattern: nil,
			},
			".biz.ua": &server_elt{
				name: "whois.biz.ua",
				kind: 0,
				pattern: nil,
			},
			".org.ru": &server_elt{
				name: "whois.nic.net.ru",
				kind: 0,
				pattern: nil,
			},
			".sa.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".uk.net": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".travel": &server_elt{
				name: "whois.nic.travel",
				kind: 0,
				pattern: nil,
			},
			".edu.cn": &server_elt{
				name: "whois.edu.cn",
				kind: 0,
				pattern: nil,
			},
			".hu.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".no.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".nov.su": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".spb.su": &server_elt{
				name: "whois.flexireg.net",
				kind: 0,
				pattern: nil,
			},
			".gov.za": &server_elt{
				name: "whois.gov.za",
				kind: 0,
				pattern: nil,
			},
			".br.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".ru.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".se.com": &server_elt{
				name: "whois.centralnic.net",
				kind: 0,
				pattern: nil,
			},
			".za.org": &server_elt{
				name: "whois.za.org",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 6,
		suffix: map[string]*server_elt{
			".co.za": &server_elt{
				name: "whois.registry.net.za",
				kind: 0,
				pattern: nil,
			},
			".uk.co": &server_elt{
				name: "whois.uk.co",
				kind: 0,
				pattern: nil,
			},
			".co.pl": &server_elt{
				name: "whois.co.pl",
				kind: 0,
				pattern: nil,
			},
			".ac.ru": &server_elt{
				name: "whois.free.net",
				kind: 0,
				pattern: nil,
			},
			".pp.ru": &server_elt{
				name: "whois.nic.net.ru",
				kind: 0,
				pattern: nil,
			},
			".co.ua": &server_elt{
				name: "whois.co.ua",
				kind: 0,
				pattern: nil,
			},
			".pp.ua": &server_elt{
				name: "whois.pp.ua",
				kind: 0,
				pattern: nil,
			},
			".ac.za": &server_elt{
				name: "whois.ac.za",
				kind: 0,
				pattern: nil,
			},
			".co.ca": &server_elt{
				name: "whois.co.ca",
				kind: 0,
				pattern: nil,
			},
			".ac.uk": &server_elt{
				name: "whois.ja.net",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 5,
		suffix: map[string]*server_elt{
			".coop": &server_elt{
				name: "whois.nic.coop",
				kind: 0,
				pattern: nil,
			},
			".post": &server_elt{
				name: "whois.dotpostregistry.net",
				kind: 0,
				pattern: nil,
			},
			".aero": &server_elt{
				name: "whois.aero",
				kind: 0,
				pattern: nil,
			},
			".asia": &server_elt{
				name: "whois.nic.asia",
				kind: 0,
				pattern: nil,
			},
			".jobs": &server_elt{
				name: "whois.nic.jobs",
				kind: 1,
				pattern: nil,
			},
			".mobi": &server_elt{
				name: "whois.afilias.net",
				kind: 0,
				pattern: nil,
			},
			".name": &server_elt{
				name: "whois.nic.name",
				kind: 0,
				pattern: nil,
			},
			".arpa": &server_elt{
				name: "whois.iana.org",
				kind: 0,
				pattern: nil,
			},
			".info": &server_elt{
				name: "whois.afilias.net",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 4,
		suffix: map[string]*server_elt{
			".net": &server_elt{
				name: "whois.verisign-grs.com",
				kind: 1,
				pattern: regexp.MustCompile("\\QNo match for\\E"),
			},
			".org": &server_elt{
				name: "whois.pir.org",
				kind: 0,
				pattern: regexp.MustCompile("\\QNOT FOUND\\E"),
			},
			".edu": &server_elt{
				name: "whois.educause.edu",
				kind: 0,
				pattern: nil,
			},
			".gov": &server_elt{
				name: "whois.dotgov.gov",
				kind: 0,
				pattern: nil,
			},
			".pro": &server_elt{
				name: "whois.afilias.net",
				kind: 0,
				pattern: nil,
			},
			".xxx": &server_elt{
				name: "whois.registrar.adult",
				kind: 0,
				pattern: nil,
			},
			".xyz": &server_elt{
				name: "whois.nic.xyz",
				kind: 0,
				pattern: regexp.MustCompile("\\QDOMAIN NOT FOUND\\E"),
			},
			".com": &server_elt{
				name: "whois.verisign-grs.com",
				kind: 1,
				pattern: regexp.MustCompile("\\QNo match for\\E"),
			},
			".biz": &server_elt{
				name: "whois.nic.biz",
				kind: 0,
				pattern: regexp.MustCompile("\\QNo Data Found\\E"),
			},
			".cat": &server_elt{
				name: "whois.nic.cat",
				kind: 0,
				pattern: nil,
			},
			".tel": &server_elt{
				name: "whois.nic.tel",
				kind: 0,
				pattern: nil,
			},
			".int": &server_elt{
				name: "whois.iana.org",
				kind: 0,
				pattern: nil,
			},
		},
	},
	{
		length: 3,
		suffix: map[string]*server_elt{
			".mz": &server_elt{
				name: "whois.nic.mz",
				kind: 0,
				pattern: nil,
			},
			".br": &server_elt{
				name: "whois.registro.br",
				kind: 0,
				pattern: nil,
			},
			".de": &server_elt{
				name: "whois.denic.de",
				kind: 0,
				pattern: nil,
			},
			".dz": &server_elt{
				name: "whois.nic.dz",
				kind: 0,
				pattern: nil,
			},
			".ee": &server_elt{
				name: "whois.tld.ee",
				kind: 0,
				pattern: nil,
			},
			".md": &server_elt{
				name: "whois.nic.md",
				kind: 0,
				pattern: nil,
			},
			".mn": &server_elt{
				name: "whois.nic.mn",
				kind: 0,
				pattern: nil,
			},
			".mw": &server_elt{
				name: "whois.nic.mw",
				kind: 0,
				pattern: nil,
			},
			".tl": &server_elt{
				name: "whois.nic.tl",
				kind: 0,
				pattern: nil,
			},
			".aw": &server_elt{
				name: "whois.nic.aw",
				kind: 0,
				pattern: nil,
			},
			".bh": &server_elt{
				name: "whois.nic.bh",
				kind: 0,
				pattern: nil,
			},
			".gd": &server_elt{
				name: "whois.nic.gd",
				kind: 0,
				pattern: nil,
			},
			".lv": &server_elt{
				name: "whois.nic.lv",
				kind: 0,
				pattern: nil,
			},
			".pw": &server_elt{
				name: "whois.nic.pw",
				kind: 0,
				pattern: nil,
			},
			".ae": &server_elt{
				name: "whois.aeda.net.ae",
				kind: 0,
				pattern: nil,
			},
			".at": &server_elt{
				name: "whois.nic.at",
				kind: 0,
				pattern: nil,
			},
			".mg": &server_elt{
				name: "whois.nic.mg",
				kind: 0,
				pattern: nil,
			},
			".ng": &server_elt{
				name: "whois.nic.net.ng",
				kind: 0,
				pattern: nil,
			},
			".qa": &server_elt{
				name: "whois.registry.qa",
				kind: 0,
				pattern: nil,
			},
			".th": &server_elt{
				name: "whois.thnic.co.th",
				kind: 0,
				pattern: nil,
			},
			".co": &server_elt{
				name: "whois.nic.co",
				kind: 0,
				pattern: nil,
			},
			".cx": &server_elt{
				name: "whois.nic.cx",
				kind: 0,
				pattern: nil,
			},
			".it": &server_elt{
				name: "whois.nic.it",
				kind: 0,
				pattern: nil,
			},
			".li": &server_elt{
				name: "whois.nic.li",
				kind: 0,
				pattern: nil,
			},
			".re": &server_elt{
				name: "whois.nic.re",
				kind: 0,
				pattern: nil,
			},
			".tv": &server_elt{
				name: "tvwhois.verisign-grs.com",
				kind: 1,
				pattern: nil,
			},
			".ua": &server_elt{
				name: "whois.ua",
				kind: 0,
				pattern: nil,
			},
			".ag": &server_elt{
				name: "whois.nic.ag",
				kind: 0,
				pattern: nil,
			},
			".kw": &server_elt{
				name: "whois.nic.kw",
				kind: 0,
				pattern: nil,
			},
			".pm": &server_elt{
				name: "whois.nic.pm",
				kind: 0,
				pattern: nil,
			},
			".vg": &server_elt{
				name: "whois.nic.vg",
				kind: 0,
				pattern: nil,
			},
			".ar": &server_elt{
				name: "whois.nic.ar",
				kind: 0,
				pattern: nil,
			},
			".cz": &server_elt{
				name: "whois.nic.cz",
				kind: 0,
				pattern: nil,
			},
			".ge": &server_elt{
				name: "whois.nic.ge",
				kind: 0,
				pattern: nil,
			},
			".gp": &server_elt{
				name: "whois.nic.gp",
				kind: 0,
				pattern: nil,
			},
			".ro": &server_elt{
				name: "whois.rotld.ro",
				kind: 0,
				pattern: nil,
			},
			".sm": &server_elt{
				name: "whois.nic.sm",
				kind: 0,
				pattern: nil,
			},
			".gg": &server_elt{
				name: "whois.gg",
				kind: 0,
				pattern: nil,
			},
			".io": &server_elt{
				name: "whois.nic.io",
				kind: 0,
				pattern: regexp.MustCompile("\\QNOT FOUND\\E"),
			},
			".sl": &server_elt{
				name: "whois.nic.sl",
				kind: 0,
				pattern: nil,
			},
			".bi": &server_elt{
				name: "whois1.nic.bi",
				kind: 0,
				pattern: nil,
			},
			".cm": &server_elt{
				name: "whois.netcom.cm",
				kind: 0,
				pattern: nil,
			},
			".ml": &server_elt{
				name: "whois.dot.ml",
				kind: 0,
				pattern: nil,
			},
			".by": &server_elt{
				name: "whois.cctld.by",
				kind: 0,
				pattern: nil,
			},
			".nz": &server_elt{
				name: "whois.srs.net.nz",
				kind: 0,
				pattern: nil,
			},
			".uk": &server_elt{
				name: "whois.nic.uk",
				kind: 0,
				pattern: nil,
			},
			".bj": &server_elt{
				name: "whois.nic.bj",
				kind: 0,
				pattern: nil,
			},
			".kr": &server_elt{
				name: "whois.kr",
				kind: 0,
				pattern: nil,
			},
			".kg": &server_elt{
				name: "whois.kg",
				kind: 0,
				pattern: nil,
			},
			".sx": &server_elt{
				name: "whois.sx",
				kind: 0,
				pattern: nil,
			},
			".td": &server_elt{
				name: "whois.nic.td",
				kind: 0,
				pattern: nil,
			},
			".la": &server_elt{
				name: "whois.nic.la",
				kind: 0,
				pattern: nil,
			},
			".rs": &server_elt{
				name: "whois.rnids.rs",
				kind: 0,
				pattern: nil,
			},
			".ru": &server_elt{
				name: "whois.tcinet.ru",
				kind: 0,
				pattern: nil,
			},
			".sg": &server_elt{
				name: "whois.sgnic.sg",
				kind: 0,
				pattern: nil,
			},
			".uz": &server_elt{
				name: "whois.cctld.uz",
				kind: 0,
				pattern: nil,
			},
			".id": &server_elt{
				name: "whois.id",
				kind: 0,
				pattern: nil,
			},
			".se": &server_elt{
				name: "whois.iis.se",
				kind: 0,
				pattern: nil,
			},
			".sh": &server_elt{
				name: "whois.nic.sh",
				kind: 0,
				pattern: nil,
			},
			".su": &server_elt{
				name: "whois.tcinet.ru",
				kind: 0,
				pattern: nil,
			},
			".bm": &server_elt{
				name: "whois.afilias-srs.net",
				kind: 0,
				pattern: nil,
			},
			".hr": &server_elt{
				name: "whois.dns.hr",
				kind: 0,
				pattern: nil,
			},
			".so": &server_elt{
				name: "whois.nic.so",
				kind: 0,
				pattern: nil,
			},
			".zm": &server_elt{
				name: "whois.zicta.zm",
				kind: 0,
				pattern: nil,
			},
			".ws": &server_elt{
				name: "whois.website.ws",
				kind: 0,
				pattern: nil,
			},
			".gi": &server_elt{
				name: "whois2.afilias-grs.net",
				kind: 0,
				pattern: nil,
			},
			".gs": &server_elt{
				name: "whois.nic.gs",
				kind: 0,
				pattern: nil,
			},
			".ht": &server_elt{
				name: "whois.nic.ht",
				kind: 0,
				pattern: nil,
			},
			".ms": &server_elt{
				name: "whois.nic.ms",
				kind: 0,
				pattern: nil,
			},
			".pr": &server_elt{
				name: "whois.afilias-srs.net",
				kind: 0,
				pattern: nil,
			},
			".tc": &server_elt{
				name: "whois.nic.tc",
				kind: 0,
				pattern: nil,
			},
			".to": &server_elt{
				name: "whois.tonic.to",
				kind: 0,
				pattern: nil,
			},
			".ai": &server_elt{
				name: "whois.nic.ai",
				kind: 0,
				pattern: regexp.MustCompile("\\QNo Object Found\\E"),
			},
			".hn": &server_elt{
				name: "whois.nic.hn",
				kind: 0,
				pattern: nil,
			},
			".nc": &server_elt{
				name: "whois.nc",
				kind: 0,
				pattern: nil,
			},
			".rw": &server_elt{
				name: "whois.ricta.org.rw",
				kind: 0,
				pattern: nil,
			},
			".as": &server_elt{
				name: "whois.nic.as",
				kind: 0,
				pattern: nil,
			},
			".ax": &server_elt{
				name: "whois.ax",
				kind: 0,
				pattern: nil,
			},
			".ch": &server_elt{
				name: "whois.nic.ch",
				kind: 0,
				pattern: nil,
			},
			".gy": &server_elt{
				name: "whois.registry.gy",
				kind: 0,
				pattern: nil,
			},
			".is": &server_elt{
				name: "whois.isnic.is",
				kind: 0,
				pattern: nil,
			},
			".jp": &server_elt{
				name: "whois.jprs.jp",
				kind: 0,
				pattern: nil,
			},
			".lk": &server_elt{
				name: "whois.nic.lk",
				kind: 0,
				pattern: nil,
			},
			".bn": &server_elt{
				name: "whois.bnnic.bn",
				kind: 0,
				pattern: nil,
			},
			".ca": &server_elt{
				name: "whois.cira.ca",
				kind: 0,
				pattern: nil,
			},
			".im": &server_elt{
				name: "whois.nic.im",
				kind: 0,
				pattern: nil,
			},
			".kn": &server_elt{
				name: "whois.nic.kn",
				kind: 0,
				pattern: nil,
			},
			".st": &server_elt{
				name: "whois.nic.st",
				kind: 0,
				pattern: nil,
			},
			".vc": &server_elt{
				name: "whois2.afilias-grs.net",
				kind: 0,
				pattern: nil,
			},
			".ac": &server_elt{
				name: "whois.nic.ac",
				kind: 0,
				pattern: nil,
			},
			".lu": &server_elt{
				name: "whois.dns.lu",
				kind: 0,
				pattern: nil,
			},
			".sy": &server_elt{
				name: "whois.tld.sy",
				kind: 0,
				pattern: nil,
			},
			".gf": &server_elt{
				name: "whois.mediaserv.net",
				kind: 0,
				pattern: nil,
			},
			".iq": &server_elt{
				name: "whois.cmc.iq",
				kind: 0,
				pattern: nil,
			},
			".ls": &server_elt{
				name: "whois.nic.ls",
				kind: 0,
				pattern: nil,
			},
			".me": &server_elt{
				name: "whois.nic.me",
				kind: 0,
				pattern: nil,
			},
			".pe": &server_elt{
				name: "kero.yachay.pe",
				kind: 0,
				pattern: nil,
			},
			".sk": &server_elt{
				name: "whois.sk-nic.sk",
				kind: 0,
				pattern: nil,
			},
			".ss": &server_elt{
				name: "whois.nic.ss",
				kind: 0,
				pattern: nil,
			},
			".tg": &server_elt{
				name: "whois.nic.tg",
				kind: 0,
				pattern: nil,
			},
			".bw": &server_elt{
				name: "whois.nic.net.bw",
				kind: 0,
				pattern: nil,
			},
			".bz": &server_elt{
				name: "AFILIAS",
				kind: 0,
				pattern: nil,
			},
			".cr": &server_elt{
				name: "whois.nic.cr",
				kind: 0,
				pattern: nil,
			},
			".dk": &server_elt{
				name: "whois.dk-hostmaster.dk",
				kind: 0,
				pattern: nil,
			},
			".fm": &server_elt{
				name: "whois.nic.fm",
				kind: 0,
				pattern: nil,
			},
			".mq": &server_elt{
				name: "whois.mediaserv.net",
				kind: 0,
				pattern: nil,
			},
			".tz": &server_elt{
				name: "whois.tznic.or.tz",
				kind: 0,
				pattern: nil,
			},
			".cn": &server_elt{
				name: "whois.cnnic.cn",
				kind: 0,
				pattern: nil,
			},
			".fo": &server_elt{
				name: "whois.nic.fo",
				kind: 0,
				pattern: nil,
			},
			".mu": &server_elt{
				name: "whois.nic.mu",
				kind: 0,
				pattern: nil,
			},
			".my": &server_elt{
				name: "whois.mynic.my",
				kind: 0,
				pattern: nil,
			},
			".tk": &server_elt{
				name: "whois.dot.tk",
				kind: 0,
				pattern: nil,
			},
			".il": &server_elt{
				name: "whois.isoc.org.il",
				kind: 0,
				pattern: nil,
			},
			".mk": &server_elt{
				name: "whois.marnet.mk",
				kind: 0,
				pattern: nil,
			},
			".mm": &server_elt{
				name: "whois.registry.gov.mm",
				kind: 0,
				pattern: nil,
			},
			".pt": &server_elt{
				name: "whois.dns.pt",
				kind: 0,
				pattern: nil,
			},
			".sa": &server_elt{
				name: "whois.nic.net.sa",
				kind: 0,
				pattern: nil,
			},
			".tm": &server_elt{
				name: "whois.nic.tm",
				kind: 0,
				pattern: nil,
			},
			".uy": &server_elt{
				name: "whois.nic.org.uy",
				kind: 0,
				pattern: nil,
			},
			".au": &server_elt{
				name: "whois.auda.org.au",
				kind: 0,
				pattern: nil,
			},
			".do": &server_elt{
				name: "whois.nic.do",
				kind: 0,
				pattern: nil,
			},
			".fi": &server_elt{
				name: "whois.fi",
				kind: 0,
				pattern: nil,
			},
			".in": &server_elt{
				name: "whois.registry.in",
				kind: 0,
				pattern: nil,
			},
			".mr": &server_elt{
				name: "whois.nic.mr",
				kind: 0,
				pattern: nil,
			},
			".nl": &server_elt{
				name: "whois.domain-registry.nl",
				kind: 0,
				pattern: nil,
			},
			".si": &server_elt{
				name: "whois.register.si",
				kind: 0,
				pattern: nil,
			},
			".tn": &server_elt{
				name: "whois.ati.tn",
				kind: 0,
				pattern: nil,
			},
			".cc": &server_elt{
				name: "ccwhois.verisign-grs.com",
				kind: 1,
				pattern: nil,
			},
			".fj": &server_elt{
				name: "whois.usp.ac.fj",
				kind: 0,
				pattern: nil,
			},
			".gl": &server_elt{
				name: "whois.nic.gl",
				kind: 0,
				pattern: nil,
			},
			".ir": &server_elt{
				name: "whois.nic.ir",
				kind: 0,
				pattern: nil,
			},
			".ki": &server_elt{
				name: "whois.nic.ki",
				kind: 0,
				pattern: nil,
			},
			".ps": &server_elt{
				name: "whois.pnina.ps",
				kind: 0,
				pattern: nil,
			},
			".sn": &server_elt{
				name: "whois.nic.sn",
				kind: 0,
				pattern: nil,
			},
			".us": &server_elt{
				name: "whois.nic.us",
				kind: 0,
				pattern: nil,
			},
			".nf": &server_elt{
				name: "whois.nic.nf",
				kind: 0,
				pattern: nil,
			},
			".am": &server_elt{
				name: "whois.amnic.net",
				kind: 0,
				pattern: nil,
			},
			".be": &server_elt{
				name: "whois.dns.be",
				kind: 0,
				pattern: nil,
			},
			".dm": &server_elt{
				name: "whois.nic.dm",
				kind: 0,
				pattern: nil,
			},
			".ie": &server_elt{
				name: "whois.weare.ie",
				kind: 0,
				pattern: nil,
			},
			".ky": &server_elt{
				name: "whois.kyregistry.ky",
				kind: 0,
				pattern: nil,
			},
			".lt": &server_elt{
				name: "whois.domreg.lt",
				kind: 0,
				pattern: nil,
			},
			".ma": &server_elt{
				name: "whois.registre.ma",
				kind: 0,
				pattern: nil,
			},
			".nu": &server_elt{
				name: "whois.iis.nu",
				kind: 0,
				pattern: nil,
			},
			".sb": &server_elt{
				name: "whois.nic.net.sb",
				kind: 0,
				pattern: nil,
			},
			".ve": &server_elt{
				name: "whois.nic.ve",
				kind: 0,
				pattern: nil,
			},
			".sc": &server_elt{
				name: "whois2.afilias-grs.net",
				kind: 0,
				pattern: nil,
			},
			".af": &server_elt{
				name: "whois.nic.af",
				kind: 0,
				pattern: nil,
			},
			".bg": &server_elt{
				name: "whois.register.bg",
				kind: 0,
				pattern: nil,
			},
			".ci": &server_elt{
				name: "whois.nic.ci",
				kind: 0,
				pattern: nil,
			},
			".gh": &server_elt{
				name: "whois.nic.gh",
				kind: 0,
				pattern: nil,
			},
			".je": &server_elt{
				name: "whois.je",
				kind: 0,
				pattern: nil,
			},
			".no": &server_elt{
				name: "whois.norid.no",
				kind: 0,
				pattern: nil,
			},
			".om": &server_elt{
				name: "whois.registry.om",
				kind: 0,
				pattern: nil,
			},
			".vu": &server_elt{
				name: "whois.dnrs.neustar",
				kind: 0,
				pattern: nil,
			},
			".bo": &server_elt{
				name: "whois.nic.bo",
				kind: 0,
				pattern: nil,
			},
			".pl": &server_elt{
				name: "whois.dns.pl",
				kind: 0,
				pattern: nil,
			},
			".ug": &server_elt{
				name: "whois.co.ug",
				kind: 0,
				pattern: nil,
			},
			".tr": &server_elt{
				name: "whois.nic.tr",
				kind: 0,
				pattern: nil,
			},
			".fr": &server_elt{
				name: "whois.nic.fr",
				kind: 0,
				pattern: regexp.MustCompile("\\QNo entries found\\E"),
			},
			".hk": &server_elt{
				name: "whois.hkirc.hk",
				kind: 0,
				pattern: nil,
			},
			".hm": &server_elt{
				name: "whois.registry.hm",
				kind: 0,
				pattern: nil,
			},
			".ke": &server_elt{
				name: "whois.kenic.or.ke",
				kind: 0,
				pattern: nil,
			},
			".ly": &server_elt{
				name: "whois.nic.ly",
				kind: 0,
				pattern: nil,
			},
			".mx": &server_elt{
				name: "whois.mx",
				kind: 0,
				pattern: nil,
			},
			".tf": &server_elt{
				name: "whois.nic.tf",
				kind: 0,
				pattern: nil,
			},
			".tw": &server_elt{
				name: "whois.twnic.net.tw",
				kind: 0,
				pattern: nil,
			},
			".wf": &server_elt{
				name: "whois.nic.wf",
				kind: 0,
				pattern: nil,
			},
			".yt": &server_elt{
				name: "whois.nic.yt",
				kind: 0,
				pattern: nil,
			},
			".ec": &server_elt{
				name: "whois.nic.ec",
				kind: 0,
				pattern: nil,
			},
			".gq": &server_elt{
				name: "whois.dominio.gq",
				kind: 0,
				pattern: nil,
			},
			".kz": &server_elt{
				name: "whois.nic.kz",
				kind: 0,
				pattern: nil,
			},
			".pf": &server_elt{
				name: "whois.registry.pf",
				kind: 0,
				pattern: nil,
			},
			".cl": &server_elt{
				name: "whois.nic.cl",
				kind: 0,
				pattern: nil,
			},
			".eu": &server_elt{
				name: "whois.eu",
				kind: 0,
				pattern: nil,
			},
			".hu": &server_elt{
				name: "whois.nic.hu",
				kind: 0,
				pattern: nil,
			},
			".lc": &server_elt{
				name: "whois2.afilias-grs.net",
				kind: 0,
				pattern: nil,
			},
			".na": &server_elt{
				name: "whois.na-nic.com.na",
				kind: 0,
				pattern: nil,
			},
		},
	},
}
