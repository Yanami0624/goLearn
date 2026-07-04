package concurrent

import (
	"context"
	"sync"
	"time"
)

type Packet struct {
	idx int
	e   int
}

var channel chan Packet

type Buffer struct {
	N    int
	buf  []int
	head int
	tail int
	cap  int
}

var lock sync.Locker
var full sync.Cond
var empt sync.Cond

func NewBufffer(n int) Buffer {
	N := max(min(n, 128), 2)
	buf := make([]int, n)
	head := 0
	tail := head + 1
	cap := 0

	channel = make(chan Packet, N)

	return Buffer{N, buf, head, tail, cap}
}

func (b *Buffer) Add(ctx context.Context) {
	select {
	case <-time.After(time.Second * 10):
		return
	case <-ctx.Done():
		return
	default:
		lock.Lock()
		defer lock.Unlock()
		for b.cap == b.N-1 {
			full.Wait()
		}

		b.buf[b.tail] = b.tail
		b.tail = (b.tail + 1) % b.N
		b.cap++

		if b.cap == 1 {
			empt.Signal()
		}
	}
}

func (b *Buffer) Send(ctx context.Context) {
	select {
	case <-time.After(time.Second * 10):
		return
	case <-ctx.Done():
		return
	default:
		lock.Lock()
		defer lock.Unlock()
		for b.cap == 0 {
			empt.Wait()
		}

		idx := (b.head + 1) % b.N
		e := b.buf[idx]
		pkt := Packet{idx, e}
		channel <- pkt

		b.head = idx
		b.cap--

		if b.cap == b.N-2 {
			full.Signal()
		}
	}
}

func (b *Buffer) Ack(ctx context.Context) {
	select {
	case <-time.After(time.Second * 10):
		return
	case <-ctx.Done():
		return
	case pkt := <-channel:
		
	}
}
