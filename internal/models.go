package internal

type User struct {
	Id         int               `json:"id"`
	Login      string            `json:"login"`
	Email      string            `json:"email"`
	Characters map[int]Character `json:"-"`
}

type Character struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Race   string `json:"race"`
}
