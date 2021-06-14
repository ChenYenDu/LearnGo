package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {

	p := parser{
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++
		fields := strings.Fields(in.Text())

		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}

		// fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields[1], p.lines)
			// fmt.Println("Wrong input!", fields[1])
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		p.total += visits

		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
