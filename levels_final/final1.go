package levels_final

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {
	s1, s2, s3, s4, s5 := "-", "-", "-", "-", "-"

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if line == "очередь" {
			fmt.Println("1. " + s1)
			fmt.Println("2. " + s2)
			fmt.Println("3. " + s3)
			fmt.Println("4. " + s4)
			fmt.Println("5. " + s5)
			continue
		}
		if line == "количество" {
			filled := 0
			if s1 != "-" {
				filled++
			}
			if s2 != "-" {
				filled++
			}
			if s3 != "-" {
				filled++
			}
			if s4 != "-" {
				filled++
			}
			if s5 != "-" {
				filled++
			}
			fmt.Println("Осталось свободных мест:", 5-filled)
			fmt.Println("Всего человек в очереди:", filled)
			continue
		}

		if line == "конец" {
			fmt.Println("1. " + s1)
			fmt.Println("2. " + s2)
			fmt.Println("3. " + s3)
			fmt.Println("4. " + s4)
			fmt.Println("5. " + s5)
			return
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Запись на место невозможна: некорректный ввод")
			continue
		}

		name := parts[0]
		num, err := strconv.Atoi(parts[1])
		if err != nil || num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", parts[1])
			continue
		}

		filled := 0
		if s1 != "-" {
			filled++
		}
		if s2 != "-" {
			filled++
		}
		if s3 != "-" {
			filled++
		}
		if s4 != "-" {
			filled++
		}
		if s5 != "-" {
			filled++
		}

		if filled == 5 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
			continue
		}

		isTaken := false
		switch num {
		case 1:
			if s1 != "-" {
				isTaken = true
			}
		case 2:
			if s2 != "-" {
				isTaken = true
			}
		case 3:
			if s3 != "-" {
				isTaken = true
			}
		case 4:
			if s4 != "-" {
				isTaken = true
			}
		case 5:
			if s5 != "-" {
				isTaken = true
			}
		}

		if isTaken {
			fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			continue
		}

		switch num {
		case 1:
			s1 = name
		case 2:
			s2 = name
		case 3:
			s3 = name
		case 4:
			s4 = name
		case 5:
			s5 = name
		}
	}
}
