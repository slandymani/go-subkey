package common

// KeyPair is the secret and public key
type KeyPair interface {
	Signer
	Verifier

	// Secret returns the secret in bytes.
	Secret() [32]byte
	// Public returns the secret in bytes.
	Public() [32]byte

	// SS58Address returns the Base58 string of the address.
	// network: must be one of the known networks
	SS58Address(network Network, ctype ChecksumType) (string, error)
}

type Signer interface {
	// Sign signs the message and returns the signature.
	Sign(msg []byte) ([64]byte, error)
}

type Verifier interface {
	// Verify verifies the signature.
	Verify(msg []byte, signature [64]byte) bool
}