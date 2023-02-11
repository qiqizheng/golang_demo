package main

import "fmt"

//输出hello world
// func main() {
// 	fmt.Println("hello world")
// }

// func main() {
// 	//字符串拼接
// 	fmt.Println("GO" + "lang")
// 	//运算运算
// 	fmt.Println("1+1=", 1+1)
// 	fmt.Println("7.0/3.0 =", 7.0/3.0)
// 	//逻辑运算
// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
// }

//变量的定义及输出
// func main() {

// 	var a = "initial"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)

// 	var d = true
// 	fmt.Println(d)

// 	var e int
// 	fmt.Println(e)

// 	f := "apple"
// 	fmt.Println(f)
// }

// const s string = "constant"

// func main() {
// 	fmt.Println(s)

// 	const n = 500000

// 	const d = 3e20 / n
// 	fmt.Println(d)

// 	fmt.Println(int64(d))

// 	fmt.Println(math.Sin(n))
// }

// for 循环
// func main() {
// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}

// 		fmt.Println(n)
// 	}
// }

func main() {

	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println(" 7 is old")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")

	} else {
		fmt.Println(num, "has multiple digits")
	}

}
