package cryptogram

import (
	"GRE3000/const_conf"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strings"
)

func mapSortFmt(params map[string]string, is_sort_by_key bool) string {
	swap_list := []string{}
	swap_map := map[string]string{}
	tmp_str_list := []string{}

	if !is_sort_by_key {
		for key, value := range params {
			swap_map[value] = key
			swap_list = append(swap_list, value)
		}
		sort.Strings(swap_list)
		for _, each_ := range swap_list {
			tmp_str_list = append(tmp_str_list, fmt.Sprintf("%s=%s", swap_map[each_], each_))
		}
	} else {
		for key, value := range params {
			swap_map[key] = value
			swap_list = append(swap_list, key)
		}
		sort.Strings(swap_list)
		for _, each_ := range swap_list {
			tmp_str_list = append(tmp_str_list, fmt.Sprintf("%s=%s", each_, swap_map[each_]))
		}
	}
	return strings.Join(tmp_str_list, "&")
}

func GetMD5Hash(text string) string {
	hash_er := md5.New()
	hash_er.Write([]byte(text))
	return hex.EncodeToString(hash_er.Sum(nil))
}

func KeyValueSign(params map[string]string, external_key string, is_sort_by_key bool) string {
	raw_str := mapSortFmt(params, is_sort_by_key) + "&key=" + external_key
	return strings.ToUpper(GetMD5Hash(raw_str))
}

func AESCBCEncrypt(aes_key, aes_iv, text string) string {
	plain_text := []byte(text)
	if len(text) == 0 {
		plain_text = append(plain_text, 0)
	}
	for i := 1; i <= len(plain_text)%aes.BlockSize; i++ {
		plain_text = append(plain_text, 0)
	}
	block, _ := aes.NewCipher([]byte(aes_key))
	cipher_text := make([]byte, len(plain_text))
	mode := cipher.NewCBCEncrypter(block, []byte(aes_iv))
	mode.CryptBlocks(cipher_text, plain_text)
	return fmt.Sprintf("%x", cipher_text)
}

func AESCBCDecrypt(aes_key, aes_iv, cipher_text string) string {
	cipher_complex_text, _ := hex.DecodeString(cipher_text)
	block, err := aes.NewCipher([]byte(aes_key))
	if err != nil {
		panic(err)
	}
	if len(cipher_complex_text) < aes.BlockSize {
		panic("Length of cipher_complex_text is too short")
	}
	if len(cipher_complex_text)%aes.BlockSize != 0 {
		panic("Length error")
	}
	mode := cipher.NewCBCDecrypter(block, []byte(aes_iv))
	mode.CryptBlocks(cipher_complex_text, cipher_complex_text)
	x := len(cipher_complex_text) - 1
	for ; x-1 >= 0 && cipher_complex_text[x-1] == 0; x-- {
	}
	return fmt.Sprintf("%s", cipher_complex_text[:x])
}

func GetSha1(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		io.WriteString(h, s)
	}
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func SafeCookieName(raw_cookie_name string) string {
	return GetMD5Hash(fmt.Sprintf(const_conf.PlatformCookieName, raw_cookie_name))
}
