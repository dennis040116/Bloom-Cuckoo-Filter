package filter

import (
	"testing"

	"github.com/dennis040116/Bloom-Cuckoo-Filter/encryptor"
)

func TestLocalBloomFilter(t *testing.T) {
	lbf := NewLocalBloomFilter(100, 6, encryptor.NewEncryptor())
	str := "dennis"
	lbf.Add([]byte(str))
	if ok, _ := lbf.Exists([]byte(str)); !ok {
		t.Errorf("%s: %v", str, ok)
	}
	t.Errorf("true")

	notexist := "dennis1"
	if ok, _ := lbf.Exists([]byte(notexist)); !ok {
		t.Errorf("%s : not exist", notexist)
	}
}
