package models

type JwtToken struct {
	Token string `json:"token" form:"token"`
}

type JwtClaims struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}
