package user

type User struct {
	Id    int    `db:"id"`
	Login string `db:"login"`
}
