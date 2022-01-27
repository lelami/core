package core

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name"`
	Phone int64  `json:"phone" binding:"required"`
}

type AuthUser struct {
	Id    int   `json:"-" db:"id"`
	Phone int64 `json:"phone" binding:"required"`
	Code  int   `json:"code" binding:"required"`
}
