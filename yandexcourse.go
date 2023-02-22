package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

type errorslice struct {
	errs []error
}

func (err errorslice) Error() string {
	sb := strings.Builder{}
	for _, e := range err.errs {
		sb.WriteString(e.Error())
		sb.WriteString("; ")
	}
	return strings.TrimRight(sb.String(), "; ")
}

var isNumeric func(string) bool = regexp.MustCompile(`\d`).MatchString

var (
	ErrFoundNumbers = errors.New("found numbers")
	ErrLineTooLong  = errors.New("line is too long")
	ErrNoTwoSpaces  = errors.New("no two spaces")
)

func MyCheck(input string) error {
	err := errorslice{}
	if isNumeric(input) {
		err.errs = append(err.errs, ErrFoundNumbers)
	}
	if len(input) >= 20 {
		err.errs = append(err.errs, ErrLineTooLong)
	}
	if strings.Count(input, " ") != 2 {
		err.errs = append(err.errs, ErrNoTwoSpaces)
	}
	if len(err.errs) > 0 {
		return err
	}
	return nil
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
