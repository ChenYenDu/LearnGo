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
	lerr    error
}

func newParser() *parser {
	// constructor pattern
	return &parser{
		sum: make(map[string]result),
	}
}

func parse(p *parser, line string) (r result) {

	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)

	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		// return parsed, err
		return
	}

	var err error

	// fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])
	r.domain = fields[0]
	r.visits, err = strconv.Atoi(fields[1])

	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields[1], p.lines)
		// fmt.Println("Wrong input!", fields[1])
		// return parsed, err
		// return
	}

	// return parsed, err
	return
}

func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	p.total += r.visits

	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
		// map elements are not addressable values
	}
	clone := p.sum[r.domain]
	_ = &clone
}

func err(p *parser) error {
	return p.lerr
}
