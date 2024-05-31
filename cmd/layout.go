/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"fmt"
	"github.com/c2dp/zenkit/cmd/bcheck/info"
	"github.com/c2dp/zenkit/cmd/common"

	"github.com/spf13/cobra"
)

// layoutCmd represents the layout command
var layoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "大数据基础平台集群组件布局.",
	Long: `获取大数据基础平台集群组件布局.

包含大数据基础平台所有组件的分布情况.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cSetting := common.GetSetting()
			getServiceLayoutInfo(cSetting)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	clusterInfoCmd.AddCommand(layoutCmd)

}

func getServiceLayoutInfo(settings *info.ClusterSettings) {
	fmt.Printf("\n大数据基础平台集群组件布局:\n\n")
	for _, srv := range settings.Services {
		fmt.Printf("SERVICE: %s\n", srv.Type)
		for _, role := range srv.Roles {
			fmt.Printf("ROLE: %s\n", role.Type)
			for _, instance := range role.Instances {
				fmt.Printf("InstanceId: %s\t\tHostname: %s\n", instance.ID, instance.Hostname)

			}
		}
		fmt.Println("")
	}
}
