/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"fmt"
	"github.com/c2dp/zenkit/cmd/bcheck/info"
	"github.com/c2dp/zenkit/cmd/common"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "大数据基础平台集群节点DNS信息.",
	Long: `大数据基础平台组件集群节点IP映射关系.

大数据基础平台集群节点IP与主机名的映射关系.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cSetting := common.GetSetting()
			getClusterIPMapping(cSetting)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	clusterInfoCmd.AddCommand(dnsCmd)
}

func getClusterIPMapping(settings *info.ClusterSettings) {
	fmt.Printf("\n大数据集群节点信息获取:\n")
	for hostname, ip := range info.GetHostsMap(settings) {
		color.Set(color.BgBlue)
		fmt.Printf("IP:%s\t\tHostname:%s", ip, hostname)
		color.Unset()
		fmt.Printf("\n")
	}
}
