package cmd

import (
	"github.com/MatForsberg/GoBudget/accounts"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createAccountCmd)
}

var createAccountCmd = &cobra.Command{
	Use:   "createAccount",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		account := accounts.Account{Name: args[0], Enabled: true}
		account.Create()
	},
}
