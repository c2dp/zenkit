package common

import (
	"github.com/c2dp/zenkit/cmd/bcheck/check"
	"github.com/c2dp/zenkit/cmd/bcheck/info"
	"github.com/c2dp/zenkit/cmd/bcheck/nssh"
	"log"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

func GetSetting() *info.ClusterSettings {
	conn, err := nssh.NoKeyClient("")
	if err != nil {
		log.Panicf("创建SSH客户端失败, ErrMsg: %s", err.Error())
	}
	defer conn.Close()

	var settings info.ClusterSettings
	// clusterSettingContent, _ := os.ReadFile("F:\\tmpdata\\cluster001.yml")
	clusterSettingContent := check.GetClusterSettings(conn)
	yaml.Unmarshal(clusterSettingContent, &settings)
	return &settings

}
func GetSSHClient() *ssh.Client {
	conn, err := nssh.NoKeyClient("")
	if err != nil {
		log.Panicf("创建SSH客户端失败, ErrMsg: %s", err.Error())
	}
	defer conn.Close()
	return conn
}
