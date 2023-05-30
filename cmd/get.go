/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-empty2/tools"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Image struct {
	Name string
	URL  string
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config") // 設定ファイル名 (config.yamlなど)
		viper.SetConfigType("yaml")   // 設定ファイルの種類 (yaml, jsonなど)
		viper.AddConfigPath(".")      // 設定ファイルのパス

		// 設定ファイルの読み込み
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("設定ファイルの読み込みに失敗しました:", err)
			return
		}

		var images []Image
		err = viper.UnmarshalKey("images", &images)
		if err != nil {
			fmt.Println("データの取得に失敗しました:", err)
			return
		}

		// データの表示
		for _, img := range images {
			tools.GetImage(img.URL, img.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
