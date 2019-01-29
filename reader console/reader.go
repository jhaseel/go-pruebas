package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func parseSlice(s string)[]string {
	return strings.Split(s,",")
}

func validarNUmero(num[]string){
	list :=make([]int,0)
	for i,v := range num {
		val,err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error en pos",i,"=>",err)
		} else {
			list=append(list,val)
		}
	}
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text,_:= reader.ReadString('\n')
	numeros:= parseSlice(text)
	
	fmt.Println(numeros)
	//validarNUmero(numeros)

}