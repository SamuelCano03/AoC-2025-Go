package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func ps(a ...any) {
	fmt.Println(a...)
}

func read() (string) {
	line, err := reader.ReadString('\n')
	if err != nil{
		return "error"
	}
	return strings.TrimSpace(line)
}

func pre() {
	ans := 0;
	for{
		line := read();
		if line=="error"{break;}
		if strings.Contains(line,"x"){
			parts:=strings.Split(line,":");
			dimPart:=parts[0];
			countsPart:=strings.Fields(parts[1]);
			dims:=strings.Split(dimPart,"x");
			w,_:=strconv.Atoi(dims[0]);
			h,_:=strconv.Atoi(dims[1]);
			sm:=0;
			tot:=w*h;
			for _,s:=range countsPart{
				val,_:=strconv.Atoi(s);
				sm+=val;
			}
			ps(w,h,tot,sm*9);
			if tot>=sm*9{ans+=1;}
		}
	}
	ps(ans);
}
func part1(){
	pre();
}
func part2(){

}
func main(){
	isPart1:=true;
	if isPart1{part1();
	}    else {part2();}
}

