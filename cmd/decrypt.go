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
	decryptKey string
	encrypted  string
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "복호화",
	Long: `encryptKey 값을 이용하여 암호화 된 값을 복호화 합니다.
example:
    k-crypto decrypt -k 금강산 -i 11223344`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := table.NewTable(decryptKey)
		if err != nil {
			fmt.Println(err)
			return
		}

		decrypted, err := t.Decrypt(encrypted)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(table.AssembleDecryptedRunes(decrypted))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.PersistentFlags().StringVarP(&decryptKey, "key", "k", "금강산", "암호화 키값")
	decryptCmd.PersistentFlags().StringVarP(&encrypted, "input", "i", "속히 상경하라", "암호화 할 값(한글, 띄어쓰기만 지원)")
}
