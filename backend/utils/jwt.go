package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims กำหนดข้อมูลที่ต้องการเก็บใน JWT
type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
