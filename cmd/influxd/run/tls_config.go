package run

import (
	"crypto/tls"
	"fmt"
	"strings"
)

type TLSConfig struct {
	Ciphers    []string `toml:"ciphers"`
	MinVersion string   `toml:"min-version"`
	MaxVersion string   `toml:"max-version"`
}

func (c TLSConfig) Validate() error {
	_, err := c.Parse()
	return err
}

func (c TLSConfig) Parse() (out *tls.Config, err error) {
	if len(c.Ciphers) > 0 {
		if out == nil {
			out = new(tls.Config)
		}

		for _, name := range c.Ciphers {
			cipher, ok := ciphersMap[strings.ToUpper(name)]
			if !ok {
				return nil, unknownCipher(name)
			}
			out.CipherSuites = append(out.CipherSuites, cipher)
		}
	}

	if c.MinVersion != "" {
		if out == nil {
			out = new(tls.Config)
		}

		version, ok := versionsMap[strings.ToUpper(c.MinVersion)]
		if !ok {
			return nil, unknownVersion(c.MinVersion)
		}
		out.MinVersion = version
	}

	if c.MaxVersion != "" {
		if out == nil {
			out = new(tls.Config)
		}

		version, ok := versionsMap[strings.ToUpper(c.MaxVersion)]
		if !ok {
			return nil, unknownVersion(c.MaxVersion)
		}
		out.MaxVersion = version
	}

	return out, nil
}

var ciphersMap = map[string]uint16{
	"TLS_RSA_WITH_RC4_128_SHA":                tls.TLS_RSA_WITH_RC4_128_SHA,
	"TLS_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	"TLS_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	"TLS_RSA_WITH_AES_128_CBC_SHA256":         tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
	"TLS_RSA_WITH_AES_128_GCM_SHA256":         tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_RSA_WITH_AES_256_GCM_SHA384":         tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":        tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_RC4_128_SHA":          tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":     tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":   tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305":    tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305":  tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
}

func unknownCipher(name string) error {
	// TODO(jeff): list available?
	return fmt.Errorf("unknown cipher suite: %q", name)
}

var versionsMap = map[string]uint16{
	"SSL3.0": tls.VersionSSL30,
	"TLS1.0": tls.VersionTLS10,
	"1.0":    tls.VersionTLS11,
	"TLS1.1": tls.VersionTLS11,
	"1.1":    tls.VersionTLS11,
	"TLS1.2": tls.VersionTLS12,
	"1.2":    tls.VersionTLS12,
}

func unknownVersion(name string) error {
	// TODO(jeff): list available?
	return fmt.Errorf("unknown tls version: %q", name)
}
