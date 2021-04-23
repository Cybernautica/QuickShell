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
	payload() string
}

func (p program) payload() string {

	var pay string
	if p.language == "python" {

		r := rev.Python(p.ip, p.port)
		pay = r
	}

	return pay

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
		r := t.payload()
		fmt.Println(r)

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
	reverseshellCmd.Flags().StringP("lang", "l", "", "Programming language to generate rev shell")
	reverseshellCmd.MarkFlagRequired("ip")
	reverseshellCmd.MarkFlagRequired("port")
	reverseshellCmd.MarkFlagRequired("lang")

}
