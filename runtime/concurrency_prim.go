// Go concurrency primitives
//
// "go" statement support
// channels

package runtime

import (
	"math"
	"math/big"
	. "sint/core"
)

func initConcurrencyPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "sint:go", 1, false, primGo)
	addPrimitive(ctx, "goroutine-id", 0, false, primGoroutineId)
	addPrimitive(ctx, "make-channel", 0, true, primMakeChannel)
	addPrimitive(ctx, "channel?", 0, true, primChannelp)
	addPrimitive(ctx, "channel-send", 2, false, primChannelSend)
	addPrimitive(ctx, "channel-receive", 1, false, primChannelReceive)
	addPrimitive(ctx, "channel-length", 1, false, primChannelLength)
	addPrimitive(ctx, "channel-capacity", 1, false, primChannelCapacity)
	addPrimitive(ctx, "close-channel", 1, false, primCloseChannel)
}

func primGo(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	err := ctx.InvokeConcurrent(a0)
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return ctx.UnspecifiedVal, 1
}

func primGoroutineId(ctx *Scheme, _, _ Val, rest []Val) (Val, int) {
	return ctx.GoroutineId, 1
}

func primMakeChannel(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	capacity := 0
	if a0 != ctx.UndefinedVal {
		iv, ok := a0.(*big.Int)
		if !ok || !iv.IsInt64() || iv.Int64() < 0 || iv.Int64() > math.MaxInt {
			return ctx.Error("make-channel: Invalid capacity", a0)
		}
		capacity = int(iv.Int64())
	}
	return &Chan{Ch: make(chan Val, capacity)}, 1
}

func primChannelp(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if _, ok := a0.(*Chan); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primChannelSend(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if ch, ok := a0.(*Chan); ok {
		ch.Ch <- a1
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("channel-send: not a channel", a0)
}

func primChannelReceive(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if ch, ok := a0.(*Chan); ok {
		v, ok := <-ch.Ch
		if !ok {
			ctx.MultiVals = []Val{ctx.FalseVal}
			return ctx.UnspecifiedVal, 2
		}
		ctx.MultiVals = []Val{ctx.TrueVal}
		return v, 2
	}
	return ctx.Error("channel-receive: not a channel", a0)
}

func primChannelLength(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if ch, ok := a0.(*Chan); ok {
		return big.NewInt(int64(len(ch.Ch))), 1
	}
	return ctx.Error("channel-length: not a channel", a0)
}

func primChannelCapacity(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if ch, ok := a0.(*Chan); ok {
		return big.NewInt(int64(cap(ch.Ch))), 1
	}
	return ctx.Error("channel-capacity: not a channel", a0)
}

func primCloseChannel(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if ch, ok := a0.(*Chan); ok {
		close(ch.Ch)
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("close-channel: not a channel", a0)
}
