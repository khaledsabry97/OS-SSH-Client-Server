package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := getToArray("in.txt")
	arr = sort(arr)
	output(arr)
}

func readFile() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter File Name : ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text) -1]
	fmt.Println(text)
	return text
}


func getToArray(fileName string) [] int {
	_, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error!")
		//os.Exit(1)
		return nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr [] int
	var txt []string
	var val int
	for scanner.Scan() {
		txt = strings.Fields(scanner.Text())
		val, _ = strconv.Atoi(txt[0])
		arr = append(arr,val )
	}
	return arr
}
func sort(arr [] int) [] int{
	size := len(arr)
	for i := 0 ;i< size ;i ++ {
		for j := i ; j < size;j++{
			if(arr[i] > arr[j]){
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func output(arr [] int){
	f, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	size := len(arr)

	str := ""
	for i:= 0 ;i < size ;i ++ {
		str += strconv.Itoa(arr[i])
		if( i != size -1) {
			str += "\n"
		}
	}

	_,_ = f.WriteString(str)

}