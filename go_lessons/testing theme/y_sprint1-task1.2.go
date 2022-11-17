package testing_theme

// User в системе.
type User struct {
	FirstName string
	LastName  string
}

// FullName возвращает фамилию и имя человека.
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}
