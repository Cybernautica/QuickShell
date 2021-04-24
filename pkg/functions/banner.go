package functions

import "fmt"

const banner = `  

		██████╗ ██╗   ██╗██╗ ██████╗██╗  ██╗███████╗██╗  ██╗███████╗██╗     ██╗     
		██╔═══██╗██║   ██║██║██╔════╝██║ ██╔╝██╔════╝██║  ██║██╔════╝██║     ██║     
		██║   ██║██║   ██║██║██║     █████╔╝ ███████╗███████║█████╗  ██║     ██║     
		██║▄▄ ██║██║   ██║██║██║     ██╔═██╗ ╚════██║██╔══██║██╔══╝  ██║     ██║     
		╚██████╔╝╚██████╔╝██║╚██████╗██║  ██╗███████║██║  ██║███████╗███████╗███████╗
		╚══▀▀═╝  ╚═════╝ ╚═╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝

                     A CLI TOOL TO GENERATE QUICK REVERSE SHELLS
With this tool you can generate easy and sophisticated reverse or bind shell commands to help you during penetration tests.

	                 Don't waste more time with .txt files storing your Reverse shells!

                                                                             `

func Banner() {

	fmt.Println(banner)

}
