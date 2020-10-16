package defs

type Config struct {
	Port               int          `yaml:"port" json:"port"`
	SocksPort          int          `yaml:"socks-port" json:"socks-port"`
	AllowLan           bool         `yaml:"allow-lan" json:"allow-lan"`
	Mode               string       `yaml:"mode" json:"mode"`
	LogLevel           string       `yaml:"log-level" json:"log-level"`
	ExternalController string       `yaml:"external-controller" json:"external-controller"`
	Secret             string       `yaml:"secret" json:"secret"`
	Proxies            []Proxy      `yaml:"proxies" json:"proxies"`
	ProxyGroups        []ProxyGroup `yaml:"proxy-groups" json:"proxy-groups"`
	Rules              []string     `yaml:"rules" json:"rules"`
}

type Proxy struct {
	Name      string   `yaml:"name" json:"name"`
	Type      string   `yaml:"type" json:"type"`
	Server    string   `yaml:"server" json:"server"`
	Port      string   `yaml:"port" json:"port"`
	Uuid      string   `yaml:"uuid" json:"uuid"`
	AlterId   string   `yaml:"alterId" json:"alterId"`
	Cipher    string   `yaml:"cipher" json:"cipher"`
	Tls       bool     `yaml:"tls" json:"tls"`
	Network   string   `yaml:"network" json:"network"`
	WsPath    string   `yaml:"ws-path" json:"ws-path"`
	WsHeaders WsHeader `yaml:"ws-headers" json:"ws-headers"`
}

type ProxyGroup struct {
	Name     string   `yaml:"name" json:"name"`
	Proxies  []string `yaml:"proxies" json:"proxies"`
	Type     string   `yaml:"type" json:"type"`
	Url      string   `yaml:"url" json:"url"`
	Interval int      `yaml:"interval" json:"interval"`
}

type WsHeader struct {
	Host string `yaml:"Host" json:"Host"`
}
