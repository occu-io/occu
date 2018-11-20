package models

type User struct {
	UID      int64  `json:"uid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"admin" form:"-"`
}
