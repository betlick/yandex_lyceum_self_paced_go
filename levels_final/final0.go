package levels_final

import (
	"fmt"
	"time"
)

func main0() {
	var t, s1, s2, s3 string
	fmt.Scanln(&t)
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)
	var a, b, c float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	sum := int(a*float64(100))%100 + int(b*float64(100))%100 + int(c*float64(100))%100
	sumKop := sum % 100
	sumRub := int(a) + int(b) + int(c) + (sum / 100)
	data, _ := time.Parse("02.01.2006", t)
	data = data.AddDate(0, 0, 15)
	fmt.Printf(
		fmt.Sprintf(
			"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n",
			s2,
			s1,
			s3,
		),
	)
	fmt.Printf(
		fmt.Sprintf(
			"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n",
			data.Format("02.01.2006"),
		),
	)
	fmt.Printf(
		fmt.Sprintf(
			"Общая сумма выплат составит %d руб. %d коп.\n\n",
			sumRub,
			sumKop,
		),
	)
	fmt.Println("С уважением,")
	fmt.Println("Гл. бух. Иванов А.Е.")
}
