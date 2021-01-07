package main

import "fmt"

func rev(str string) string {
	if len(str) <= 0 {
		return str
	}
	temp := str[1:]
	firstchar := string([]rune(str)[0])
	//return reverseString(myStr.substring(1)) + myStr.charAt(0);
	//fmt.Println(temp)
	defer fmt.Println("Function called ", len(temp)+1, "time")
	return rev(temp) + firstchar
}

func main() {
	var str string
	fmt.Println("Enter String to Reverse")
	fmt.Scan(&str)
	str2 := rev(str)
	fmt.Println(str2)
}
