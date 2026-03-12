/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/luis13005/stresstest/cmd/stresstester"
	"github.com/spf13/cobra"
)

var (
	url string
)

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			panic(err)
		}

		repeticoes, err := cmd.Flags().GetInt("requests")
		if err != nil {
			panic(err)
		}

		concorrencias, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			panic(err)
		}

		stresstester.Tester(&url, repeticoes, concorrencias)
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.Flags().StringP("url", "u", "", "Escolha a url que deseja testar")
	stressCmd.Flags().IntP("requests", "r", 1, "Quantidade de repetições que será realizado o teste.")
	stressCmd.Flags().IntP("concurrency", "c", 1, "Quantidade de threads simultâneas.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
