/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/25
 * Time: 16:42
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine1() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second)
}

func goroutine2() {
	Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// 等待发布新闻
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")
}

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

func Producer() <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "海老握り"
		ch <- "鮪とろ握り"
		ch <- "hello, 世界"
		close(ch)
	}()
	return ch
}

func goroutine3() {
	var ch = Producer()
	for s := range ch {
		fmt.Println("Consumed", s)
	}
}

func Publish2(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) //缺少会导致死锁
	}()
	return ch
}

func goroutine4() {
	wait := Publish2("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for the news...")
	<-wait
	fmt.Println("The news is out, time to leave.")
}

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0
		n++
		ch <- n
	}()
	n := <-ch
	n++
	fmt.Println(n)
}

func goroutine5() {
	sharingIsCaring()
}

type AtomicInt struct {
	mu sync.Mutex
	n  int
}

func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.n += n
	a.mu.Unlock()
}

func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func lockItUp() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1)
		close(wait)
	}()
	n.Add(1)
	<-wait
	fmt.Println(n.Value())
}

func goroutine6() {
	lockItUp()
}

func race() {
	var wg sync.WaitGroup
	wg.Add(5)
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Print(i)
	//		wg.Done()
	//	}()
	//} // 55555
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Print(n)
			wg.Done()
		}(i)
	} // 20413
	//for i := 0; i < 5; i++ {
	//	n := i
	//	go func() {
	//		fmt.Print(n)
	//		wg.Done()
	//	}()
	//} // 20314
	wg.Wait()
	fmt.Println()
}

func goroutine7() {
	race()
}

func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
	}
	wg.Done()
}

func goroutine8() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
	}
}

func main() {
	goroutine1()
	goroutine2()
	goroutine3()
	goroutine4()
	goroutine5()
	goroutine6()
	goroutine7()
	goroutine8()
}
