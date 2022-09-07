/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/gyarang/k-crypto/table"
	"github.com/spf13/cobra"
)

var (
	encryptKey string
	raw        string
	showTable  bool
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "암호화",
	Long: `encryptKey 값으로 암호표를 만들어 raw 값을 암호화하여 출력합니다.
example:
    k-crypto encrypt -k 금강산 -i "속히 상경하라" -t`,
	Run: func(cmd *cobra.Command, args []string) {
		table, err := table.NewTable(encryptKey)
		if err != nil {
			panic(err)
		}

		encrypted, err := table.Encrypt(raw)
		if err != nil {
			panic(err)
		}

		fmt.Println("암호화 값:")
		fmt.Println(encrypted)
		if showTable {
			fmt.Println("암호표:")
			table.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.PersistentFlags().StringVarP(&encryptKey, "key", "k", "금강산", "암호화 키값")
	encryptCmd.PersistentFlags().StringVarP(&raw, "input", "i", "속히 상경하라", "암호화 할 값(한글, 띄어쓰기만 지원)")
	encryptCmd.PersistentFlags().BoolVarP(&showTable, "table", "t", false, "암호표 출력 여부")
}
