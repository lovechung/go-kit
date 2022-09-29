package bootstrap

import (
	"net"
	"os"
)

type ServiceInfo struct {
	Name     string
	Version  string
	Id       string
	Metadata map[string]string
}

func NewServiceInfo(name, version, id string) ServiceInfo {
	if id == "" {
		id = getIP()
		if id == "" {
			id, _ = os.Hostname()
		}
	}
	return ServiceInfo{
		Name:     name,
		Version:  version,
		Id:       id,
		Metadata: map[string]string{},
	}
}

func (s *ServiceInfo) GetInstanceId() string {
	return s.Id + "." + s.Name
}

func (s *ServiceInfo) SetMataData(k, v string) {
	s.Metadata[k] = v
}

func getIP() string {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range adds {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}

		}
	}
	return ""
}
