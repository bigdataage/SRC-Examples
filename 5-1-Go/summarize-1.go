// Go语言的语法总结。

//运行：

//go run 1.go

//或

// 1.    go build 1.go -o 1
// 2.    ./1





package main    //当前包的名字。
import "fmt"    //导入的包的名字。Package implementing formatted I/O.
import "regexp"





func main() {  //main.main是每一个可执行程序的入口。main包中的main函数。
    fmt.Println("Hello, world-1.")     //输出默认含换行符。
    fmt.Println("Hello, world-2.\n")
    fmt.Printf("Hello, world-3.")      //输出默认不含换行符。 
    fmt.Printf("Hello, world-4.\n") 
    fmt.Println("######################## 1 ######################################")  


	var v1t int = 10
	var v2t = 20
	v3t := 30
    fmt.Println("%v %v %v", v1t, v2t, v3t)
    fmt.Printf("%v %v %v\n", v1t, v2t, v3t)
    fmt.Println("######################## 2 ######################################") 


	var v1f float64 = 10.98
	var v2f = 20.98
	v3f := 30.98
    fmt.Println("%v %v %v", v1f, v2f, v3f)
    fmt.Printf("%v %v %v\n", v1f, v2f, v3f)
    fmt.Printf("v1f v2f v3f\n")
    fmt.Println("######################## 3 ######################################")


	v1 := "1aaa"
	v2 := "2bbb"
	v3 := "3ccc"
	v4 := "4ddd"
	v5 := "5eee"
	v6 := "6fff"
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v\n", v1, v2, v3, v4, v5, v6)
	v1, v2, v3, v4, v5, v6 = v6, v5, v4, v3, v2, v1
    fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v\n", v1, v2, v3, v4, v5, v6)
    fmt.Println("######################## 4 ######################################")


    var v1s string = "aaa"
	var v2s = "bbb"
	v3s := "ccc"
	v1s, v2s, v3s = v3s, v2s, v1s
    fmt.Printf("v1s=%v, v2s=%v, v3s=%v\n", v1s, v2s, v3s)
    fmt.Println("######################## 5 ######################################")



    //正则表达式。
    str1 := "456abcd123"    //字符串
    reg1, _ := regexp.Compile("123")  //构造模式，自变量一个，函数值2个。

    //byte1 := [2]byte{0, 1}
    boo1:= reg1.MatchString(str1)
    fmt.Println(boo1)
    var digitsRegexp = regexp.MustCompile(`(\d+)\D+(\d+)`)


    someString := "1000abcd123"
    fmt.Println(digitsRegexp.FindStringSubmatch(someString))


    reg, _  := regexp.Compile("[a-z]\\d")
    str2 := "abcd bb12 ww 3 w4 56 kk789"
    num  := reg.FindAllString(str2, 1)
    fmt.Println(num)
}




