/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k-crypto",
	Short: "일본이 쩔쩔맨 K-암호 Go 구현체",
	Long: `k-crypto 는 일제시대 김재명이 고안한 암호체계의 Go 구현체입니다.
참조: https://www.hani.co.kr/arti/society/society_general/1057247.html`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}