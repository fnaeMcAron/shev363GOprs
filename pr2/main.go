//практическая 2
//Шевченко Андрей
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	TypeSingle = "single"
	TypeDouble = "Double"
	TypeSuite  = "Suite"

	StatusFree        = "Free"
	StatusBooked      = "Booked"
	StatusMaintenance = "Maintenance"
)

type Order struct {
	ID           int
	Items        []int
	Total        float64
	Adress       string
	isCompleated bool
}

type LogEntry struct {
	IP_Adress string
	HTTP_Code int
	TimePoint time.Time
}

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type HotelRoom struct {
	typeRoom   string
	statusRoom string
	priceRoom  float64
}

type Counts struct {
	CountLetters  int
	CountWords    int
	CountSentence int
}

func NewOrder(order *map[int]Order) {
	var lastIntKey int
	var Items []int
	var Total float64
	var Adress string
	for k := range *order {
		if k > lastIntKey {
			lastIntKey = k + 1
		}
	}
	for {
		fmt.Print("Введите ID предметов:  (когда закончите введите -1)")
		var Item int
		fmt.Scanln(&Item)
		if Item == -1 {
			break
		} else if Item <= 0 {
			fmt.Print("Такого ID нет")
		} else {
			Items = append(Items, Item)
		}
	}
	for {
		fmt.Print("Введите Цену:")
		fmt.Scanln(&Total)
		if Total > 0 {
			break
		}
	}
	for {
		fmt.Print("Введите адрес:")
		fmt.Scanln(&Adress)
		if Adress != "" {
			break
		}
	}

	orderValue := Order{
		ID:           lastIntKey,
		Items:        Items,
		Total:        Total,
		Adress:       Adress,
		isCompleated: false,
	}

	(*order)[lastIntKey] = orderValue

	fmt.Println("Заказ добавлен")
}

func validateUser(name string, age int, email string) error {
	if name != "" && len(name) > 50 {
		return errors.New("Длина имени не должна быть больше 50 и он не может быть пустым")
	}
	if age < 18 && age > 120 {
		return errors.New("Возраст должен быть между 18 и 120")
	}
	IndexCheck := strings.Index(email, "@")
	if IndexCheck == -1 {

		return errors.New("email должен содержать знак @")
	}
	return nil
}

func CountVotes(Votes []string) {
	var VoteBoris int
	var VoteAnna int
	var VoteVictor int
	var AllVotes float64
	var j string
	for _, j = range Votes {
		if j == "Борис" {
			VoteBoris++
		}
		if j == "Анна" {
			VoteAnna++
		}
		if j == "Виктор" {
			VoteVictor++
		}
	}
	fmt.Printf("Проголосовавших за Анну: %d\n", VoteAnna)
	fmt.Printf("Проголосовавших за Виктора: %d\n", VoteVictor)
	fmt.Printf("Проголосовавших за Бориса: %d\n", VoteBoris)
	AllVotes = float64(VoteBoris+VoteAnna+VoteVictor) / 100
	fmt.Printf("Процент проголосовавших за Анну: %f\n", AllVotes*float64(VoteAnna))
	fmt.Printf("Процент проголосовавших за Виктора: %f\n", AllVotes*float64(VoteVictor))
	fmt.Printf("Процент проголосовавших за Бориса: %f\n", AllVotes*float64(VoteBoris))
}

func CheckSalary(Employees []Employee) (float64, float64) {
	var sumSalary float64
	for _, value := range Employees {
		sumSalary += value.Salary
	}
	AVGsumSalary := sumSalary / float64(len(Employees))
	return sumSalary, AVGsumSalary
}

func SortLogEntries(LogEntries []LogEntry) []LogEntry {
	var SortLogEntries []LogEntry
	for _, value := range LogEntries {
		if value.HTTP_Code >= 400 && value.HTTP_Code <= 599 {
			SortLogEntries = append(SortLogEntries, value)
		}
	}
	return SortLogEntries
}

func BookedRoom(reservRoom *map[string]HotelRoom) {
	var TakeRoom string
	fmt.Println("(Выберите комнату) Свободные комнаты:")
	for i, value := range *reservRoom {
		if value.statusRoom == StatusFree {
			fmt.Println(i)
		}
	}
	for {
		fmt.Scanln(&TakeRoom)
		if TakeRoom != "" {
			break
		}
	}
	for i := range *reservRoom {
		if i == TakeRoom {
			//(*reservRoom)[i].statusRoom = StatusBooked
		}
	}
	fmt.Println("Комната: " + TakeRoom + " была забронирована")

}

func UniqueTagsSeeker(Posts [][]string) map[string]bool {
	Tags := make(map[string]bool)

	for _, post := range Posts {
		for _, tag := range post {
			Tags[tag] = true
		}
	}
	return Tags
}

func textStats(text string) Counts {
	var mLetters = len(text)
	var mWords = len(strings.Fields(text))
	dots := strings.Count(text, ".")
	voskl := strings.Count(text, "!")
	questions := strings.Count(text, "?")
	mSentence := dots + voskl + questions

	return Counts{
		CountLetters:  mLetters,
		CountWords:    mWords,
		CountSentence: mSentence,
	}
}

