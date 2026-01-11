
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type pnt struct {
	x,y,z int;
}
type pir struct {
	dis float64;
	a,b int;
}

func dis(a pnt, b pnt)(d float64){
	dx,dy,dz:=a.x-b.x,a.y-b.y,a.z-b.z;
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz));
}
func part1() {
	var mat [] pnt;
	for{
		var line string;
		line,er := read()
		if er!=nil{break}
		part := strings.Split(line,",");
		a,_ := strconv.Atoi(part[0]);
		b,_ := strconv.Atoi(part[1]);
		c,_ := strconv.Atoi(part[2]);
		mat = append(mat, pnt{a,b,c})
	}
	var lis []pir;
	n := len(mat);
	for i:=0;i<n-1;i+=1{
		for j:=i+1;j<n;j+=1{
			lis = append(lis, pir{dis(mat[i],mat[j]),i,j});
		}
	}
	sort.Slice(lis,func(i,j int) bool{
		return lis[i].dis < lis[j].dis;
	});
	adj := make([][]int, n);
	for i:=range 1000{
		a,b := lis[i].a,lis[i].b;
		adj[a] = append(adj[a], b);
		adj[b] = append(adj[b], a);
	}
	vis := make([]bool,n);
	var val[]int;
	for i:= range n{
		if vis[i]!= false {continue}
		vis[i]=true;
		var queue []int;
		cnt := 1;
		queue = append(queue, i);
		for len(queue)>0{
			top:=queue[len(queue)-1];
			queue = queue[:len(queue)-1];
			for _,e:=range adj[top]{
				if vis[e] {continue;} 
				vis[e]=true;
				queue = append(queue, e);
				cnt += 1;
			}
		}
		val = append(val, cnt);
	}
	sort.Ints(val);
	ans := 1
	for i:=len(val)-1;i>=max(0,len(val)-3);i-=1{
		ps(val[i]);
		ans *= val[i];
	}
	ps(ans);
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i]) 
	}
	return d.parent[i]
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
		return true 
	}
	return false 
}

func part2() {
	var mat []pnt
	for {
		line, er := read()
		if er != nil {break}
		part := strings.Split(line, ",")
		a, _ := strconv.Atoi(part[0])
		b, _ := strconv.Atoi(part[1])
		c, _ := strconv.Atoi(part[2])
		mat = append(mat, pnt{a, b, c})
	}
	var lis []pir
	n := len(mat)
	for i := 0; i < n-1; i += 1 {
		for j := i + 1; j < n; j += 1 {
			lis = append(lis, pir{dis(mat[i], mat[j]), i, j})
		}
	}
	sort.Slice(lis, func(i, j int) bool {
		return lis[i].dis < lis[j].dis
	})
	dsu := NewDSU(n)
	numComponents := n 

	for _, e := range lis {
		if dsu.Union(e.a, e.b) {
			numComponents--
			if numComponents == 1 {
				valA := mat[e.a].x
				valB := mat[e.b].x
				ans := valA * valB
				ps(ans)
				return
			}
		}
	}
}

func main(){
	isPart1:=false;
	if isPart1{part1();
	}    else {part2();}
}
