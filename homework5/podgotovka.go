package homework5

import "math/rand"

func PrepareStudents() []Student {
	return []Student{
		{"Алексей", rand.Intn(2) == 1, rand.Intn(2) == 1, ""},
		{"Марина", rand.Intn(2) == 1, rand.Intn(2) == 1, ""},
		{"Игорь", rand.Intn(2) == 1, rand.Intn(2) == 1, ""},
	}
}
