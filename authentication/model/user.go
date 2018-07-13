package model

type User struct {
	//UUID     string `json:"uuid"`
	//Username string `json:"username"`
	//Password string `json:"password"`

	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
