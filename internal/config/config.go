package config

type Config struct {
	Service struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
	} `json:"service"`

	Discovery struct {
		Etcd struct {
			Addr string `json:"addr"`
		} `json:"etcd"`
		Enable bool `json:"enable"`
	} `json:"etcd"`

	TLS struct {
		CertFile           string `json:"certFile"`
		KeyFile            string `json:"keyFile"`
		ServerNameOverride string `json:"serverNameOverride"`
		Enable             bool   `json:"enable"`
	} `json:"tls"`
	Gateway struct {
		TLS struct {
			Enable   bool   `json:"enable"`
			CertFile string `json:"certFile"`
			KeyFile  string `json:"keyFile"`
		} `json:"tls"`
		OnlyUnaryRPC bool   `json:"onlyUnaryRPC"`
		Addr         string `json:"addr"`     // gateway addr
		Endpoint     string `json:"endpoint"` // gateway 代理的 gRPC 服务器地址
		Enable       bool   `json:"enable"`
	}
}
