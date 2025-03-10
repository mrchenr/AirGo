package model

// vmess 格式
type Vmess struct {
	V            string `json:"v" `   //
	Name         string `json:"ps"`   //节点名
	Address      string `json:"add"`  //节点地址
	Port         string `json:"port"` //端口
	Uuid         string `json:"id"`   //用户UUID
	Aid          string `json:"aid"`  //额外ID
	Net          string `json:"net"`  //传输协议
	Disguisetype string `json:"type"` //伪装类型
	Host         string `json:"host"` //伪装域名
	Path         string `json:"path"` //
	Tls          string `json:"tls"`  //传输层安全
	Alpn         string `json:"alpn"`
	Fp           string `json:"fp"`
	Sni          string `json:"sni"`
}

// clash meta  yaml格式
type ClashYaml struct {
	Port               int64             `yaml:"port"`
	SocksPort          int64             `yaml:"socks-port"`
	RedirPort          int64             `yaml:"redir-port"`
	AllowLan           bool              `yaml:"allow-lan"`
	Mode               string            `yaml:"mode"`
	LogLevel           string            `yaml:"log-level"`
	ExternalController string            `yaml:"external-controller"`
	Secret             string            `yaml:"secret"`
	Proxies            []ClashProxy      `yaml:"proxies"`
	ProxyGroups        []ClashProxyGroup `yaml:"proxy-groups"`
	Rules              []string          `yaml:"rules"`
}
type ClashProxy struct {
	//基础参数
	Name    string `yaml:"name" json:"name"`
	Type    string `yaml:"type" json:"type"`
	Server  string `yaml:"server" json:"server"`
	Port    int    `yaml:"port" json:"port"`
	Uuid    string `yaml:"uuid" json:"uuid"`
	Network string `yaml:"network" json:"network"`
	Udp     bool   `yaml:"udp" json:"udp"`
	//vmess参数
	Alterid int64  `yaml:"alterId" json:"alterId"`
	Cipher  string `yaml:"cipher" json:"cipher"`
	//trojan 参数
	Password string `yaml:"password" json:"password"`
	//vless流控
	Flow string `yaml:"flow" json:"flow"`

	Tls               bool        `yaml:"tls" json:"tls"`
	Sni               string      `yaml:"sni" json:"sni"`
	ClientFingerprint string      `yaml:"client-fingerprint" json:"client-fingerprint"` //Available: "chrome","firefox","safari","ios","random", currently only support TLS transport in TCP/GRPC/WS/HTTP for VLESS/Vmess and trojan.
	Alpn              []string    `yaml:"alpn" json:"alpn"`                             //h2 http/1.1
	Servername        string      `yaml:"servername" json:"servername"`                 //REALITY servername
	SkipCertVerify    bool        `yaml:"skip-cert-verify" json:"skip-cert-verify"`
	WsOpts            WsOpts      `yaml:"ws-opts" json:"ws-opts"`
	RealityOpts       RealityOpts `yaml:"reality-opts" json:"reality-opts"`
	GrpcOpts          GrpcOpts    `yaml:"grpc-opts" json:"grpc-opts"`
	H2Opts            H2Opts      `yaml:"h2-opts" json:"h2-opts"`
}

type WsOpts struct {
	Path                string            `yaml:"path"`
	Headers             map[string]string `yaml:"headers"`
	MaxEarlyData        int               `yaml:"max-early-data"`         //2048
	EarlyDataHeaderName string            `yaml:"early-data-header-name"` //Sec-WebSocket-Protocol
}
type WsHeaders struct {
	Host string `yaml:"Host"`
}

type RealityOpts struct {
	PublicKey string `yaml:"public-key"`
	ShortID   string `yaml:"short-id"`
}
type GrpcOpts struct {
	GrpcServiceName string `yaml:"grpc-service-name"` //grpc
}

type H2Opts struct {
	Host []string `yaml:"host"`
	Path string   `yaml:"path"`
}
type HttpOpts struct {
	Method  string                `yaml:"method"` //GET
	Path    map[string][]string   `yaml:"path"`
	Headers map[string]Connection `yaml:"headers"`
}
type Connection []string

type ClashProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxies  []string `yaml:"proxies"`
	Url      string   `yaml:"url"`
	Interval int      `yaml:"interval"`
}

