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

func part1() {
	ans := 0;
	for{
		line := read();
		if line == ""{break;}
		parts:=strings.Split(line," ");
		indicator := parts[0][1:len(parts[0])-1];
		vin := 0;
		for i := range len(indicator){
			if indicator[i]=='#'{
				vin = vin ^  (1<<i);
			}
		}
		n := len(parts);
		buttons := parts[1:n-1];
		var buts []int;
		for _,e:=range buttons{
			ax := e[1:len(e)-1];
			bx := strings.Split(ax,",");
			var cx []int;
			for _,ee:= range bx{
				val,_ := strconv.Atoi(ee); 
				cx = append(cx, val);
			}
			val:=0;
			for _,ee:= range cx{
				val = val ^ (1<<ee);
			}
			buts = append(buts, val);
		}
		num,nas := len(buts),1000;
		for i := range (1<<num){
			vax,cnt := 0,0;
			for j:= range num{
				if (i & (1<<j))>=1{
					vax = vax ^ buts[j];
					cnt++;
				}
			}
			// ps(vin,vax);
			if vin == vax{nas = min(nas,cnt);}
		}
		ans += nas;
	}
	ps(ans);
}

func key(target []int) string {
	return fmt.Sprint(target)
}

func minPresses(buttons [][]int, target []int, memo map[string]int) int {
	allZero := true
	for _, v := range target {
		if v != 0 {	allZero = false;break}
	}
	if allZero {return 0}
	k := key(target);
	if val,ok:=memo[k];ok{return val;}
	n,m := len(buttons),len(target);
	limit := 1<<n;
	best := -1;
	for mask := range limit{
		remainder := make([]int, m);
		copy(remainder,target);
		costPhase1,poss :=0,true;
		for b := range n{
			if (mask & (1<<b)) != 0{
				costPhase1++;
				for i:=range m {
					remainder[i] -= buttons[b][i];
				}
			}
		}
		for i := range m{
			if remainder[i]<0 || remainder[i]%2!=0{poss=false; break;}
		}
		if poss{
			nextTarget := make([]int, m);
			for i:=range m{	nextTarget[i]=remainder[i]/2; }
			res := minPresses(buttons,nextTarget,memo);
			if res!=-1{
				totalCost := costPhase1 + 2*res;
				if best == -1 || totalCost < best{best= totalCost;}
			}
		}
	}
	memo[k]=best;
	return best;
}
func part2() {
	ans := 0
	for {
		line := read()
		if line == "" {break}
		parts := strings.Fields(line)
		n := len(parts)

		joltageRaw := parts[n-1]
		joltageContent := joltageRaw[1 : len(joltageRaw)-1]
		jox := strings.Split(joltageContent, ",")
		var target []int
		for _, e := range jox {
			val, _ := strconv.Atoi(e)
			target = append(target, val)
		}
		buttonsRaw := parts[1 : n-1]
		var buttons [][]int
		
		m := len(target)
		
		for _, e := range buttonsRaw {
			inner := e[1 : len(e)-1]
			bx := strings.Split(inner, ",")
			btnEffect := make([]int, m)
			
			for _, ee := range bx {
				counterIdx, _ := strconv.Atoi(ee)
				if counterIdx < m {
					btnEffect[counterIdx] = 1
				}
			}
			buttons = append(buttons, btnEffect)
		}
		memo := make(map[string]int)
		ans += minPresses(buttons, target, memo)
	}
	ps(ans)
}
func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}

