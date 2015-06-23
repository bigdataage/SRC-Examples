//找出转录起始区域的序列。
//输入为两个文件：染色体序列文件，基因注释文件。


//example:
//
//time go run seq6.go   -file filename   -folder foldername  -up 2000 -down 2000  -output  result.txt
//
// or:
//
// go build seq6.go
// time ./seq6  -file filename   -folder foldername   -up 2000  -down 2000  -output  result.txt





package main




import (
    "os"
    "regexp"
    "flag"
    "bufio"
    "bytes"
    "io"   
    "strconv"
    "runtime"
    "io/ioutil"
    "sync"
)









var (
    file_g, folder_g, output_g  string  //全局变量均以“_g”结尾。
    up_g, down_g  int
    cores_g int = 16    //CPU数目                  
    wg_g = sync.WaitGroup{}              // 用用于等待所有 goroutine 结束。
    c_g = make(chan int, cores_g)
)


func commandLineArguments() {  
    flag.StringVar(&file_g,     "file",     "hunmanrefFlat_noNR_norepet.txt",   "gene annotations")
    flag.StringVar(&folder_g,   "folder",   "human_genome",                     "the DNA sequence of all chomosomes") 
    flag.IntVar   (&up_g,       "up",       2000,                               "the up basepairs of txStart")
    flag.IntVar   (&down_g,     "down",     2000,                               "the down basepairs of txStart")
    flag.StringVar(&output_g,   "output",   "result.txt",                       "the result file") 
    flag.Parse()
    println()
    println("The Command Line Arguments: \n")
    println("file_g   is:  ", file_g)
    println("folder_g is:  ", folder_g)
    println("up_g     is:  ", up_g)
    println("down_g   is:  ", down_g)
    println("output_g is:  ", output_g)
}









func main() { 
    runtime.GOMAXPROCS(cores_g)                                   
    commandLineArguments()  
    lines1 := readLines(file_g)   //一行一个元素。 
    //num := len(lines1)-1   
    num := 200
    fh1, _ := os.Create(output_g);  defer fh1.Close()
    
    for i:=1; i<=num; i++ {   //一次循环一行。
        go prosOneLine(lines1[i], i, fh1) 
    }
   

}











func  prosOneLine(str1 string, i int, fh1 *os.File) () { 
    c_g <- 1    
    reg1, _ := regexp.Compile("\\S+")  
    array1 := reg1.FindAllString(str1, 5)
    txStart, _ := strconv.Atoi(array1[4])       
    byte1, err6 := ioutil.ReadFile(folder_g+"/"+array1[2]+array1[3]); if err6 != nil {panic(err6) }  
    start2 := txStart-up_g;   end2 := txStart+down_g 
    region1 := string(byte1[start2:end2]) 
    _, err4 := io.WriteString(fh1, ">"+str1+"\n"+region1+"\n");  if err4 != nil {panic(err4)}    
    println(i)               
    <- c_g
}








func readLines(path2 string)  (lines2 []string) {  
    file, err := os.Open(path2);  if err != nil {panic(err) };  defer file.Close()
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        part, prefix, err1 := reader.ReadLine();  if err1 != nil {break }       
        buffer.Write(part) 
        if !prefix { 
            lines2 = append(lines2, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {err = nil }
    if err != nil {panic(err) }
    return
}







 
