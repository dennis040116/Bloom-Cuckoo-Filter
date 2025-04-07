package encryptor

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type Encryptor interface {
	Encrypt(data []byte) (int32, error)
}

type encryptor struct{}

func NewEncryptor() Encryptor {
	return &encryptor{}
}

func (e *encryptor) Encrypt(data []byte) (int32, error) {
	hasher := murmur3.New32()
	_, err := hasher.Write(data)
	if err != nil {
		return 0, err
	}
	return int32(hasher.Sum32() % math.MaxInt32), nil
}
