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

func read() (string) {
	line, err := reader.ReadString('\n')
	if err != nil{
		return ""
	}
	return strings.TrimSpace(line)
}
var memo map[string]int;
func dfs(u string, v string, mp map[string][]string) int{
	if u==v{return 1;}
	if val,ok:=memo[u]; ok{return val;}
	tot := 0;
	for _,e:= range mp[u]{
		tot += dfs(e,v,mp);
	}
	memo[u] = tot;
	return tot;
}
func pre() map[string][]string{
	mp := make(map[string][]string);
	for{
		line := read();
		if line==""{break;}
		vals := strings.Split(line, " ");
		a := vals[0][:len(vals[0])-1];
		for i := 1; i< len(vals); i++{
			mp[a] = append(mp[a], vals[i]);
		}
	}
	return mp;
}
func part1(){
	mp := pre();
	memo = make(map[string]int);
	ans := dfs("you","out",mp);
	ps(ans);
}

func part2(){
	mp := pre();
	a,b,c,d:="svr","fft","dac","out";
	memo = make(map[string]int)
	path1 := dfs(a, b, mp)
	
	memo = make(map[string]int)
	path2 := dfs(b, c, mp)
	
	memo = make(map[string]int)
	path3 := dfs(c, d, mp)

	memo = make(map[string]int)
	path4 := dfs(a, c, mp)
	
	memo = make(map[string]int)
	path5 := dfs(c, b, mp)
	
	memo = make(map[string]int)
	path6 := dfs(b, d, mp)

	ans := (path1 * path2 * path3) + (path4 * path5 * path6);
	ps(ans);
}
func main(){
	isPart1:=true;
	if isPart1{part1();
	}    else {part2();}
}

