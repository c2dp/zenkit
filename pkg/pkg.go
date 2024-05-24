package pkg

import (
	"net"
)

func GetLocalAddr() ([]*net.IPNet, error) {
	var addrSet []*net.IPNet
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addrSet = append(addrSet, ipnet)
			}
		}
	}
	return addrSet, nil
}
