package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 1. 指针
	num := 10
	pointer(&num)
	fmt.Println("num:", num)

	slice := []int{1, 2, 3, 4, 5}
	slicePointer(&slice)
	fmt.Println("切片：", slice)

	//2. goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		printOddNum()
	}()
	go func() {
		defer wg.Done()
		printEvenNum()
	}()
	wg.Wait()
	fmt.Println("所有数字打印完成！")

	// 3. 任务调度器
	tasks := []Task{
		func() {
			time.Sleep(time.Second * 1)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(time.Second * 2)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(time.Second * 3)
			fmt.Println("任务3完成")
		},
	}
	results := runTask(tasks)
	for i, result := range results {
		fmt.Printf("任务%d完成，耗时：%v\n", i, result)
	}

	// 4. Shape
	rect := Rectangle{}
	circle := Circle{}
	rect.Area()
	rect.Perimeter()
	circle.Area()
	circle.Perimeter()

	// 5. Person 和 Employee 结构体测试
	person := Person{
		Name: "张三",
		Age:  30,
	}
	employee := Employee{
		Person:     person,
		EmployeeID: "EMP001",
	}
	employee.PrintInfo()

	// 6. 通道
	ch := make(chan int, 10)
	go generateNumbers(ch)
	go printNumbers(ch)
	time.Sleep(time.Second * 10)

	// 7. 缓冲通道
	testChannel2()

	// 8. 互斥锁
	testMutex1()

	// 9. 原子操作
	testMutex2()
}

// 1. 指针
func pointer(n *int) {
	*n += 10
	fmt.Println("指针：", n)
}

// 2. 切片指针
func slicePointer(s *[]int) {
	slice := *s
	for i := range slice {
		slice[i] *= 2
	}
}

// 3. goroutine
func printOddNum() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数：", i)
	}
}

func printEvenNum() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数：", i)
	}
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
type Task func()

func runTask(tasks []Task) map[int]time.Duration {
	results := make(map[int]time.Duration)
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	var mu sync.Mutex

	for i, t := range tasks {
		go func(i int, task Task) {
			defer wg.Done()

			startTime := time.Now()
			t()
			cost := time.Since(startTime)

			mu.Lock()
			results[i] = cost
			mu.Unlock()
		}(i, t)
	}
	wg.Wait()
	return results
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct{}

func (r Rectangle) Area() {
	fmt.Println("矩形面积：", 10*20)
}

func (r Rectangle) Perimeter() {
	fmt.Println("矩形周长：", 10+20)
}

type Circle struct{}

func (c Circle) Area() {
	fmt.Println("圆形面积：", 3.14*10*10)
}

func (c Circle) Perimeter() {
	fmt.Println("圆形周长：", 2*3.14*10)
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("员工信息：", e.Name, e.Age, e.EmployeeID)
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
func generateNumbers(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("写入通道数字：", i)
		ch <- i
	}
	close(ch)
}

func printNumbers(ch chan int) {
	for num := range ch {
		fmt.Println("读取通道数字：", num)
	}
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func testChannel2() {
	ch := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch)

		for i := range 100 {
			ch <- i
			fmt.Println("生产者发送：", i)
		}
		fmt.Println("生产者发送完成")
	}()

	go func() {
		defer wg.Done()

		count := 0
		for num := range ch {
			fmt.Println("消费者接收：", num)
			count++
		}
		fmt.Println("消费者接收完成，接收数量：", count)
	}()

	wg.Wait()
	fmt.Println("所有数字接收完成")
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func testMutex1() {
	counter := 0
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			for range 1000 {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("mutex计数器值：", counter)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func testMutex2() {
	counter := int64(0)
	var wg sync.WaitGroup
	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("atomic计数器值：", counter)
}