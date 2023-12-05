package tasks

import "fmt"

/*
   Данный фрагмент кода может привести к ошибке "slice bounds out of range", так как переменная justString пытается
   сохранить только первые 100 символов из строки v, длина которой 1 << 10 (1024 символа). Если строка v содержит менее
   100 символов, возникнет паника из-за выхода за пределы среза.
   Чтобы исправить эту проблему, нужно сначала проверить длину строки перед выделением среза. Вот исправленный код:
*/

var (
	justString string
)

func createHugeString(length int) string {
	hugeString := ""
	for i := 0; i < length; i++ {
		hugeString += "a"
	}
	return hugeString
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		justString = v
	} else {
		justString = v[:100]
	}
}

func Fifteenth() {
	someFunc()
	fmt.Println(justString)
}
