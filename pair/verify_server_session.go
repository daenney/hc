package pair

import(
    "github.com/brutella/hap"
)

type VerifySession struct {
    otherPublicKey [32]byte
    publicKey [32]byte
    secretKey [32]byte
    sharedKey [32]byte
    encryptionKey [32]byte
}

func NewVerifySession() (*VerifySession) {
    return &VerifySession{}
}

// Generate Curve25519 public, secret and shared key for a specified other public key
// The other public key is also stored for further use in `otherPublicKey` property
func (s *VerifySession) GenerateKeysWithOtherPublicKey(otherPublicKey [32]byte) {
    secretKey := hap.Curve25519_GenerateSecretKey()
    publicKey := hap.Curve25519_PublicKey(secretKey)
    sharedKey := hap.Curve25519_SharedSecret(secretKey, otherPublicKey)
    
    s.otherPublicKey = otherPublicKey
    s.secretKey = secretKey
    s.publicKey = publicKey
    s.sharedKey = sharedKey
}

func (s *VerifySession) PublicKey() []byte {
    return s.publicKey[:]
}

func (s *VerifySession) EncryptionKey() []byte {
    return s.encryptionKey[:]
}