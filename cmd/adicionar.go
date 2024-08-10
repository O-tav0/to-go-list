/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adicionarCmd represents the adicionar command
var adicionarCmd = &cobra.Command{
	Use:   "adicionar",
	Short: "Utilizado para adicionar uma tarefa à sua lista",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adicionar called")
	},
}

func init() {
	rootCmd.AddCommand(adicionarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adicionarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adicionarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
