// Code generated by cdpgen. DO NOT EDIT.

package webauthn

// AuthenticatorID
type AuthenticatorID string

// AuthenticatorProtocol
type AuthenticatorProtocol string

// AuthenticatorProtocol as enums.
const (
	AuthenticatorProtocolNotSet AuthenticatorProtocol = ""
	AuthenticatorProtocolU2F    AuthenticatorProtocol = "u2f"
	AuthenticatorProtocolCTAP2  AuthenticatorProtocol = "ctap2"
)

func (e AuthenticatorProtocol) Valid() bool {
	switch e {
	case "u2f", "ctap2":
		return true
	default:
		return false
	}
}

func (e AuthenticatorProtocol) String() string {
	return string(e)
}

// CTAP2Version
type CTAP2Version string

// CTAP2Version as enums.
const (
	CTAP2VersionNotSet  CTAP2Version = ""
	CTAP2VersionCTAP2_0 CTAP2Version = "ctap2_0"
	CTAP2VersionCTAP2_1 CTAP2Version = "ctap2_1"
)

func (e CTAP2Version) Valid() bool {
	switch e {
	case "ctap2_0", "ctap2_1":
		return true
	default:
		return false
	}
}

func (e CTAP2Version) String() string {
	return string(e)
}

// AuthenticatorTransport
type AuthenticatorTransport string

// AuthenticatorTransport as enums.
const (
	AuthenticatorTransportNotSet   AuthenticatorTransport = ""
	AuthenticatorTransportUSB      AuthenticatorTransport = "usb"
	AuthenticatorTransportNFC      AuthenticatorTransport = "nfc"
	AuthenticatorTransportBLE      AuthenticatorTransport = "ble"
	AuthenticatorTransportCable    AuthenticatorTransport = "cable"
	AuthenticatorTransportInternal AuthenticatorTransport = "internal"
)

func (e AuthenticatorTransport) Valid() bool {
	switch e {
	case "usb", "nfc", "ble", "cable", "internal":
		return true
	default:
		return false
	}
}

func (e AuthenticatorTransport) String() string {
	return string(e)
}

// VirtualAuthenticatorOptions
type VirtualAuthenticatorOptions struct {
	Protocol                    AuthenticatorProtocol  `json:"protocol"`                              // No description.
	CTAP2Version                CTAP2Version           `json:"ctap2Version,omitempty"`                // Defaults to ctap2_0. Ignored if |protocol| == u2f.
	Transport                   AuthenticatorTransport `json:"transport"`                             // No description.
	HasResidentKey              *bool                  `json:"hasResidentKey,omitempty"`              // Defaults to false.
	HasUserVerification         *bool                  `json:"hasUserVerification,omitempty"`         // Defaults to false.
	HasLargeBlob                *bool                  `json:"hasLargeBlob,omitempty"`                // If set to true, the authenticator will support the largeBlob extension. https://w3c.github.io/webauthn#largeBlob Defaults to false.
	AutomaticPresenceSimulation *bool                  `json:"automaticPresenceSimulation,omitempty"` // If set to true, tests of user presence will succeed immediately. Otherwise, they will not be resolved. Defaults to true.
	IsUserVerified              *bool                  `json:"isUserVerified,omitempty"`              // Sets whether User Verification succeeds or fails for an authenticator. Defaults to false.
}

// Credential
type Credential struct {
	CredentialID         string  `json:"credentialId"`         // No description.
	IsResidentCredential bool    `json:"isResidentCredential"` // No description.
	RPID                 *string `json:"rpId,omitempty"`       // Relying Party ID the credential is scoped to. Must be set when adding a credential.
	PrivateKey           []byte  `json:"privateKey"`           // The ECDSA P-256 private key in PKCS#8 format. (Encoded as a base64 string when passed over JSON)
	UserHandle           []byte  `json:"userHandle,omitempty"` // An opaque byte sequence with a maximum size of 64 bytes mapping the credential to a specific user. (Encoded as a base64 string when passed over JSON)
	SignCount            int     `json:"signCount"`            // Signature counter. This is incremented by one for each successful assertion. See https://w3c.github.io/webauthn/#signature-counter
	LargeBlob            []byte  `json:"largeBlob,omitempty"`  // The large blob associated with the credential. See https://w3c.github.io/webauthn/#sctn-large-blob-extension (Encoded as a base64 string when passed over JSON)
}
