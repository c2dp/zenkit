/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDisk()
	},
}

func init() {
	systemCmd.AddCommand(diskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkDisk() {
	// 获取磁盘使用情况
	partition, _ := disk.Partitions(false)
	for _, part := range partition {
		usage, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("Disk %s: Total: %dGB, Used: %dGB\n", part.Mountpoint, usage.Total/1024/1024/1024, usage.Used/1024/1024/1024)
	}
}
