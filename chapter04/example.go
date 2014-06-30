package main

import (
	"fmt"
	"time"
)

// 一个简单的计时器，每秒打印一次当前时间
func tick() {
	c := time.Tick(1 * time.Second)

	for now := range c {
		fmt.Printf("%v \n", now)
	}
}

func f1(c chan int) {
	ticker := time.NewTicker(1 * time.Second).C
	timer := time.NewTimer(4 * time.Second).C

	for true {
		select {
		case m := <-c:
			fmt.Println(m)
		case <-ticker:
			fmt.Println("-------")
		case <-timer:
			fmt.Println("+++++++")
		}
	}
}

func f2(c chan int) {
	for true {
		time.Sleep(3 * time.Second)
		c <- 1
	}
}

func f3() {
	for true {

		time.Sleep(2 * time.Second)
		fmt.Println("f3")
	}
}

func main() {
	t0 := time.Now()
	fmt.Println(t0)
	for i := 0; i < 3; i = i + 1 {
		time.Sleep(time.Second * 1)
	}
	t1 := time.Now()
	fmt.Println(t1)
	dur := t1.Sub(t0)
	fmt.Println(dur.Nanoseconds())
	fmt.Println(dur.Seconds())

	dur, _ = time.ParseDuration("30h")
	fmt.Println(dur)

	c := make(chan int)
	go f1(c)
	go f2(c)

	time.Sleep(time.Minute * 1)
}
