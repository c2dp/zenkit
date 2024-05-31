/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "zenkit工具版本.",
	Long:  `zenkit工具版本.`,
	Run: func(cmd *cobra.Command, args []string) {
		print(Format("v1.0.2", "2023-07-17"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")
	buildDate = strings.ReplaceAll(buildDate, "-", "")

	var dateStr string
	if buildDate != "" {
		dateStr = fmt.Sprintf("build%s", buildDate)
	}

	return fmt.Sprintf("zenkit version %s_%s\n", version, dateStr)
}
