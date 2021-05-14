package main

import (
	"flag"
	"fmt"
	"log"

	rev "github.com/Cybernautica/QuickShell/pkg/functions"
)

//var silent = flag.Bool("silent", false, "no banner")

func main() {

	ip := flag.String("ip", "", "Please provide your ip.")
	port := flag.String("port", "", "Please provide your port.")
	lang := flag.String("lang", "", "Please provide the language for the payload.")
	shell := flag.String("r", "", "Please provide the shell type. reverse/bind")
	silent := flag.Bool("s", false, "No Banner")
	flag.Parse()

	if *silent == true {
		if *ip == "" {
			log.Fatalf("Please provide the ip!")
		}
		if *port == "" {
			log.Fatalf("Please provide the port!")
		}
		if *lang == "" {
			log.Fatalf("Please provide the payload language!")
		}
		if *shell == "" {
			log.Fatalf("Please provude the shell type. reverse/bind")
		} else {
			switch *shell {
			case "reverse":
				cmd := rev.ReverseShell(*ip, *port, *lang)

				for _, v := range cmd {
					fmt.Println(v)
				}
			case "bind":
				log.Fatalf("[!!]Bind not supported yet!!")
			}
		}
	} else {
		rev.Banner()
		if *ip == "" {
			log.Fatalf("Please provide the ip!")
		}
		if *port == "" {
			log.Fatalf("Please provide the port!")
		}
		if *lang == "" {
			log.Fatalf("Please provide the payload language!")
		}
		if *shell == "" {
			log.Fatalf("Please provude the shell type. reverse/bind")
		} else {
			switch *shell {
			case "reverse":
				cmd := rev.ReverseShell(*ip, *port, *lang)

				for _, v := range cmd {
					fmt.Println(v)
				}
			case "bind":
				log.Fatalf("[!!]Bind not supported yet!!")
			}
		}
	}

}
