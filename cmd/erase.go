/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/dchauviere/vault-pass-auth-helper/pkg/pass"
	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		storage := pass.Pass{
			BasePath:    passBasePath,
			DefaultName: passDefaultName,
		}
		storage.Erase(os.Getenv("VAULT_ADDR"), os.Getenv("VAULT_NAMESPACE"))
	},
}

func init() {
	rootCmd.AddCommand(eraseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eraseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eraseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
