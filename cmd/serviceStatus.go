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

// serviceStatusCmd represents the serviceStatus command
var serviceStatusCmd = &cobra.Command{
	Use:   "serviceStatus",
	Short: "大数据基础平台组件服务是否运行检查",
	Long: `通过端口是否监听判断服务是否与运行.
主要检测以下服务:
1. ZOOKEEPER.
2. HADOOP.
3. HBASE.
4. ELASTICSEARCH.
5. HBPWEB.
6. KAFKA.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cSetting := common.GetSetting()
			getServiceListen(cSetting)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceStatusCmd)

}

func getServiceListen(cSetting *info.ClusterSettings) {
	cSettings := common.GetSetting()
	var serviceRoleHost map[string]map[string]string
	serviceRoleHost = make(map[string]map[string]string)
	for k, v := range info.BigdataRoleList {
		fmt.Printf("\n服务: %v 端口是否监听：\n", k)
		serviceRoleHost = *info.GetRoleNode(cSettings, k, v)
		info.CheckServicePort(k, serviceRoleHost)
	}

}
