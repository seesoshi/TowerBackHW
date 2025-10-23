package main

import (
	"fmt"
	"strings"
)

func toLitera(n int) string {
	/* перебором i от -2^32 до 12306 получено что макс число равно 2^32+1
	 */
	var nstr = ""
	ones := [10]string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	tens := [10]string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	fromtentotwenty := [10]string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	hundreds := [10]string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятсот"}
	onesthousands := [10]string{"", "одна тысяча", "две тысячи", "три тысячи", "четыре тысячи", "пять тысяч", "шесть тысяч", "семь тысяч", "восемь тысяч", "девять тысяч"}
	largenumbers := [5]string{"миллион", "миллиард", "триллион", "квадриллион", "квинтиллион"}
	//числа до тысячи
	if (n%100 > 9) && (n%100 < 20) {
		nstr = hundreds[(n/100)%10] + " " + fromtentotwenty[n%10]
	} else {
		nstr = hundreds[(n/100)%10] + " " + tens[(n/10)%10] + " " + ones[n%10]
	}
	//тысячи
	n /= 1000
	if n%1000 != 0 {
		if (n%100 > 9) && (n%100 < 20) {
			nstr = hundreds[(n/100)%10] + " " + fromtentotwenty[n%10] + " тысяч " + nstr
		} else if n%10 == 0 {
			nstr = hundreds[(n/100)%10] + " " + tens[(n/10)%10] + " тысяч " + nstr
		} else {
			nstr = hundreds[(n/100)%10] + " " + tens[(n/10)%10] + " " + onesthousands[n%10] + " " + nstr
		}
	}
	//для миллионов и тд
	var i = 0
	n /= 1000
	for n > 0 {
		//проверка на пустой порядок 000
		if n%1000 != 0 {
			if (n%100 > 9) && (n%100 < 20) {
				nstr = hundreds[(n/100)%10] + " " + fromtentotwenty[n%10] + " " + largenumbers[i] + "ов " + nstr
			} else if n%10 == 0 {
				nstr = hundreds[(n/100)%10] + " " + tens[(n/10)%10] + " " + ones[n%10] + " " + largenumbers[i] + "ов " + nstr
			} else {
				//окончание
				var endstr = ""
				if (n%10 > 1) && (n%10 < 5) {
					endstr = "а"
				} else {
					endstr = "ов"
				}
				nstr = hundreds[(n/100)%10] + " " + tens[(n/10)%10] + " " + ones[n%10] + " " + largenumbers[i] + endstr + " " + nstr
			}
		}
		n /= 1000
		i++
	}
	//удаление лишних пробелов
	nstr = strings.TrimSpace(nstr)
	for strings.Contains(nstr, "  ") {
		nstr = strings.Replace(nstr, "  ", " ", -1)
	}
	return nstr

}

func main() {
	var n int
	fmt.Printf("Введите целое (%T) число ", n)
	fmt.Scan(&n)
	fmt.Printf("Число n = %#v \n", n)
	if n >= 12307 {
		fmt.Println("Число не изменилось")
	} else {
		var flag = true
	outer:
		for n < 12307 {
			if n < 0 {
				n *= -1
			} else if n%7 == 0 {
				n *= 39
			} else if n%9 == 0 {
				n *= 13
				n++
			} else {
				n += 2
				n *= 3
			}
			if (n%13 == 0) && (n%9 == 0) {
				fmt.Print("service error")
				flag = false
				break outer
			} else {
				n++
			}
		}
		if flag {
			fmt.Printf("Новое число %#v", n)
			fmt.Printf(" %#v ", toLitera(n))
		}
	}
}
