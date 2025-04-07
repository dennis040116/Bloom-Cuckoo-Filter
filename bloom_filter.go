package filter

import (
	"github.com/demdxx/gocast"
	"github.com/dennis040116/Bloom-Cuckoo-Filter/encryptor"
)

type BloomFilter interface {
	Exists(data []byte) (bool, error)
	Add(data []byte) error
}

type localBloomFilter struct {
	m, n, k   int32
	bitmap    []int
	encryptor encryptor.Encryptor
}

func NewLocalBloomFilter(m, k int32, encryptor encryptor.Encryptor) BloomFilter {
	bitmap := make([]int, m/32+1) //每个int元素有32个bit位,m/32+1是为了防止构造时有除不尽的问题
	return &localBloomFilter{
		m:         m,
		k:         k,
		bitmap:    bitmap,
		encryptor: encryptor,
	}
}

func (bf *localBloomFilter) Exists(data []byte) (bool, error) {

	khash, err := bf.getKEncrypted(data)
	if err != nil {
		return false, err
	}
	for _, offset := range khash {
		index := offset & int32(len(bf.bitmap)-1)
		bitOffset := offset & 31
		if bf.bitmap[index]&(1<<bitOffset) == 0 {
			return false, nil
		}
	}

	return true, nil
}

func (bf *localBloomFilter) Add(data []byte) error {
	bf.n++
	khash, err := bf.getKEncrypted(data)

	if err != nil {
		return err
	}

	for _, offset := range khash {
		index := offset & int32(len(bf.bitmap)-1)
		bitOffset := offset & 31
		bf.bitmap[index] |= (1 << bitOffset)
	}

	return nil
}

func (bf *localBloomFilter) getKEncrypted(data []byte) ([]int32, error) {
	encrypteds := make([]int32, 0, bf.k)
	origin := data
	for i := 0; int32(i) < bf.k; i++ {
		encrypted, err := bf.encryptor.Encrypt(origin)
		if err != nil {
			return []int32{}, nil
		}
		encrypteds = append(encrypteds, encrypted)
		hash_bytes := []byte(gocast.ToString(encrypted))
		hash_bytes = append(hash_bytes, byte(i))
		origin = hash_bytes
	}
	return encrypteds, nil
}
