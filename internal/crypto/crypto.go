package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	scryptN      = 1 << 15
	scryptR      = 8
	scryptP      = 1
	keyLen       = 32
	saltLen      = 16
	nonceSizeGCM = 12
)

var ErrInvalidCiphertext = errors.New("invalid ciphertext format")

func deriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, keyLen)
}

func Encrypt(plaintext []byte, password string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	key, err := deriveKey(password, salt)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceSizeGCM)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	final := make([]byte, 0, len(salt)+len(nonce)+len(ciphertext))
	final = append(final, salt...)
	final = append(final, nonce...)
	final = append(final, ciphertext...)

	encoded := base64.StdEncoding.EncodeToString(final)
	return encoded, nil
}

func Decrypt(encoded string, password string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	if len(data) < saltLen+nonceSizeGCM {
		return nil, ErrInvalidCiphertext
	}

	salt := data[:saltLen]
	nonce := data[saltLen : saltLen+nonceSizeGCM]
	ciphertext := data[saltLen+nonceSizeGCM:]

	key, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed")
	}

	return plaintext, nil
}