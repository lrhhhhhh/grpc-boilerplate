package config

type Config struct {
	Service struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
	} `json:"service"`
	TLS struct {
		CertFile string `json:"certFile"`
		KeyFile  string `json:"keyFile"`
	} `json:"tls"`
	Etcd struct {
		Addr string `json:"addr"`
	} `json:"etcd"`
}
