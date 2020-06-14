package main

type config struct {
	Log       logObject
	Api       apiObject
	Dns       dnsObject
	Routing   routingObject
	Policy    policyObject
	Inbounds  []inboundObject
	Outbounds []outboundObject
	TransPort transportObject
	Stats     statsObject
	Reverse   reverseObject
}

type logObject struct {
	Access   string
	Error    string
	Loglevel string
}
type inboundObject struct {
	Port           interface{} // number | "env:variable" | string
	Listen         string
	Protocol       string
	Settings       interface{} // InboundConfigurationObject
	StreamSettings streamSettingsObject
	Tag            string
	Sniffing       sniffingObject
	Allocate       allocateObject
}
type sniffingObject struct {
	Enabled      bool
	DestOverride []string
}
type allocateObject struct {
	Strategy    string
	Refresh     uint16
	Concurrency uint16
}
type outboundObject struct {
	SendThrough    string
	Protocol       string
	Settings       interface{} // OutboundConfigurationObject
	Tag            string
	StreamSettings streamSettingsObject
	ProxySettings  proxySettingsObject
	Mx             muxObject
}
type proxySettingsObject struct {
	Tag string
}
type streamSettingsObject interface{}
type muxObject interface{}
type apiObject struct {
	Tag      string
	Services []string
}
type dnsObject struct {
	Hosts    map[string]string
	Servers  []serverObject // [string | ServerObject ]
	ClientIp string
	Tag      string
}
type serverObject struct {
	Address   string
	Port      uint16
	Domains   []string
	ExpectIPs []string
}
type routingObject struct {
	DomainsStrategy string // "AsIs" | "IPIfNonMatch" | "IPOnDemand"
	Rules           []ruleObject
	Balancers       []balancerObject
}
type ruleObject interface{}
type balancerObject interface{}
type policyObject interface{}
type transportObject interface{}
type statsObject interface{}
type reverseObject interface{}
