package progress

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type Progress struct {
	UpdateInterval time.Duration // Default value is 250ms
	max            int
	current        int
	title          string
	isFinished     bool
	isRunning      bool
	startedAt      time.Time
	ctx            context.Context
	cancel         context.CancelFunc
	*sync.Mutex
}

const CLEAR_LINE = "\r\x1b[2K"

func NewProgress(max int) *Progress {
	p := new(Progress)
	p.max = max
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.Mutex = new(sync.Mutex)

	return p
}

func (p *Progress) SetTitle(t string) {
	p.Lock()
	defer p.Unlock()
	p.title = t

	// Print title immediately
	p.print()
}

func (p *Progress) SetMax(m int) error {
	p.Lock()
	defer p.Unlock()

	if m < p.current {
		return fmt.Errorf("Max is less than current value: %v", p.current)
	}

	p.max = m

	// Print new progress immediately
	p.print()

	return nil
}

func (p *Progress) Run(title string) error {
	if p.isFinished {
		return errors.New("Progress is finished")
	}

	if p.isRunning {
		return errors.New("Progress is running")
	}

	p.title = title
	p.startedAt = time.Now()
	p.isRunning = true

	go func() {
		var d time.Duration
		if p.UpdateInterval > 0 {
			d = p.UpdateInterval
		} else {
			d = time.Millisecond * 250
		}
		for {
			select {
			case <-p.ctx.Done():
				break
			case <-time.Tick(d):
				p.Lock()
				p.print()
				p.Unlock()
			}
		}
	}()

	return nil
}

func (p *Progress) print() {
	fmt.Fprintf(os.Stderr, "%s[%v/%v] - %v%% %s", CLEAR_LINE, p.current, p.max,
		p.current*100/p.max, p.title)
}

func (p *Progress) Incr(n int) error {
	p.Lock()
	defer p.Unlock()

	if p.max < p.current+n {
		return errors.New("Current value is out of bound")
	}

	p.current += n
	return nil
}

func (p *Progress) Finish() error {
	if p.isFinished {
		return errors.New("Progress is finished")
	}

	if !p.isRunning {
		return errors.New("Progress is not running")
	}

	p.cancel()
	p.isFinished = true
	p.isRunning = false

	used := humanReadableDuration(p.startedAt)
	uss := time.Now().Unix() - p.startedAt.Unix()
	var ops int64
	if uss == 0 {
		ops = int64(p.current)
	} else {
		ops = int64(p.current) / uss
	}

	fmt.Fprintf(os.Stderr, "%sDone, used %s, %v op/s.\n", CLEAR_LINE, used, ops)
	return nil
}

func humanReadableDuration(t time.Time) string {
	ms := time.Now().UnixMilli() - t.UnixMilli()
	if ms < 0 {
		ms = -ms
	}

	if ms < 1000 {
		return fmt.Sprintf("%vms", ms)
	}

	if ms = ms / 1000; ms < 60 {
		return fmt.Sprintf("%vs", ms)
	}

	if ms = ms / 60; ms < 60 {
		return fmt.Sprintf("%vmin", ms)
	}

	if ms = ms / 60; ms < 24 {
		return fmt.Sprintf("%vh", ms)
	}

	ms /= 24
	return fmt.Sprintf("%vday", ms)
}
