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
		id = GetLocalIP()
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

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
