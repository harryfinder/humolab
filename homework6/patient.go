package homework6

type MedBook struct {
	Annotation bool
	Diagnose   string
}

type Organs struct {
	Leg    int
	Hands  int
	Livers int
	Heart  bool
}

type Patient struct {
	Name    string
	Age     int
	Order   int
	Organs  Organs
	MedBook MedBook
}
