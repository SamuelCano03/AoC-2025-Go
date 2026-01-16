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

func pre() ([][]byte,[]int,[]int) {
	mp := make([][]byte,0);
	for{
		line:=read();
		if line ==""{break}
		mp = append(mp,[]byte(line));
	}
	x:=[]int{-1,0,1,0,-1,-1,1,1};
	y:=[]int{0,1,0,-1,-1,1,-1,1};
	return mp,x,y;
}
func part1(){
	mp,x,y:=pre();
	n,m:=len(mp),len(mp[0]);
	ans:=0;
	for i := range n{
		for j:= range m{
			if mp[i][j] != '@'{continue}
			cnt := 0;
			for k := range 8{
				ni,nj:=i+x[k], j+y[k];
				if ni<0 || ni>=n || nj<0 || nj>=m{continue}
				if mp[ni][nj] == '@'{cnt++;}
			}
			if cnt<4{ans++;}
		}
	}
	ps(ans);
}
func part2(){
	mp,x,y:=pre();
	n,m:=len(mp),len(mp[0]);
	ans,flg:=1,-1;
	for {
		if flg==0{break;}
		ans += flg;
		flg =0;
		for i := range n{
			for j:= range m{
				if mp[i][j] != '@'{continue}
				cnt := 0;
				for k := range 8{
					ni,nj:=i+x[k], j+y[k];
					if ni<0 || ni>=n || nj<0 || nj>=m{continue}
					if mp[ni][nj] == '@'{cnt++;}
				}
				if cnt<4{flg++;mp[i][j]='.'}
			}
		}
	}
	ps(ans);
}
func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}

