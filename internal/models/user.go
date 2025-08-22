package models

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

type User struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
	CreatedAt  time.Time
}

type UserModel struct {
	DB *sql.DB
}

type Argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// User methods
func (u *User) GeneratePasswordHash(password string) (string, error) {
	params := &Argon2Params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt := make([]byte, params.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	fullHash := fmt.Sprintf(format, argon2.Version, params.memory, params.iterations, params.parallelism, b64Salt, b64Hash)
	return fullHash, nil
}

func ComparePasswordAndHash(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	params := &Argon2Params{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.memory, &params.iterations, &params.parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	params.keyLength = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}
