package repository

type User struct {
	Id           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash Hashed `db:"password_hash"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	BirthDate    string `db:"birth_date"`
	Email        string `db:"email"`
	PhoneNumber  string `db:"phone_number"`
}
