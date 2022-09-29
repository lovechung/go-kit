package bootstrap

import (
	"net"
	"os"
)

type ServiceInfo struct {
	Env      string
	Name     string
	Version  string
	Id       string
	Metadata map[string]string
}

func NewServiceInfo(env, name, version, id string) ServiceInfo {
	if id == "" {
		id = GetLocalIP()
	}
	return ServiceInfo{
		Env:      env,
		Name:     name,
		Version:  version,
		Id:       id,
		Metadata: map[string]string{},
	}
}

func (s *ServiceInfo) GetInstanceId() string {
	hostname, _ := os.Hostname()
	return s.Id + "_" + hostname + s.Name
}

func (s *ServiceInfo) SetMataData(k, v string) {
	s.Metadata[k] = v
}

func GetLocalIP() string {
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
