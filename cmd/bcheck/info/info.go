package info

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	BigdataRoleList = map[string][]string{
		"ZOOKEEPER":     {"ZKNODE"},
		"HADOOP":        {"NODEMANAGER", "ZKFC", "DATANODE", "HISTORYSERVER", "JOURNALNODE", "NAMENODE", "RESOURCEMANAGER"},
		"HBASE":         {"HMASTER", "HREGIONSERVER"},
		"KAFKA":         {"BROKER"},
		"ELASTICSEARCH": {"NODE"},
		"HBPWEB":        {"HBPWEBSERVER"},
	}

	RoleHost map[string]map[string]string

	ServiceRolePort = map[string]map[string]int{
		"ZOOKEEPER": {
			"ZKNODE": 2181,
		},
		"HADOOP": {
			"NAMENODE":        50070,
			"DATANODE":        50010,
			"RESOURCEMANAGER": 18088,
			"NODEMANAGER":     8042,
		},
		"HBASE": {
			"HMASTER":       60010,
			"HREGIONSERVER": 60020,
		},
		"KAFKA": {
			"BROKER": 9092,
		},
		"ELASTICSEARCH": {
			"NODE": 9200,
		},
		"HBPWEB": {
			"HBPWEBSERVER": 8878,
		},
	}
)

func GetHostsMap(settings *ClusterSettings) map[string]string {
	HostsMap := make(map[string]string)
	for _, host := range settings.Hosts {
		HostsMap[host.Hostname] = host.IP
	}
	if len(settings.Hosts) > 1 {
		fmt.Printf("集群类型: 集群\t 节点数量：%d", len(settings.Hosts))
	} else if len(settings.Hosts) == 1 {

		fmt.Printf("集群类型: 单机\t 节点数量：%d", len(settings.Hosts))
	} else {
		log.Fatalf("获取集群类型失败\n")
	}
	fmt.Printf("\n主机名IP映射关系: \n")
	return HostsMap
}

func GetServiceNodeHostname(settings *ClusterSettings, servicename string, role string) string {
	for _, v := range settings.Services {
		if v.Type == servicename {
			for _, v := range v.Roles {
				if v.Type == role {
					return v.Instances[0].Hostname
				}
			}
		}
	}
	return "null"
}

func GetServiceNodeMessage(settings *ClusterSettings, servicename string, rolelist []string) {
	for _, v := range settings.Services {
		if v.Type == servicename {
			fmt.Printf("服务：%s\n", v.Type)
			for _, v := range v.Roles {
				for _, role := range rolelist {
					if v.Type == role {
						fmt.Printf("%s角色类型数量：%d\n", role, len(v.Instances))
						for _, v := range v.Instances {
							fmt.Printf("角色类型：%s, 主机名：%s\n", role, v.Hostname)

						}
					}
				}
			}
		}
	}
}

func GetRoleNode(settings *ClusterSettings, servicename string, rolelist []string) *map[string]map[string]string {
	RoleHost = make(map[string]map[string]string)
	for _, v := range settings.Services {
		if v.Type == servicename {
			roleHost := make(map[string]string)
			for _, v := range v.Roles {
				for _, role := range rolelist {
					if v.Type == role {
						// fmt.Printf("%s角色类型数量：%d\n", role, len(v.Instances))
						for _, v := range v.Instances {
							// fmt.Printf("角色类型：%s, 主机名：%s\n", role, v.Hostname)
							role_idx := role + "_" + strconv.Itoa(v.Idx)
							roleHost[role_idx] = v.Hostname
						}
					}
				}
			}
			RoleHost[servicename] = roleHost
		}
	}

	return &RoleHost
}

func ScanPort(protocol string, hostname string, port int) bool {
	// fmt.Printf("scanning port %d \n", port)
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func CheckServicePort(service string, RoleNode map[string]map[string]string) {
	// fmt.Printf("服务: %v 端口是否监听：\n", service)
	for role, host := range RoleNode[service] {
		// fmt.Printf("role: %v, host: %v\n", role, host)
		newRole := strings.Split(role, "_")[0]
		if newRole != "ZKFC" {
			portNumber := ServiceRolePort[service][newRole]
			isUsed := ScanPort("tcp", host, portNumber)
			if !isUsed {
				fmt.Printf("角色:%v, 主机名：%v,端口：%v 没有监听，请在8877web上检查服务%v是否正常运行\n", role, host, portNumber, service)
			} else {
				fmt.Printf("角色:%v, 主机名：%v, 端口：%v 监听中\n", role, host, portNumber)
			}
		}

	}

}
