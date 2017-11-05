package backend

import (
	"crypto"
	_ "crypto/sha512" // use sha512
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bluemir/zumo/datatype"
)

func (b *backend) CreateToken(username, unhashedKey string) (*datatype.Token, error) {
	// TODO Check user exist....
	if len(unhashedKey) < 8 {
		return nil, fmt.Errorf("Token too short")
	}

	token := &datatype.Token{
		Username:  username,
		HashedKey: hash(unhashedKey),
	}

	token, err := b.store.PutToken(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (b *backend) Token(tokenStr string) (*datatype.Token, error) {
	if tokenStr[:6] != "Basic " {
		return nil, fmt.Errorf("Not collect token type")
	}

	buf, err := base64.StdEncoding.DecodeString(tokenStr[6:])
	if err != nil {
		return nil, fmt.Errorf("token decode fail")
	}

	arr := strings.SplitN(string(buf), ":", 2)
	if len(arr) < 2 {
		return nil, fmt.Errorf("Not enought argument")
	}

	return b.store.GetToken(arr[0], hash(arr[1]))
}
func hash(str string) string {
	hashed := crypto.SHA512.New().Sum([]byte(str))
	return hex.EncodeToString(hashed)
}
