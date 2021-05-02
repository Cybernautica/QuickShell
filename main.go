package main

import (
	"github.com/Cybernautica/QuickShell/cmd"
	a "github.com/Cybernautica/QuickShell/pkg/functions"
)

//var silent = flag.Bool("silent", false, "no banner")

func main() {
	// flag.Parse()
	// if *silent == true {
	// 	cmd.Execute()
	// } else {
	// 	a.Banner()
	// 	cmd.Execute()
	// }
	a.Banner()
	cmd.Execute()
}
