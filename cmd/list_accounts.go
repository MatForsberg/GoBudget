package cmd

import (
	"github.com/MatForsberg/GoBudget/accounts"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listAccountsCmd)
}

var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range accounts.GetAll() {
			a.Print()
		}
	},
}
