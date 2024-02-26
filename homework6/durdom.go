package homework6

import (
	"fmt"
	"log"
	"math/rand"
)

func TakeOrder(patient *Patient) {
	log.Println("Пациент зашёл за чеком")
	if patient.Age <= 15 {
		fmt.Println("Слишком мал для операции!")
		return
	} else if patient.Age >= 150 {
		fmt.Println("Слишком...стар?")
		return
	}
	patient.Order = rand.Intn(20) + 1
	log.Printf("Пациент забрал чек. Номер: %d\n", patient.Order)
}

func BeginOperation(patient *Patient) {
	log.Println("Пациент пошёл на операцию.")
	peoples := rand.Intn(20) + patient.Order
	for i := 1; i <= peoples; i++ {
		if i == patient.Order {
			log.Println("Дождался!", i)
			patient.MedBook.Annotation = true
			patient.MedBook.Diagnose = "Установка нового сердца"
			Operation(patient)
			return
		} else {
			log.Println("Очередной человек прошёл мимо: ", i)
		}
	}
}

func Operation(patient *Patient) {
	// Операция началась
	luck := rand.Intn(4)
	switch luck {
	case 0:
		fmt.Println("Ой... а где нога?")
		patient.Organs.Leg--
	case 1:
		fmt.Println("Не повезло! Рука пропала!")
		patient.Organs.Hands--
	case 2:
		fmt.Println("А у него всегда была одна почка?")
		patient.Organs.Livers--
	case 3:
		fmt.Println("Первый человек с новым сердцем!")
		patient.Organs.Heart = true
		patient.MedBook.Annotation = false
		patient.MedBook.Diagnose = "Живой!"
	}

	if !patient.Organs.Heart {
		log.Printf("Пациент вернулся! Но что-то со мной не так: %+v\n", patient)
	} else {
		log.Printf("Пациент вернулся! И чувствует себя отлично: %+v\n", patient)
	}
}
