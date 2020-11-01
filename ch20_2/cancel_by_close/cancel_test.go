package cancel_by_close

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 多路机制
func isCancelled(ctx context.Context) bool{
	select {
	case <- ctx.Done():   // 接收取消
		return true
	default:
		return false
	}
}


func TestCancel(t *testing.T) {
	ctx ,cancel := context.WithCancel(context.Background()) // 父节点
	for i:=0;i<5 ;i++  {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i,"Cancelled")
		}(i,ctx)
	}
	//cancel_1(cancelChan) // 只是执行了 4
	cancel()  // channel 机制 的广播
	time.Sleep(time.Millisecond * 1)

}
