package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ps(a ...any) {
	fmt.Println(a...)
}

func read() (string) {
	line, err := reader.ReadString('\n')
	if err != nil{
		return ""
	}
	return strings.TrimSpace(line)
}
func atoi(s string)int{
	a,_:=strconv.Atoi(s);
	return a;
}
type pr struct{
	d rune;
	a int;
}
func pre() []pr{
	mat := make([]pr,0);
	for {
		line := read();
		if line==""{break;}
		mat = append(mat, pr{rune(line[0]),atoi(line[1:])})
	}
	return mat;
}
func part1(){
	mat := pre();
	n := 50
	cnt :=0;
	for _,e:=range mat{
		if e.d=='L'{
			n -= e.a;
		}else{
			n += e.a;
		}
		n = ((n%100)+100)%100;
		if n==0{cnt++;}
	}
	ps(cnt);
}
func part2(){
	mat := pre();
	n := 50
	cnt :=0;
	for _,e:=range mat{
		cnt += e.a/100;
		res:=e.a%100;
		x:=n;
		if e.d=='L'{
			n -= res;
		}else{
			n += res;
		}
		if n<=0 || n>99{if x!=0{cnt++;}}
		n = ((n%100)+100)%100;
		// ps(string(e.d),e.a,x,n,cnt);
	}
	ps(cnt);
}
func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}

