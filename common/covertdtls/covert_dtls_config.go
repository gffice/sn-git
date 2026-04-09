package covertdtls

import (
	"strings"

	"github.com/theodorsm/covert-dtls/pkg/fingerprints"
)

const (
	// CovertDTLSConfigRandomize is a config string used for CovertDTLSConfig to enable ClientHello randomization.
	CovertDTLSConfigRandomize = "randomize"
	// CovertDTLSConfigMimic is a config string used for CovertDTLSConfig to enable ClientHello mimicking of the latest Chrome or Firefox version.
	CovertDTLSConfigMimic = "mimic"
	// CovertDTLSConfigMimic is a config string used for CovertDTLSConfig to enable ClientHello mimicking of a random Chrome or Firefox fingerprint.
	CovertDTLSConfigRandomizeMimic = "randomizemimic"
)

// CovertDTLSConfig is used to configure the covert-dtls library for fingerprint-resistance for the DTLS handshake.
// If mimic is enabled, the ClientHello of the latest version Chrome or Firefox is mimicked.
// If randomize is enabled, the ClientHello is randomized and it's fingerprint will be unique every time.
// If both mimic and randomize is enabled, a random ClientHello message of Chrome or Firefox is mimicked.
// Fingerprint is used to set the ClientHello to be mimicked with a hex string.
type CovertDTLSConfig struct {
	Randomize   bool
	Mimic       bool
	Fingerprint fingerprints.ClientHelloFingerprint
}

// ParseConfigString creates a CovertDTLSConfig from a config string.
// Valid configurations strings are: mimic, randomize, randomizemimic.
// Using randomizemimic is recommended for best stability and decent fingerprint-resistance, randomize is less stable.
// All other strings will return an empty config with every feature disabled.
func ParseConfigString(str string) CovertDTLSConfig {
	config := CovertDTLSConfig{}
	str = strings.ToLower(str)
	switch str {
	case CovertDTLSConfigRandomize:
		config.Randomize = true
	case CovertDTLSConfigMimic:
		config.Mimic = true
	case CovertDTLSConfigRandomizeMimic:
		config.Randomize = true
		config.Mimic = true
	default:
	}

	return config
}
