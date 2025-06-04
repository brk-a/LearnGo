package sub_domain_finder

import (
	"fmt"

	"github.com/miekg/dns"
)

func SubDomainFinder(){
	var msg dns.Msg
	domain := dns.Fqdn("youtube.com")
	msg.SetQuestion(domain, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err!=nil{
		panic(err)
	}
	if len(in.Answer)<1{
		fmt.Println("no records found")
		return
	}

	for _, answer:=range in.Answer{
		if a, ok := answer.(*dns.A); ok {
			fmt.Print("IP address found: ", a.A, "\n")
		}
	}
}