package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/benjaminboruff/base-go-app/internal/utils"
	"golang.org/x/crypto/argon2"
	"time"
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

// ****************************
// User methods DO NOT interact
// with the DB directly
// ****************************

func (u *User) GeneratePasswordHash(password string) error {
	params := &utils.Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	salt := make([]byte, params.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	u.Password = fmt.Sprintf(format, argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)
	return nil
}
