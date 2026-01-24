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
func atoi(b string)int64{
	e,_:=strconv.ParseInt(b,10,64);
	return e;
}
func pre() [][]int64 {
	line := read();
	parts:=strings.Split(line,",");
	mat := make([][]int64,0);
	for _,e := range parts{
		part := strings.Split(e,"-");
		a,b:=atoi(part[0]),atoi(part[1]);
		mat = append(mat, []int64{a,b});
	}
	return mat;
}
var ans int64 =0;
func part1(){
	mat:=pre();
	for _,e:=range mat{
		a,b:=e[0],e[1];
		mp := make(map[int64]bool);
		for num := range 1000000{
			for i:=2;i<=2;i++{
				if i*len(fmt.Sprint(num))>len(fmt.Sprint(b)){break;}
				nm:=atoi(strings.Repeat(fmt.Sprint(num),i));
				if a<=nm && nm<=b{mp[nm]=true;}
			}
		}
		for num := range mp{
			ans += num;
		}
	}
	ps(ans);
}
func part2(){
	mat:=pre();
	for _,e:=range mat{
		a,b:=e[0],e[1];
		mp := make(map[int64]bool);
		for num := range 1000000{
			for i:=2;i<=10;i++{
				if i*len(fmt.Sprint(num))>len(fmt.Sprint(b)){break;}
				nm:=atoi(strings.Repeat(fmt.Sprint(num),i));
				if a<=nm && nm<=b{mp[nm]=true;}
			}
		}
		for num := range mp{
			ans += num;
		}
	}
	ps(ans);
}
func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}

