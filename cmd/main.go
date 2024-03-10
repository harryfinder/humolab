package main

//
//import (
//	"bufio"
//	"fmt"
//	"humo/homework1"
//	"humo/homework2"
//	"humo/homework3"
//	"humo/homework4"
//	"humo/homework5"
//	"humo/homework6"
//	"humo/homework7"
//	"log"
//	"math/rand"
//	"os"
//	"strconv"
//	"strings"
//	"time"
//)
//
//func main() {
//
//	log.Printf("Программа запущена: %s\n", time.Now().Format("15:04 02/01/2006"))
//
//	fmt.Println("Здравствуйте! Пожалуйста, выберите день обучения:")
//	fmt.Println("- 1 День. Знакомство")
//	fmt.Println("- 2 День. Первая программа - вычисление периметр квадрата, прямоугольник и тд. также поиск юбилея")
//	fmt.Println("- 3 День. Функции и пакеты: Создание и использование функций, организация кода в пакеты.")
//	fmt.Println("- 4 День. Объединение все задачи в единную программу,1 также бонусные задачи")
//	fmt.Println("- 5 День. Дурная школа :)")
//	fmt.Println("- 6 День. Дурная больница(использрвание указатели(pointer)")
//	fmt.Println("- 7 День. Продвинутая работа с коллекциями: Манипуляции с массивами и слайсами.")
//
//	var day, task int
//	fmt.Scanln(&day)
//
//	switch day {
//	case 1:
//		fmt.Println("Выберите задание из 1 дня:")
//		fmt.Println("1 - Знакомство")
//		fmt.Println("...")
//		fmt.Scanln(&task)
//		switch task {
//		case 1:
//			homework1.Introduction()
//		default:
//			fmt.Println("таких заданий нет, точнее больше нету :)")
//
//		}
//	case 2:
//		fmt.Println("Выберите задание из 2 дня:")
//		fmt.Println("1 - Периметр квадрата")
//		fmt.Println("2 - Периметр прямоугольника")
//		fmt.Println("3 - площадь прямоугольника")
//		fmt.Println("4 - Бонусная - вычисление Юбилея")
//
//		fmt.Scanln(&task)
//		switch task {
//		case 1:
//			homework2.PerimeterSquare()
//		case 2:
//			homework2.PerimeterRectangle()
//		case 3:
//			homework2.AreaRectangle()
//		case 4:
//			homework2.YearsToJubilee()
//		default:
//			fmt.Println("Такого задания нет.")
//		}
//	case 3:
//		homework3.OnlyPrint()
//	case 4:
//		homework4.OnlyPrint()
//		fmt.Println("Бонус - 1. Вывести таблицу умножения")
//		fmt.Println("Бонус - 2. Вычислить число Фибоначчи")
//		fmt.Scanln(&task)
//		switch task {
//		case 1:
//			homework4.PrintMultiplicationTable()
//		case 2:
//			fmt.Println("Введите порядковый номер числа Фибоначчи:")
//			var n int
//			fmt.Scanln(&n)
//			fmt.Printf("Число Фибоначчи №%d: %d\n", n, homework4.Fibonacci(n))
//		default:
//			fmt.Println("Выбрано несуществующее задание.")
//		}
//	case 5:
//		fmt.Println("Привет из дурной школы")
//		homework5.ConductExam()
//		fmt.Println("Пока, не учитесь под систему")
//	case 6:
//		fmt.Println("Привет из дурной больницы!")
//		patient := &homework6.Patient{
//			Name: "Гордон Фриманович Ломов", Age: rand.Intn(126) + 5, Organs: homework6.Organs{Leg: 2, Hands: 2, Livers: 2, Heart: false},
//		}
//		homework6.TakeOrder(patient)
//
//		homework6.BeginOperation(patient)
//		fmt.Println("Спасибо за визит в дурную больницу!")
//	case 7:
//		reader := bufio.NewReader(os.Stdin)
//
//		fmt.Println("Добро пожаловать в 7-й день!")
//		fmt.Println("Выберите задание:")
//		fmt.Println("- 1. Отзеркалить массив строк")
//		fmt.Println("- 2. Разделить массив чисел на чётные и нечётные")
//		fmt.Println("- 3. Бонусная задача")
//		fmt.Scanln(&task)
//		switch task {
//
//		case 1:
//			fmt.Println("Введите массив строк через пробел (например, 'Hello how are you ?'):")
//			input, _ := reader.ReadString('\n')
//			stringsSlice := strings.Fields(strings.TrimSpace(input))
//			homework7.ReverseSliceStrings(stringsSlice)
//		case 2:
//			fmt.Println("Введите массив чисел через пробел (например, '1 2 3 4 5 6 7 8 9'):")
//			input, _ := reader.ReadString('\n')
//			numStrs := strings.Fields(strings.TrimSpace(input))
//			var numbers []int
//			for _, numStr := range numStrs {
//				num, err := strconv.Atoi(numStr)
//				if err != nil {
//					fmt.Println("Ошибка при конвертации строки в число:", err)
//					return
//				}
//				numbers = append(numbers, num)
//			}
//			homework7.SplitSliceIntoEvenOdd(numbers)
//		case 3:
//			fmt.Println("Бонусное задание: 'Дурная больница' для нескольких клиентов.")
//			patients := []homework6.Patient{
//				{Name: "Никодим", Age: 25},
//				{Name: "Илюха", Age: 30},
//				{Name: "Сашка", Age: 35},
//				{Name: "Пациент 4", Age: 40},
//			}
//			homework6.ProcessPatients(patients)
//			fmt.Println("теперь у дурной больницы много клиентов :) Пока ! ")
//
//		default:
//			fmt.Println("Такого задания нет.")
//		}
//
//	default:
//		fmt.Println("Такого дня обучения нет или задания на этот день ещё не добавлены.")
//	}
//
//	log.Printf("Программа завершена: %s\n", time.Now().Format("15:04 02/01/2006"))
//}
