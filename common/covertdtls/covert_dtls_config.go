package covertdtls

import (
	"strings"

	"github.com/theodorsm/covert-dtls/pkg/fingerprints"
)

type CovertDTLSConfig struct {
	Randomize   bool
	Mimic       bool
	Fingerprint fingerprints.ClientHelloFingerprint
}

func ParseConfigString(str string) CovertDTLSConfig {
	config := CovertDTLSConfig{}
	str = strings.ToLower(str)
	switch {
	case strings.Contains(str, "random"):
		config.Randomize = true
	case strings.Contains(str, "mimic"):
		config.Mimic = true
	default:
	}

	return config
}
