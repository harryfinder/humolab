package homework8_9

type User struct {
	Name    string
	Age     int
	Country string
}

var UsersMap = make(map[string]User)
