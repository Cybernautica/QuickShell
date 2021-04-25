package cmd

import (
	"fmt"

	rev "github.com/Cybernautica/QUICKSHELL/pkg/functions"
	"github.com/spf13/cobra"
)

type program struct {
	ip       string
	port     string
	language string
}

type reverseShell interface {
	payload()
}

func (p program) payload() {

	i := rev.ReverseShell(p.ip, p.port, p.language)

	for _, val := range i {
		fmt.Print("\n")
		fmt.Println(val)
	}

}

// reverseshellCmd represents the reverseshell command
var reverseshellCmd = &cobra.Command{
	Use:   "reverseshell",
	Short: "Generate quick reverse shell ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetString("port")
		lang, _ := cmd.Flags().GetString("lang")

		var t reverseShell = program{ip: ip, port: port, language: lang}

		t.payload()

	},
}

func init() {
	rootCmd.AddCommand(reverseshellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//reverseshellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	reverseshellCmd.Flags().StringP("ip", "i", "", "Provide the IP Adddress")
	reverseshellCmd.Flags().StringP("port", "p", "", "Provide the port Number")
	reverseshellCmd.Flags().StringP("lang", "l", "", "Programming language or binary name to generate rev shell")
	reverseshellCmd.MarkFlagRequired("ip")
	reverseshellCmd.MarkFlagRequired("port")
	reverseshellCmd.MarkFlagRequired("lang")

}
