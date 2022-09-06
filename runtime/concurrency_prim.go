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

func initConcurrencyPrimitives(c *Scheme) {
	addPrimitive(c, "sint:go", 1, false, primGo)
	addPrimitive(c, "make-channel", 0, true, primMakeChannel)
	addPrimitive(c, "channel?", 0, true, primChannelp)
	addPrimitive(c, "channel-send", 2, false, primChannelSend)
	addPrimitive(c, "channel-receive", 1, false, primChannelReceive)
	addPrimitive(c, "channel-length", 1, false, primChannelLength)
	addPrimitive(c, "channel-capacity", 1, false, primChannelCapacity)
	addPrimitive(c, "close-channel", 1, false, primCloseChannel)
}

func primGo(ctx *Scheme, args []Val) (Val, int) {
	ctx.InvokeConcurrent(args[0])
	return ctx.UnspecifiedVal, 1
}

func primMakeChannel(ctx *Scheme, args []Val) (Val, int) {
	capacity := 0
	if len(args) > 0 {
		iv, ok := args[0].(*big.Int)
		if !ok || !iv.IsInt64() || iv.Int64() < 0 || iv.Int64() > math.MaxInt {
			return ctx.Error("make-channel: Invalid capacity: " + args[0].String())
		}
		capacity = int(iv.Int64())
	}
	return &Chan{Ch: make(chan Val, capacity)}, 1
}

func primChannelp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Chan); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primChannelSend(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Chan); ok {
		ch.Ch <- args[1]
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("channel-send: not a channel: " + args[0].String())
}

func primChannelReceive(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Chan); ok {
		v, ok := <-ch.Ch
		if !ok {
			ctx.MultiVals = []Val{ctx.FalseVal}
			return ctx.UnspecifiedVal, 2
		}
		ctx.MultiVals = []Val{ctx.TrueVal}
		return v, 2
	}
	return ctx.Error("channel-receive: not a channel: " + args[0].String())
}

func primChannelLength(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Chan); ok {
		return big.NewInt(int64(len(ch.Ch))), 1
	}
	return ctx.Error("channel-length: not a channel: " + args[0].String())
}

func primChannelCapacity(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Chan); ok {
		return big.NewInt(int64(cap(ch.Ch))), 1
	}
	return ctx.Error("channel-capacity: not a channel: " + args[0].String())
}

func primCloseChannel(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Chan); ok {
		close(ch.Ch)
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("close-channel: not a channel: " + args[0].String())
}
