package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "GoBudget",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }