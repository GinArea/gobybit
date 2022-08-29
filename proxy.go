package bybit

import (
	"fmt"
	"net/url"
)

func ParseProxy(proxy string) (*url.URL, error) {
	if proxy == "" {
		return nil, nil
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil ||
		(proxyURL.Scheme != "http" &&
			proxyURL.Scheme != "https" &&
			proxyURL.Scheme != "socks5") {
		if proxyURL, err := url.Parse("http://" + proxy); err == nil {
			return proxyURL, nil
		}
	}
	if err != nil {
		return nil, fmt.Errorf("invalid proxy %q: %v", proxy, err)
	}
	return proxyURL, nil
}
