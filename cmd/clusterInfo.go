/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"github.com/c2dp/zenkit/cmd/common"
	"log"

	"github.com/spf13/cobra"
)

// clusterInfoCmd represents the clusterInfo command
var clusterInfoCmd = &cobra.Command{
	Use:   "clusterInfo",
	Short: "大数据集群信息",
	Long: `大数据集群信息.

1. 大数据基础平台集群节点DNS信息.
2. 大数据基础平台集群组件布局.`,
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {

		arg, err := cmd.Flags().GetString("module")
		err = cmd.ValidateArgs([]string{arg})
		if err != nil {
			log.Fatalf("Flag argument is invaild, vaild flag argument in %s.\nUse \"lightwork clusterInfo --help\" for more information about a command.", cmd.ValidArgs)
		}
		if arg == "all" {
			cSetting := common.GetSetting()
			getClusterIPMapping(cSetting)
			getServiceLayoutInfo(cSetting)
		}

	},
}

func init() {
	rootCmd.AddCommand(clusterInfoCmd)
	clusterInfoCmd.Flags().StringP("module", "m", "", "同时执行多个模块, 现只有一个参数: all, 即执行所有模块.")
	clusterInfoCmd.ValidArgs = []string{"all"}
}