// ios shadowrocket
type Shadowrocket struct {
	Host          string  `json:"host"`
	File          string  `json:"file"`
	ObfsParam     string  `json:"obfsParam"`
	Alpn          string  `json:"alpn"`
	Cert          string  `json:"cert"`
	Created       float64 `json:"created"`
	Updated       float64 `json:"updated"`
	Tls           bool    `json:"tls"`
	Mtu           string  `json:"mtu"`
	Flag          string  `json:"flag"`
	PrivateKey    string  `json:"privateKey"`
	Hpkp          string  `json:"hpkp"`
	Uuid          string  `json:"uuid"`
	Type          string  `json:"type"`
	Downmbps      string  `json:"downmbps"`
	User          string  `json:"user"`
	Xtls          int64   `json:"xtls"`
	Plugin        string  `json:"plugin"`
	Method        string  `json:"method"`
	Data          string  `json:"data"`
	Udp           int64   `json:"udp"`
	Filter        string  `json:"filter"`
	ProtoParam    string  `json:"protoParam"`
	Reserved      string  `json:"reserved"`
	AlterId       string  `json:"alterId"`
	Upmbps        string  `json:"upmbps"`
	Keepalive     string  `json:"keepalive"`
	AllowInsecure int64   `json:"allowInsecure"`
	Port          string  `json:"port"`
	Obfs          string  `json:"obfs"`
	Dns           string  `json:"dns"`
	PublicKey     string  `json:"publicKey"`
	Peer          string  `json:"peer"`
	Weight        int64   `json:"weight"`
	Title         string  `json:"title"`
	Proto         string  `json:"proto"`
	Password      string  `json:"password"`
	ShortId       string  `json:"shortId"`
	Chain         string  `json:"chain"`
	Ip            string  `json:"ip"`
}

// ios surge
type SurgeConf struct {
	General    General    `ini:"General"`
	Replica    Replica    `ini:"Replica"`
	Panel      Panel      `ini:"Panel"`
	Proxy      Proxy      `ini:"Proxy"`
	ProxyGroup ProxyGroup `ini:"Proxy Group"`
	Rule       Rule       `ini:"Rule"`
}
type General struct {
	Loglevel            string   `ini:"loglevel""`
	Doh_server          string   `ini:"doh-server "`
	Dns_server          []string `ini:"dns-server"`
	Tun_excluded_routes []string `ini:"tun-excluded-routes"`
	Skip_proxy          []string `ini:"skip-proxy"`

	Wifi_assist             bool   `ini:"wifi-assist"`
	Allow_wifi_access       bool   `ini:"allow-wifi-access"`
	Wifi_access_http_port   int64  `ini:"wifi-access-http-port"`
	Wifi_access_socks5_port int64  `ini:"wifi-access-socks5-port"`
	Http_listen             string `ini:"http-listen"`
	Socks5_listen           string `ini:"socks5-listen"`

	External_controller_access string `ini:"external-controller-access"`
	Replica                    bool   `ini:"replica"`

	Tls_provider             string `ini:"tls-provider"`
	Network_framework        bool   `ini:"network-framework"`
	Exclude_simple_hostnames bool   `ini:"exclude-simple-hostnames"`
	Ipv6                     bool   `ini:"ipv6"`

	Test_timeout      int64  `ini:"test-timeout"`
	Proxy_test_url    string `ini:"proxy-test-url"`
	Geoip_maxmind_url string `ini:"geoip-maxmind-url"`
}
type Replica struct {
	Hide_apple_request       bool `ini:"hide-apple-request "`
	Hide_crashlytics_request bool `ini:"hide-crashlytics-request"`
	Use_keyword_filter       bool `ini:"use-keyword-filter"`
	Hide_udp                 bool `ini:"hide-udp "`
}
type Panel struct {
	SubscribeInfo string `ini:"SubscribeInfo"`
}
type Proxy struct {
	ProxyText string `ini:"ProxyText"`
}
type ProxyGroup struct {
	Proxy    []string `ini:"Proxy"`
	Auto     []string `ini:"auto"`
	Fallback []string `ini:"fallback"`
}
type Rule struct {
	RuleText string `ini:"RuleText"`
}
