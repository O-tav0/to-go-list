/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"to-go-list/service"
)

// completarCmd represents the completar command
var completarCmd = &cobra.Command{
	Use:   "completar",
	Short: "Completa uma task com o ID passado",
	Long: `Completa uma task com o ID passado`,
	Run: func(cmd *cobra.Command, args []string) {
		service.CompletarTarefa(args[0])
		fmt.Println("Tarefa de ID " + args[0] + " completada!!!")
	},
}

func init() {
	rootCmd.AddCommand(completarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
