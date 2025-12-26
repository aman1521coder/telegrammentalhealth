package main
/* 
import (
	"fmt"
	"net/http"
	"sync"
	"time"
)


for [eopedplew]

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (n *TreeNode) Insert(k int) {
	if k < n.value {
		if n.left == nil {
			n.left = &TreeNode{value: k}

		} else {
			n.left.Insert(k)

		}

	} else if k > n.value {
		if n.right == nil {
			n.right = &TreeNode{value: k}
		} else {
			n.right.Insert(k)
		}
	}

}
func (n *TreeNode) Serach(k int) bool {
	if n.value == k {
		return true
	}
	if k < n.value {
		if n.left == nil {
			return false
		} else {
			return n.left.Serach(k)
		}
	}
	if k > n.value {
		if n.right == nil {
			return false
		} else {
			return n.right.Serach(k)
		}
	}
	return false
}
func Tiker() {
	x := time.NewTicker(11 * time.Millisecond)
	defer x.Stop()
	counter := 0
	for {
		counter++
		fmt.Println("tick : ", counter)
		<-x.C
	}
}

func SendingtoChannel(x chan int) {
	for i := 0; i < 5; i++ {
		x<-i
	}
}
func RecivingFromChannels(x chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-x)
	}
}
func Checksite(url string, wg *sync.WaitGroup) {
	start:=time.Now()
	defer wg.Done()
	resp,err:=http.Get(url)
	fmt.Println("Checking site")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
fmt.Println(time.Since(start))
}
func ChecksitewithChannel(url string, done chan bool) {
	start:=time.Now()
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
		return
	}
	done<-true
	fmt.Println(resp.Status)
	fmt.Println(time.Since(start))
}
func main() {


	   	for _, v := range []int{1, 2, 3} {

	   	   go func(x int ) {
	   	       fmt.Println(v)

	   	   }(v)

	   }
	
	ch:=make(chan bool)
	wg:=sync.WaitGroup{}
		wg.Add(1)
	go Checksite("https://www.google.com",&wg)
	wg.Wait()
	go ChecksitewithChannel("https://www.google.com",ch)
fmt.Println(<-ch)

}
 */
 