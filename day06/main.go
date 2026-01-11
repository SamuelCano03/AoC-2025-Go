
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Interval struct{
	L int 
	R int
}

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

func main() {
	var a [][]string;
	for {
		e,er:=read();
		if er!=nil{break}
		o := strings.Fields(e);
		a = append(a, o);
	}
	n,m:=len(a),len(a[0]);
	var ans int64;
	for i := range m{
		var pre int64 =0
		if a[n-1][i] == "*" {pre=1}
		for j := range n-1{
			num,_:=strconv.ParseInt(a[j][i],10,64)
			if a[n-1][i]=="*"{
				pre *= num
			}else{
				pre += num
			}
		}
		ans += pre
	}
	ps(ans)
}
