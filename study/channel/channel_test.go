package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	intChan := make(chan int)

	go func() {
		fmt.Println("begin run goroutine...")
		time.Sleep(3 * time.Second)
		intChan <- 1
		fmt.Println("finish run goroutine.")
	}()

	value, ok := <-intChan
	if !ok {
		fmt.Errorf("get intChan err!")
	}
	fmt.Println("value:", value)
}

func TestCacheChannel(t *testing.T) {
	intChan := make(chan int, 3)

	go func() {
		fmt.Println("begin run goroutine...")
		for i := 0; i < 20; i++ {
			fmt.Println("set value:", i)
			intChan <- i
		}
		fmt.Println("finish run goroutine.")
	}()

	for i := 0; i < 20; i++ {
		value := <-intChan
		fmt.Println("get value:", value)
	}
}

type User struct {
	Name  string
	Age   int
	Class Class
}

type Class struct {
	Name string
}

// use channel translate struct is value translate
func TestTranslateStructChannel(t *testing.T) {
	userChan := make(chan User, 1)

	go func() {
		user := User{Name: "xianbin", Age: 18, Class: Class{Name: "c1"}}
		fmt.Printf("src user: %+v\n", user)
		userChan <- user

		user.Age = 20
		fmt.Printf("src user: %+v\n", user)
	}()

	newUser := <-userChan
	fmt.Printf("src user: %+v\n", newUser)
}

func TestCloseChannel(t *testing.T) {
	intChan := make(chan int, 5)
	signCh := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			intChan <- i
			fmt.Println("set intCh value:", i)
			time.Sleep(time.Second)
		}

		close(intChan)
		fmt.Println("the intCh channel is closed")

		signCh <- 0
	}()

	go func() {
		for {
			i, ok := <-intChan
			fmt.Printf("get int channel %d, %v \n", i, ok)
			if !ok {
				fmt.Println("monitor intCh channel is closed")
				break
			}

			// 防止耗尽资源
			time.Sleep(time.Second * 4)
		}

		signCh <- 1
	}()

	<-signCh
	<-signCh
	fmt.Println("end")
}

func TestMergeChannel(t *testing.T) {
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)

	go func(in1, in2 <-chan int, out chan<- int) {
		for {
			select {
			case v := <-in1:
				out <- v
			case v := <-in2:
				out <- v
			}
		}
	}(input1, input2, output)

	go func() {
		for i := 0; i < 10; i++ {
			input1 <- i
			fmt.Println("set value:", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 20; i < 30; i++ {
			input2 <- i
			fmt.Println("set value:", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			select {
			case value := <-output:
				fmt.Println("get value:", value)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("main process exist.")
}

func TestQuitChannel(t *testing.T) {
	intChan := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-intChan:
				fmt.Println("get value:", v)
			case <-quit:
				fmt.Println("goroutine exit.")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		intChan <- i
	}
	quit <- true
	time.Sleep(time.Second)
	fmt.Println("main process exist.")
}

func TestProductCustomerChannel(t *testing.T) {
	intChan := make(chan int)
	prodChan := make(chan bool)
	custChan := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			intChan <- i
			fmt.Println("Product set value:", i)
			time.Sleep(time.Second)
		}
		prodChan <- true
	}()

	go func() {
		for {
			select {
			case v := <-intChan:
				fmt.Println("Customer get value:", v)
			case <-prodChan:
				custChan <- false
				return
			}
		}
	}()

	<-custChan
	fmt.Println("main process exist.")
}

func TestSequnseChannel(t *testing.T) {
	intChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		v := <-intChan
		fmt.Println(v)
	}()

	intChan <- 1
	fmt.Println(2)
}

func TestTimeoutChannel(t *testing.T) {
	intChan := make(chan int)
	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case v := <-intChan:
				fmt.Println("get value:", v)
			case <-time.After(time.Second * 3):
				quitChan <- true
				fmt.Println("goroutine timeout exit")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		intChan <- i
	}

	<-quitChan
	fmt.Println("goroutine timeout, main process exit")
}

func TestInAndOutChannel(t *testing.T) {
	intChan := make(chan int)
	quit := make(chan bool)

	go func(inChan chan<- int) {
		for i := 0; i < 10; i++ {
			inChan <- i
			fmt.Println("set value:", i)
			time.Sleep(time.Microsecond * 500)
		}
		quit <- true
	}(intChan)

	go func(outChan <-chan int) {
		for {
			select {
			case v := <-outChan:
				fmt.Println("get value:", v)
			case <-quit:
				fmt.Println("get inChan quit")
				return
			}
		}
	}(intChan)

	<-quit
	fmt.Println("goroutine timeout, main process exit")
}

func TestMaxChannel(t *testing.T) {
	maxNum := 3
	limitChan := make(chan bool, maxNum)
	quit := make(chan bool)

	for i := 0; i < 100; i++ {
		fmt.Println("start worker:")
		limitChan <- true
		go func(i int) {
			fmt.Println("do worker start: ", i)
			time.Sleep(time.Millisecond * 20)
			fmt.Println("do worker finish: ", i)

			<-limitChan

			if i == 99 {
				fmt.Println("goroutine quit")
				quit <- true
			}

		}(i)
	}

	<-quit
	fmt.Println("main process exit")
}

//func TestSignalChannel(t *testing.T) {
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//
//	go func() {
//		number := 0
//		for {
//			number++
//			fmt.Println("number is", number)
//			time.Sleep(time.Second * 2)
//		}
//	}()
//
//	fmt.Println("Press Ctrl + C exit process.")
//	<-quit
//	fmt.Println("main process exit")
//}

var lockChan = make(chan int, 1)
var gameCurrency = 20

func TestSynchronizeChannel(t *testing.T) {
	quit := make(chan bool, 2)

	go func() {
		for i := 0; i < 10; i++ {
			currency := rand.Intn(12) + 1
			go testSynchronizeUse(currency)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			currency := (rand.Intn(12) + 1)
			go testSynchronize_gain(currency)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	<-quit
	<-quit
	fmt.Println("main process exit")
}

func testSynchronizeUse(currency int) {
	lockChan <- 0

	if gameCurrency >= currency {
		srcGameCurrency := gameCurrency
		gameCurrency -= currency
		fmt.Printf("old %d, use %d，left %d\n", srcGameCurrency, currency, gameCurrency)
	} else {
		fmt.Printf("want use %d, just only %d\n", currency, gameCurrency)
	}

	<-lockChan
}

func testSynchronize_gain(currency int) {
	lockChan <- 0

	srcGameCurrency := gameCurrency
	gameCurrency += currency
	fmt.Printf("old %d, get %d，left %d\n", srcGameCurrency, currency, gameCurrency)

	<-lockChan
}
