package levels_final

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var t = TimeNow()

func currentDayOfTheWeek() string {
	weekday := t.Weekday()
	switch weekday {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	default:
		return "Воскресенье"
	}
}

func dayOrNight() string {
	if t.Hour() >= 10 && t.Hour() <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	weekday := t.Weekday()
	current := int(weekday)
	friday := 5
	if current <= friday {
		return friday - current
	}
	return 7 - (current - friday)
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	day := currentDayOfTheWeek()
	if strings.ToLower(day) == strings.ToLower(answer) {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	check := dayOrNight()
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	if strings.ToLower(check) == strings.ToLower(answer) {
		return true, nil
	}
	return false, nil
}
