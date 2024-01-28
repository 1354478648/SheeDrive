package utility

import "github.com/gogf/gf/crypto/gmd5"

// 密码加密（不加盐）
func EncryptPassword(password string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password))
}
