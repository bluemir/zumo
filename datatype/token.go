package datatype

// Token is
type Token struct {
	Username  string
	HashedKey string
	Salt      string
	Lables    map[string]string
}
