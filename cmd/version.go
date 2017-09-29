package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
)

func init() {
  RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Go Budget",
  Long:  `All software has versions. This is Go Budget's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Go Budget v0.0.0")
  },
}