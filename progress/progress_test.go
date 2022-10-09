package progress

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestProgress(t *testing.T) {
	p := NewProgress(5)
	_ = p.Run("start")
	for i := 0; i < 5; i++ {
		_ = p.Incr(1)
		time.Sleep(time.Microsecond)
	}
	_ = p.Finish()
	t.Log("ok")
}

func TestProgressParallel(t *testing.T) {
	p := NewProgress(10000)
	c, cancel := context.WithCancel(context.Background())

	_ = p.Run("start")
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-c.Done():
					return
				case <-time.After(time.Millisecond * time.Duration(rand.Int63n(10)+1)):
					_ = p.Incr(1)
				}
			}
		}()
	}

	time.Sleep(time.Second * 2)
	p.SetTitle("change title")
	time.Sleep(time.Second)
	_ = p.SetMax(12000)
	p.UpdateInterval = time.Millisecond * 600
	time.Sleep(time.Second * 2)
	cancel()
	_ = p.Finish()
	t.Log("ok")
}
