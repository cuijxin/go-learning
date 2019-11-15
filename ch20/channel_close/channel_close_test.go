package channelclose

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 关闭channel
		// ch <- 11 // 往已经关闭的channel上继续发消息，会产生一个panic
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 判断channel是否已经关闭
		for {
			if data, ok := <-ch; !ok {
				break
			} else {
				fmt.Println(data)
			}
		}

		// 如果不判断channel是否已经关闭，超出的消息是一条channel里类型的零值
		// for i := 0; i < 11; i++ {
		// 	data := <-ch
		// 	fmt.Println(data)
		// }

		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