func main() {
	var Exit bool = false
	var TaskManage int
	for {
		fmt.Println("")
		if Exit {
			break
		}
		fmt.Println("Выберите задание (введите значение от 1-16) (-1 чтобы выйти): ")
		fmt.Scanln(&TaskManage)
		switch TaskManage {
		case 1:
			PriceDays := map[string]int{
				"ПН": 2100,
				"ВТ": 2100,
				"СР": 2100,
				"ЧТ": 2100,
				"ПТ": 2850,
				"СБ": 2850,
				"ВС": 2850,
			}
			price := PriceDays["ВТ"] + PriceDays["СР"] + PriceDays["ЧТ"] + PriceDays["ПТ"] + PriceDays["СБ"] + PriceDays["ВС"] + PriceDays["ЧТ"] + PriceDays["ПТ"]
			fmt.Println(price)
		case 2:
			var WeightBagageMain float64
			var WeightBagageHand float64
			var WeightBagageHandDop float64
			for {
				fmt.Print("Введите вес основного багажа: ")
				fmt.Scanln(&WeightBagageMain)
				fmt.Print("Введите вес ручной клади: ")
				fmt.Scanln(&WeightBagageHand)
				fmt.Print("Введите вес доп. ручной клади: ")
				fmt.Scanln(&WeightBagageHandDop)
				if WeightBagageMain >= 0 && WeightBagageHand >= 0 && WeightBagageHandDop >= 0 {
					break
				}
			}
			fmt.Printf("Общий вес багажа: %f", WeightBagageMain+WeightBagageHand+WeightBagageHandDop)
			fmt.Println()
		case 3:
			Orders := map[int]Order{
				1: {
					ID:           1,
					Items:        []int{2, 3, 4},
					Total:        200.5,
					Adress:       "ул. Ленина 25",
					isCompleated: false,
				},
			}
			NewOrder(&Orders)
			fmt.Println(Orders)
		case 4:
			var CondidatsVotes = []string{"Анна", "Борис", "Виктор", "Борис", "Виктор", "Борис", "Виктор", "Анна", "Анна", "Анна"}
			CountVotes(CondidatsVotes)
		case 5:
			err := validateUser("Грыгорий", 20, "email@email.com")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Ошибок в регистрации не было обнаружено")
			}
		case 6:
			var PostsAndTags = [][]string{
				{"go", "backend"},
				{"git", "go", "tools"},
			}
			var UniqueTags = UniqueTagsSeeker(PostsAndTags)
			fmt.Println(UniqueTags)
		case 7:
			Employees := []Employee{
				{
					ID:       1,
					Name:     "Grigoriy",
					Position: "Слесарь",
					Salary:   35000,
				},
				{
					ID:       2,
					Name:     "Alex",
					Position: "Программист",
					Salary:   25000,
				},
				{
					ID:       3,
					Name:     "Alena",
					Position: "Бухгалтер",
					Salary:   15000,
				},
			}
			SumSalary, AvgSumSalary := CheckSalary(Employees)
			fmt.Printf("Общая получаемая сумма рабочими: %f Средняя получаемая сумма: %f \n", SumSalary, AvgSumSalary)
		case 8:
			LogEntries := []LogEntry{
				{
					IP_Adress: "198.5.3.1",
					HTTP_Code: 405,
					TimePoint: time.Now(),
				},
				{
					IP_Adress: "195.5.1.1",
					HTTP_Code: 129,
					TimePoint: time.Now(),
				},
				{
					IP_Adress: "178.5.3.1",
					HTTP_Code: 506,
					TimePoint: time.Now(),
				},
			}
			fmt.Println("Логи до этого:")
			for _, value := range LogEntries {
				fmt.Println(value)
			}
			var SortLogEntries []LogEntry = SortLogEntries(LogEntries)
			fmt.Println("Отсортированные логи:")
			for _, value := range SortLogEntries {
				fmt.Println(value)
			}
		case 9:
			reservRoom := map[string]HotelRoom{
				"109": {
					typeRoom:   TypeSingle,
					statusRoom: StatusFree,
					priceRoom:  305,
				},
				"111": {
					typeRoom:   TypeDouble,
					statusRoom: StatusBooked,
					priceRoom:  400,
				},
				"101": {
					typeRoom:   TypeSuite,
					statusRoom: StatusMaintenance,
					priceRoom:  500,
				},
			}
			BookedRoom(&reservRoom)
		case 10:
			//уже час ночи, кто нибудь, допишите вывод за меня, я уже не понимаю
			//текст брал с https://fish-text.ru/
			fmt.Println(textStats("Повседневная практика показывает, что понимание сути ресурсосберегающих технологий не даёт нам иного выбора, кроме определения стандартных подходов. Как принято считать, представители современных социальных резервов набирают популярность среди определенных слоев населения, а значит, должны быть смешаны с не уникальными данными до степени совершенной неузнаваемости, из-за чего возрастает их статус бесполезности."))
		case 11:

		case 12:

		case 13:

		case 14:

		case 15:

		case 16:

		case -1:
			Exit = true
		default:
			fmt.Println("Данного значения нет")
		}
	}
}
