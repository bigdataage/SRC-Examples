package main



import (
    "os"
    //"regexp"
    //"flag"
    "bufio"
    "bytes"
    "io"  
    "strings" 
    //"strconv"
    //"io/ioutil"
)





func main() {   
    a := -3/2
    b := -3.0/2.0
    println(a, b)
    lines  := readLines("txt")
    lines1 := readLines1("txt")
    num := len(lines)
    num1 := len(lines1)
    println()
    println(num, num1)
    println()
    
}




func readLines(path string)  (lines []string) {  //读取文件内容，返回一个字符串数组, 一行作为数组的一个元素.
    file, err := os.Open(path);  if err != nil {panic(err) };  defer file.Close()
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        part, prefix, err1 := reader.ReadLine();  if err1 != nil {break }       
        buffer.Write(part) 
        if !prefix { 
            if (!strings.EqualFold(buffer.String(), "")) {lines = append(lines, buffer.String()) }
            buffer.Reset()
        }
    }
    if err == io.EOF {err = nil }
    if err != nil {panic(err) }
    return
}




func readLines1(path string)  (lines []string) {  //读取文件内容，返回一个字符串数组, 一行作为数组的一个元素.
    file, err := os.Open(path);  if err != nil {panic(err) };  defer file.Close()
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        part, prefix, err1 := reader.ReadLine();  if err1 != nil {break }       
        buffer.Write(part) 
        if !prefix { 
            lines = append(lines, buffer.String()) 
            buffer.Reset()
        }
    }
    if err == io.EOF {err = nil }
    if err != nil {panic(err) }
    return
}
