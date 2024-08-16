/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"to-go-list/service"
	"github.com/spf13/cobra"
)

// deletarCmd represents the deletar command
var deletarCmd = &cobra.Command{
	Use:   "deletar",
	Short: "Deleta um registro com o ID passado",
	Long: `Deleta um registro com o ID passado`,
	Run: func(cmd *cobra.Command, args []string) {
		service.DeletarTarefa(args[0])
		fmt.Println("Registro com o ID " + args[0] + " deletado com sucesso!!")
	},
}

func init() {
	rootCmd.AddCommand(deletarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
