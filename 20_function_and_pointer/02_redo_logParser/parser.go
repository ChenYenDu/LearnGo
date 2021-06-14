package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result // total visites per domain
	domains []string          //  unique domain names
	total   int               // total visits to all domains
	lines   int
}

func newParser() parser {
	// constructor pattern
	return parser{
		sum: make(map[string]result),
	}
}

func parse(p parser, line string) (parsed result, err error) {

	// var (
	// 	parsed result
	// 	err    error
	// )

	fields := strings.Fields(line)

	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		// return parsed, err
		return
	}

	// fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])
	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields[1], p.lines)
		// fmt.Println("Wrong input!", fields[1])
		// return parsed, err
		return
	}

	// return parsed, err
	return
}

func update(p parser, parsed result) parser {
	domain, visits := parsed.domain, parsed.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	p.lines++

	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}

	return p
}
