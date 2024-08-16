/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"to-go-list/service"
)

// listarCmd represents the listar command
var listarCmd = &cobra.Command{
	Use:   "listar",
	Short: "Utilizado para listar todas as tarefas pendentes(Ainda não concluídas) da sua lista.",
	Long: `

Você pode, opcionalmente, utilizar a flag -t para trazer todas as tarefas já criadas,
 independentmente de estarem concluídas ou não.`,
	Run: func(cmd *cobra.Command, args []string) {
		if(cmd.Flags().Changed("todos")) {
			service.ListarTarefas(false);
		} else {
			service.ListarTarefas(true)
		}
	},
}

func init() {
	rootCmd.AddCommand(listarCmd)
	listarCmd.Flags().BoolP("todos", "t",true,"Mostra todas as tarefas, independentemente do seu status.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
