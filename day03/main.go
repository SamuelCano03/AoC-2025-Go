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
// func atoi(a string)int{
// 	e,_:=strconv.Atoi(a);
// 	return e;
// }
func pre() [][]int {
	mat := make([][]int,0);
	for{
		vec := make([]int,0);
		line:=read();
		if line==""{break;}
		for _,a := range line{
			e := int(a-'0');
			vec = append(vec, e);
		}
		mat = append(mat, vec);
	}
	return mat;
}
func part1(){
	mat := pre();
	n := len(mat[0]);
	ans := 0;
	for _,e := range mat{
		b,a,j:=0,0,0;
		for i:=n-2;i>=0;i--{
			if a<=e[i]{a=e[i];j=i;}
		}
		for i:=j+1;i<n;i++{b=max(b,e[i]);}
		ans += a*10+b;
	}
	ps(ans);
}
func part2() {
	mat := pre()
	ans := 0;

	for _, fila := range mat {
		n := len(fila)
		drops := n - 12;
		stack := make([]int, 0, n)

		for _, digito := range fila {
			for len(stack) > 0 && stack[len(stack)-1] < digito && drops > 0 {
				stack = stack[:len(stack)-1]
				drops--
			}
			stack = append(stack, digito)
		}
		stack = stack[:12]
		var sb strings.Builder
		for _, val := range stack {
			sb.WriteByte(byte('0' + val))
		}
		val, _ := strconv.Atoi(sb.String())
		ans += val
	}
	ps(ans)
}
func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}

