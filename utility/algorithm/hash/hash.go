package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	compiler_time "github.com/s9rA16Bf4/go-evil/utility/variables/compiler-time"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

func Hashing_md5(value string) string {
	temp := md5.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_sha1(value string) string {
	temp := sha1.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_sha224(value string) string {
	temp := sha256.Sum224([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_sha256(value string) string {
	temp := sha256.Sum256([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}
func Hashing_sha384(value string) string {
	temp := sha512.Sum384([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_sha512(value string) string {
	temp := sha512.Sum512([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_sha3_224(value string) string {
	temp := sha3.Sum224([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}
func Hashing_sha3_256(value string) string {
	temp := sha3.Sum256([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}
func Hashing_sha3_384(value string) string {
	temp := sha3.Sum384([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}
func Hashing_sha3_512(value string) string {
	temp := sha3.Sum512([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp[:]))
	return hex.EncodeToString(temp[:])
}

func Hashing_blake2s_256(value string) string {
	temp, _ := blake2s.New256([]byte(randstr.String(32)))
	temp_hash := temp.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp_hash))
	return hex.EncodeToString(temp_hash)
}
func Hashing_blake2b_256(value string) string {
	temp, _ := blake2b.New256([]byte(randstr.String(32)))
	temp_hash := temp.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp_hash))
	return hex.EncodeToString(temp_hash)
}
func Hashing_blake2b_384(value string) string {
	temp, _ := blake2b.New384([]byte(randstr.String(32)))
	temp_hash := temp.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp_hash))
	return hex.EncodeToString(temp_hash)
}

func Hashing_blake2b_512(value string) string {
	temp, _ := blake2b.New512([]byte(randstr.String(32)))
	temp_hash := temp.Sum([]byte(value))
	compiler_time.Set_variable(hex.EncodeToString(temp_hash))
	return hex.EncodeToString(temp_hash)
}
