package check

import (
	"fmt"
	"github.com/c2dp/zenkit/cmd/bcheck/info"
	"log"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

func GetClusterSettings(conn *ssh.Client) []byte {
	/*
	* 获取cluster*.yml配置文件内容
	 */
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	clusterSeetingContent, err := session.CombinedOutput("cat /usr/lib/cloudmanager/components/lark/data/clusters/*/*.yml")
	if err != nil {
		log.Fatal("unable get cluster*.yml  status", err)
	}
	return clusterSeetingContent

}

func IsHdfsSafeMode(conn *ssh.Client) {
	/*
	* HDFS是否进入安全模式检测模块。
	 */
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	fmt.Printf("\n\nHDFS 安全模式检测：\n结果判断条件：\nHDFS进入安全模式，则只能读取HDFS上的文件，不能把数据写入HDFS.\n  1. 开启，则是进入了安全模式;\n  2. 关闭，则是没有进入安全模式。\n开始检测，检测时间较长，请耐心等待！\n开始时间：%s", time.Now().Local())
	SafeModeStatus, err := session.CombinedOutput("/usr/lib/*/SERVICE-HADOOP-*/bin/hdfs dfsadmin -safemode get")
	Res := "开启"
	if err != nil {
		log.Fatal("unable get safemode status", err.Error())
	}
	if strings.Contains(string(SafeModeStatus), "ON") {
		fmt.Printf("\n结束时间：%s\n检测结果：%s\n", time.Now().Local(), Res)
	}
	Res = "关闭"
	fmt.Printf("\n结束时间：%s\n检测结果：%s\n", time.Now().Local(), Res)

}

func HdfsHealthy(conn *ssh.Client) {
	/*
	* HDFS文件系统状态报告检测模块。
	 */
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	fmt.Printf("\n\nHDFS 数据坏块检测：\n结果判断条件：\nHDFS是否存在数据坏块.\n  1. 有，则存在数据坏块;\n  2. 没有，则是不存在数据坏块。\n开始检测，检测时间较长，请耐心等待！\n开始时间：%s", time.Now().Local())
	Res := "有"
	dfsHealth, err := session.CombinedOutput("/usr/lib/*/SERVICE-HADOOP-*/bin/hdfs fsck / | grep -i Status")
	if err != nil {
		log.Fatal("unable get dfsReport status", err)
	}
	if strings.Contains(string(dfsHealth), "HEALTHY") {
		Res = "没有"

	}
	fmt.Printf("\n结束时间：%s\n检测结果：%s\n", time.Now().Local(), Res)
}

func HdfsReport(conn *ssh.Client) {
	/*
	* HDFS文件系统状态报告检测模块。
	 */
	fmt.Printf("HDFS检测报告:\n")
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	dfsReport, err := session.CombinedOutput("/usr/lib/*/SERVICE-HADOOP-*/bin/hdfs dfsadmin -report")
	if err != nil {
		log.Fatal("unable get dfsReport status", err)
	}
	fmt.Println(string(dfsReport))

}

func HBaseFsck(conn *ssh.Client) {
	/*
	* HBase 数据不一致检测模块。
	 */
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	fmt.Printf("\n\nHBase 数据不一致检测：\n结果判断条件：\n  1. 数量不为0，则是有数据不一致;\n  2. 数量为0，则是没有数据不一致。\n开始检测，检测时间较长，请耐心等待！\n开始时间：%s", time.Now().Local())
	dfsReport, err := session.CombinedOutput("/usr/lib/*/SERVICE-HBASE-*/bin/hbase hbck | grep 'inconsistencies detected.'| grep -v grep")
	if err != nil {
		log.Fatal("unable get hbase fsck status", err)
	}

	pattern, _ := regexp.Compile(`[0-9]+ inconsistencies detected.`)
	temp := pattern.Find(dfsReport)
	tmppattern, _ := regexp.Compile(`[0-9]+`)

	fmt.Printf("\n结束时间：%s\n检测结果：%s\n", time.Now().Local(), tmppattern.Find([]byte(temp)))
}

func ElasticsearchCheck(conn *ssh.Client, esnode string) {
	/*
	* Elasticsearch 检测模块。
	 */
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	var esreport info.ElasticsearchStruct
	cmd := "curl --location --request GET 'http://localhost:9200/_cluster/health'"
	newcmd := strings.Replace(cmd, "localhost", esnode, -1)
	EsReport, err := session.Output(newcmd)
	if err != nil {
		log.Fatal("unable get elasticsearch health status: ", err)
	}
	yaml.Unmarshal(EsReport, &esreport)
	fmt.Printf("\n\nElasticsearch健康值检测:\n结果判断条件: \n1. red, 则是Elasticsearch集群不健康, 有部分主分片未正常分配, 请联系大数据技术支持排查根因;\n2. yellow, 则是Elasticsearch集群不健康, 有部分副本分片未正常分配,能够正常读写数据,可以联系大数据技术支持排查下原因;\n3. green, 则是Elasticsearch集群健康, 能正常对外提供查询服务。\n开始检测, 检测时间较长, 请耐心等待！\n开始时间: %s", time.Now().Local())
	fmt.Printf("\n结束时间：%s\n检测结果：%s\n", time.Now().Local(), esreport.Status)

}
