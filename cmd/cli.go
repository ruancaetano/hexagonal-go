/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ruancaetano/hexagonal-go/adapters/cli"
	"github.com/spf13/cobra"
)

var (
	action       string
	productId    string
	productName  string
	productPrice float64
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cli adapter",
	Long:  `Option to use cli adapter`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(productService, action, productId, productName, productPrice)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "get", "Set action: get|create|enable|disable")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Set product ID: required get|enable|disable")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "Set product Name:  required only for create")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Set product Price: required only for create")
}
