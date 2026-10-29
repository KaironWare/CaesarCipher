package main

import (
	"fmt"
	"strings"
)

var alp = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ_абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
var enalp = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz")
var shift int

func encrypt(clean string) string {
	var en []string

	for _, t := range clean {
		found := false

		for i, v := range alp {
			if t == v {
				index := (i + shift) % len(alp)
				en = append(en, string(alp[index]))
				found = true
				break
			}
		}

		if found {
			continue
		}

		for i, v := range enalp {
			if t == v {
				index := (i + shift) % len(enalp)
				en = append(en, string(enalp[index]))
				found = true
				break
			}
		}

		if !found {
			en = append(en, string(t))
		}
	}

	return strings.Join(en, "")
}

func decrypt(clean string) string {
	var de []string

	for _, t := range clean {
		found := false

		for i, v := range alp {
			if t == v {
				index := (i - shift + len(alp)) % len(alp)
				de = append(de, string(alp[index]))
				found = true
				break
			}
		}

		if found {
			continue
		}

		for i, v := range enalp {
			if t == v {
				index := (i - shift + len(enalp)) % len(enalp)
				de = append(de, string(enalp[index]))
				found = true
				break
			}
		}

		if !found {
			de = append(de, string(t))
		}
	}

	return strings.Join(de, "")
}

func main() {
	var com int
	var str string
	fmt.Println("1. Зашифровать / encrypt")
	fmt.Println("2. Расшифровать / decrypt")
	fmt.Scan(&com)
	fmt.Println("Текст / text :")
	fmt.Scan(&str)
	fmt.Println("Сдвиг / shift :")
	fmt.Scan(&shift)
	if com == 1 {
		fmt.Println("Результат / result :")
		fmt.Println(encrypt(str))
	} else if com == 2 {
		fmt.Println("Результат / result :")
		fmt.Println(decrypt(str))
	} else {
		fmt.Println("error")
	}
}
