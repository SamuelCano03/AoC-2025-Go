
package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ps(a ...any) {
	fmt.Println(a...)
}

func read() (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func part1() {
	var mat[]string;
	for{
		s,er := read();
		if er!=nil{break}
		mat = append(mat, s);
	}
	var queue,ax []int;
	for i,e := range mat[0]{
		if e=='S'{ queue = append(queue, i); break;}
	}    
	n,ans:= len(mat),0;
	for i := range n-1{
		ax =  ax[:0];
		for _,e := range queue{
			if mat[i+1][e]=='^'{
				ax = append(ax, e-1);
				ax = append(ax, e+1);
				ans+=1;
			}else{
				ax = append(ax, e);
			}
		}
		var pre int = -1;
		var bx[]int;
		for _,e :=range ax{
			if e==pre{continue}
			bx = append(bx, e);
			pre = e;
		}
		queue = bx;
	}
	ps(ans);
}

func part2(){
	var mat[]string;
	for{
		s,er := read();
		if er!=nil{break}
		mat = append(mat, s);
	}
	var queue,ax []int;
	for i,e := range mat[0]{
		if e=='S'{ queue = append(queue, i); break;}
	}    
	n,m,ans:= len(mat),len(mat[0]),0;
	var tam [][]int;
	tam = make([][]int,n);
	for i := range tam{tam[i]=make([]int,m)}
	tam[0][queue[0]] += 1;
	for i := range n-1{
		ax =  ax[:0];
		for _,e := range queue{
			if mat[i+1][e]=='^'{
				tam[i+1][e-1] += tam[i][e];
				tam[i+1][e+1] += tam[i][e];
				ax = append(ax, e-1);
				ax = append(ax, e+1);
			}else{
				tam[i+1][e] += tam[i][e];
				ax = append(ax, e);
			}
		}
		var pre int = -1;
		var bx[]int;
		for _,e :=range ax{
			if e==pre{continue}
			bx = append(bx, e);
			pre = e;
		}
		queue = bx;
	}
	for i := range tam[n-1] { ans += tam[n-2][i]; }
	ps(ans);
}

func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}
