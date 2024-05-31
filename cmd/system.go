/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "大数据基础平台节点系统信息.",
	Long: `大数据基础平台节点系统信息.

大数据基础平台节点系统信息:
1. 磁盘使用情况.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("system called")
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
