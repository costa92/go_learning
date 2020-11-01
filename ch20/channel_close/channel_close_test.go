package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产数据
func dataProducer(ch chan int ,wg *sync.WaitGroup)  {
	go func() {
		for i := 0;i<10;i++ {
			ch <- i  // 0 到 9 放到 channel
		}
		close(ch) // 关闭 channel 避免堵塞

		// 已经关闭上发送 数据 提示：panic: send on closed channel
		//ch <- 11
		wg.Done()
	}()
}

// 消耗数据
func dataReceiver(ch chan int ,wg *sync.WaitGroup) {
	go func() {

		for i :=0; i<10 ; i++  {
			data := <-ch  // 取数据
			fmt.Println(data)
		}
		wg.Done()
	}()
}

// 消耗数据V1 生产的时候添加有 close channel 无需计数
func dataReceiverV1(ch chan int ,wg *sync.WaitGroup) {
	go func() {

		for{ // 与 dataReceiver 不相同的 是没有计数 不需要知道生产有多少数据的
			if data,ok:= <-ch; ok { // 取数据
				fmt.Println(data)
			}else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T){
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiverV1(ch,&wg)
	// 两个receiver
	wg.Add(1)
	dataReceiverV1(ch,&wg)
	wg.Wait()
}

