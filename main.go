/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	"github.com/Cybernautica/QUICKSHELL/cmd"
)

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

func main() {

	fmt.Println(banner)

	cmd.Execute()

}
