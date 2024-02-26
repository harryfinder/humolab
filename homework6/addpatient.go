package homework6

import "fmt"

func ProcessPatients(patients []Patient) {
	for i, patient := range patients {
		fmt.Printf("Обработка пациента #%d: %s\n", i+1, patient.Name)
		TakeOrder(&patient)
		BeginOperation(&patient)
	}
}
