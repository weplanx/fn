package common

type Inject struct {
	Values *Values
}

type Values struct {
	TrustedProxies []string `env:"trusted_proxies"`
}
