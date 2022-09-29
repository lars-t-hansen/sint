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

func primGo(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	err := ctx.InvokeConcurrent(a0)
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return ctx.UnspecifiedVal, 1
}

func primGoroutineId(ctx *Scheme, _, _ Val, _ []Val) (Val, int) {
	return ctx.GoroutineId, 1
}

func primMakeChannel(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	capacity := 0
	if a0 != ctx.UndefinedVal {
		// TODO: this is checkExactIntInRange, and integer->char does the same thing
		iv, ok := a0.(*big.Int)
		if !ok || !iv.IsInt64() || iv.Int64() < 0 || iv.Int64() > math.MaxInt {
			return ctx.Error("make-channel: Invalid capacity", a0)
		}
		capacity = int(iv.Int64())
	}
	return &Chan{Ch: make(chan Val, capacity)}, 1
}

func primChannelp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Chan); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primChannelSend(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	ch, err := checkChannel(ctx, a0, "channel-send")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	ch <- a1
	return ctx.UnspecifiedVal, 1
}

func primChannelReceive(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	ch, err := checkChannel(ctx, a0, "channel-receive")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	v, ok := <-ch
	if !ok {
		ctx.MultiVals = []Val{ctx.FalseVal}
		return ctx.UnspecifiedVal, 2
	}
	ctx.MultiVals = []Val{ctx.TrueVal}
	return v, 2
}

func primChannelLength(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	ch, err := checkChannel(ctx, a0, "channel-length")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return big.NewInt(int64(len(ch))), 1
}

func primChannelCapacity(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	ch, err := checkChannel(ctx, a0, "channel-capacity")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return big.NewInt(int64(cap(ch))), 1
}

func primCloseChannel(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	ch, err := checkChannel(ctx, a0, "close-channel")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	close(ch)
	return ctx.UnspecifiedVal, 1
}

func checkChannel(ctx *Scheme, v Val, name string) (chan Val, *WrappedError) {
	if ch, ok := v.(*Chan); ok {
		return ch.Ch, nil
	}
	return nil, ctx.WrapError(name+": not a channel", v)
}
