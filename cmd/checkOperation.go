/*
Copyright © 2022 chenchao30 <chenchao30@hikvision.com.cn>
*/
package cmd

import (
	"github.com/c2dp/zenkit/cmd/bcheck/check"
	"github.com/c2dp/zenkit/cmd/bcheck/info"
	"github.com/c2dp/zenkit/cmd/bcheck/nssh"
	"github.com/c2dp/zenkit/cmd/common"
	"log"

	"github.com/spf13/cobra"
)

// checkOperationCmd represents the checkOperation command
var checkOperationCmd = &cobra.Command{
	Use:   "checkOperation",
	Short: "对组件是否正常运行进行检测.",
	Long: `检测组件是否正常运行.

主要的检测功能如下:
1.HDFS 是否进入安全模式.
2.HDFS 是否存在数据坏块.
3. HBASE是否存在数据不一致.
4. ELASTICSEARCH集群健康状态.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cSetting := common.GetSetting()
			getServiceState(cSetting)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	rootCmd.AddCommand(checkOperationCmd)

}

func getServiceState(cSetting *info.ClusterSettings) {
	conn, err := nssh.NoKeyClient("")
	if err != nil {
		log.Panicf("创建SSH客户端失败, ErrMsg: %s", err.Error())
	}
	defer conn.Close()
	cSettings := common.GetSetting()
	var esnode string
	for _, srv := range cSettings.Services {
		if srv.Type == "ELASTICSEARCH" {
			esnode = srv.Roles[0].HostnameByRoles[0]
		}
	}
	check.IsHdfsSafeMode(conn)
	check.HdfsHealthy(conn)
	check.HBaseFsck(conn)
	check.ElasticsearchCheck(conn, esnode)

}
