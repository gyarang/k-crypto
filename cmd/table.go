/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gyarang/k-crypto/table"

	"github.com/spf13/cobra"
)

var (
	tableKey string
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "암호표 출력",
	Long: `encryptKey 값으로 암호표를 만들어 출력합니다.
example:
	k-crypto table -k 금강산`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := table.NewTable(tableKey)
		if err != nil {
			panic(err)
		}
		t.Print()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
	tableCmd.PersistentFlags().StringVarP(&tableKey, "key", "k", "금강산", "암호화 키값")
}
