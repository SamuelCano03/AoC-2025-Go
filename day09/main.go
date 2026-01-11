
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

func read() (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(line), nil
}

type pair struct {
	ff, ss int;
};

func abs(x int)(int){
	if x < 0{x = -x}
	return x+1;
}
func part1(){
	id := make([]pair, 0);
	for {
		line,_ := read();
		if line == ""{break;}
		pr := strings.Split(line,",");
		a,_ := strconv.Atoi(pr[0]);
		b,_ := strconv.Atoi(pr[1]);
		id = append(id, pair{a,b})
	}
	ans,n := 0,len(id);
	for i := range n-1{
		for j:= i+1; j<n;j++{
			ans = max(ans,abs(id[i].ff-id[j].ff) * abs(id[i].ss-id[j].ss));
		}
	}
	ps(ans);
}
func isValid(id []pair, i int, j int) bool {
	mnx, mxx := min(id[i].ff, id[j].ff), max(id[i].ff, id[j].ff)
	mny, mxy := min(id[i].ss, id[j].ss), max(id[i].ss, id[j].ss)

	n := len(id)
	for k := range n {
		p1 := id[k]
		p2 := id[(k+1)%n] 

		if intersects(p1, p2, mnx, mxx, mny, mxy) {
			return false
		}
	}

	if !isInside(id, mnx, mxx, mny, mxy) {
		return false
	}

	return true
}

func intersects(p1, p2 pair, minX, maxX, minY, maxY int) bool {
	if p1.ff == p2.ff {
		x := p1.ff
		if x > minX && x < maxX {
			y1, y2 := min(p1.ss, p2.ss), max(p1.ss, p2.ss)
			if max(y1, minY) < min(y2, maxY) {
				return true
			}
		}
	} else { 
		y := p1.ss
		if y > minY && y < maxY {
			x1, x2 := min(p1.ff, p2.ff), max(p1.ff, p2.ff)
			if max(x1, minX) < min(x2, maxX) {
				return true
			}
		}
	}
	return false
}

func isInside(p []pair, minX, maxX, minY, maxY int) bool {
	testX := float64(minX+maxX) / 2.0
	testY := float64(minY+maxY) / 2.0

	intersections := 0
	n := len(p)
	for k := range n{
		p1 := p[k]
		p2 := p[(k+1)%n]

		x1, y1 := float64(p1.ff), float64(p1.ss)
		x2, y2 := float64(p2.ff), float64(p2.ss)

		if (y1 > testY) != (y2 > testY) {
			intersectX := (x2-x1)*(testY-y1)/(y2-y1) + x1
			if intersectX > testX {
				intersections++
			}
		}
	}
	return intersections%2 != 0
}
func part2(){
	var id []pair;
	for{
		line, _ := read();
		if line ==""{break;}
		pr := strings.Split(line, ",")
		a,_ := strconv.Atoi(pr[0]);
		b,_ := strconv.Atoi(pr[1]);
		id = append(id, pair{a,b});
	}
	ans,n := 0,len(id);
	for i := range n-1{
		for j:= i+1;j<n;j++{
			if !isValid(id,i,j) {continue;}
			ans = max(ans, abs(id[i].ff-id[j].ff)*abs(id[i].ss-id[j].ss));
		}
	}
	ps(ans);
}

func main(){
	isPart1:=false;
	if isPart1{
		part1();
	} else {part2();}
} 
