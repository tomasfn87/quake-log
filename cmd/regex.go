package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println("matched:", matched, " error:", err)

	matched, err = regexp.Match(`\w+fo\w+`, []byte("seafood"))
	fmt.Println("matched:", matched, " error:", err)

	var re *regexp.Regexp
	re, err = regexp.Compile(`\wat`)
	fmt.Println("re:", re, " error:", err)
	str1 := `bat mat pot sat cat rat pat vat hat`
	fmt.Println("MatchString:", re.MatchString(str1))
	fmt.Println("Find:", string(re.Find([]byte(str1))))
	fmt.Printf("FindAll: %q\n", re.FindAll([]byte(str1), -1))
	fmt.Println("FindIndex:", re.FindIndex([]byte(str1)))
	fmt.Println("FindAllIndex:", re.FindAllIndex([]byte(str1), -1))
	fmt.Println("Match:", re.Match([]byte([]byte(str1))))
	fmt.Printf("ReplaceAllLiteral: %s\n", re.ReplaceAllLiteral([]byte(str1), []byte("dog")))
	fmt.Println("ReplaceAllLiteralString:", re.ReplaceAllLiteralString(str1, "cat"))
	fmt.Println("String:", re.String())

	fmt.Println()
	var re2 *regexp.Regexp
	re2, err = regexp.Compile(`^(img-\d+)\.png`)
	f := "img-13421234123.png"
	a := re2.ReplaceAllString(f, `$1`)
	fmt.Println(a)
}
