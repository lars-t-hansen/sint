// Evaluation engine core.

package core

import (
	"math/big"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// Well-known tls key values.  The ones below 100 are embedded in Scheme code as
// constants.

const (
	CurrentInputPort  = 1
	CurrentOutputPort = 2
	CurrentErrorPort  = 3
	ErrorHandler      = 4
	FirstUserTlsKey   = 100
)

// State shared between goroutines of the same Scheme instance.  Some of these
// values are copied into the per-goroutine state for easy access.

type SharedScheme struct {
	/////////////////////////////////////////////////////////////////////////
	//
	// oblist and nextGenSym are shared mutable state.

	// Symbol table
	//
	// The oblist uses a sync.Map because its behavior -- ever-growing cache
	// that is read much more than it is written -- fits well with that.
	oblist sync.Map // string -> symbol

	// Counter for non-interned symbol name generation
	//
	// The counter uses atomic increment operations since those work very
	// well for it.
	//
	// For Go 1.19 we can upgrade to atomic.Int32 here; for now, just use
	// atomic operators on a plain int32.
	nextGensym int32

	// Counter for tls-value keys.  Atomic as for nextGensym.
	//
	// Keys have to be globally unique since a tls key (the internal name for a parameter)
	// can be created on one thread but later used on a different thread; the name
	// must not reference some other parameter on the latter thread.
	nextTlsKey int32

	// Counter for goroutine ID.  Atomic as for nextGensym.
	nextGoroutineId int64

	//////////////////////////////////////////////////////////////////////////
	//
	// Immutable state

	// Singleton values, these are copied into the Scheme structure
	UnspecifiedVal Val
	UndefinedVal   Val
	NullVal        Val
	TrueVal        Val
	FalseVal       Val
	EofVal         Val

	// Well-known symbols.
	AndSym           *Symbol
	BeginSym         *Symbol
	CaseSym          *Symbol
	CondSym          *Symbol
	DefineSym        *Symbol
	DoSym            *Symbol
	ElseSym          *Symbol
	GoSym            *Symbol
	IfSym            *Symbol
	LambdaSym        *Symbol
	LetSym           *Symbol
	LetStarSym       *Symbol
	LetValuesSym     *Symbol
	LetStarValuesSym *Symbol
	LetrecSym        *Symbol
	OrSym            *Symbol
	ParameterizeSym  *Symbol
	QuoteSym         *Symbol
	SetSym           *Symbol
	ArrowSym         *Symbol
	DotSym           *Symbol
	NewlineSym       *Symbol
	ReturnSym        *Symbol
	TabSym           *Symbol
	SpaceSym         *Symbol
}

func newSharedScheme() *SharedScheme {
	s := &SharedScheme{
		UnspecifiedVal: &Unspecified{},
		UndefinedVal:   &Undefined{},
		NullVal:        &Null{},
		TrueVal:        &True{},
		FalseVal:       &False{},
		EofVal:         &EofObject{},
		nextGensym:     1000,
		nextTlsKey:     FirstUserTlsKey,
	}

	s.AndSym = s.Intern("and")
	s.BeginSym = s.Intern("begin")
	s.CaseSym = s.Intern("case")
	s.CondSym = s.Intern("cond")
	s.DefineSym = s.Intern("define")
	s.DoSym = s.Intern("do")
	s.ElseSym = s.Intern("else")
	s.GoSym = s.Intern("go")
	s.IfSym = s.Intern("if")
	s.LambdaSym = s.Intern("lambda")
	s.LetSym = s.Intern("let")
	s.LetStarSym = s.Intern("let*")
	s.LetValuesSym = s.Intern("let-values")
	s.LetStarValuesSym = s.Intern("let*-values")
	s.LetrecSym = s.Intern("letrec")
	s.OrSym = s.Intern("or")
	s.ParameterizeSym = s.Intern("parameterize")
	s.QuoteSym = s.Intern("quote")
	s.SetSym = s.Intern("set!")
	s.ArrowSym = s.Intern("=>")
	s.DotSym = s.Intern(".")
	s.NewlineSym = s.Intern("newline")
	s.ReturnSym = s.Intern("return")
	s.TabSym = s.Intern("tab")
	s.SpaceSym = s.Intern("space")

	return s
}

// There is one new Scheme instance per goroutine, so it needs to be fairly
// lightweight.

type Scheme struct {
	Shared *SharedScheme

	// Singleton values, lifted from "Shared"
	UnspecifiedVal Val
	UndefinedVal   Val
	NullVal        Val
	TrueVal        Val
	FalseVal       Val
	EofVal         Val

	// Useful(?) values.  TODO: Flesh this out, and use it in the emitter: Most
	// literal values in programs are 0, 1, and 2, and we could have them
	// all predefined here and could just use them rather than cons them
	// up anew every time.  That said, those are *constant* values and
	// are only consed up when the program is deserialized, not at runtime,
	// so they probably are not all that useful frankly.
	Zero  *big.Int
	One   *big.Int
	OneMM *big.Int
	Zerof *big.Float
	Onef  *big.Float
	Halff *big.Float

	GoroutineId *big.Int

	///////////////////////////////////////////////////////////////////////////////
	//
	// Per-thread mutable state

	// Last-ditch unwind reporter which the engine will use to report unwinds on goroutines
	// if they propagate past any installed error handler.  A goroutine inherits the
	// unwind handler from its parent goroutine.
	UnwindReporter func(*Scheme, *UnwindPkg)

	// This is interpreted in the context of the number-of-values flag passed back
	// in the evaluator.
	MultiVals []Val

	// The tls store is used for parameter values primarily, but can also be used
	// for other things.  The key is global; see comments in SharedScheme.
	// The store is initialized from the parent's store when the thread is forked.
	// Threads never merge.
	tlsValues map[int32]Val
}

// oldScheme can be nil, in which case we create a new globally shared
// SharedScheme instance.  For new goroutines, oldScheme must never be nil.

func NewScheme(oldScheme *Scheme, unwReporter func(*Scheme, *UnwindPkg)) *Scheme {
	var ss *SharedScheme
	if oldScheme != nil {
		ss = oldScheme.Shared
	} else {
		ss = newSharedScheme()
	}
	s := &Scheme{
		Shared:         ss,
		UnspecifiedVal: ss.UnspecifiedVal,
		UndefinedVal:   ss.UndefinedVal,
		NullVal:        ss.NullVal,
		TrueVal:        ss.TrueVal,
		FalseVal:       ss.FalseVal,
		EofVal:         ss.EofVal,
		Zero:           big.NewInt(0),
		One:            big.NewInt(1),
		OneMM:          big.NewInt(1000000),
		Zerof:          big.NewFloat(0.0),
		Onef:           big.NewFloat(1.0),
		Halff:          big.NewFloat(0.5),
		tlsValues:      make(map[int32]Val),
		GoroutineId:    big.NewInt(atomic.AddInt64(&ss.nextGoroutineId, 1)),
		UnwindReporter: unwReporter,
	}

	// Inherit initial parameter values from oldScheme.
	// We're currently on oldScheme's thread and can copy without synchronization.
	// TODO: Is this the best we can do for copying a map?
	if oldScheme != nil {
		for k, v := range oldScheme.tlsValues {
			s.tlsValues[k] = v
		}
	}
	return s
}

func (c *SharedScheme) Intern(s string) *Symbol {
	if v, ok := c.oblist.Load(s); ok {
		return v.(*Symbol)
	}
	// Two threads can race to add a new symbol; if the other thread won,
	// return the symbol it created instead of the one constructed here.
	sym := &Symbol{Name: s, Value: c.UndefinedVal}
	realSym, _ := c.oblist.LoadOrStore(s, sym)
	return realSym.(*Symbol)
}

func (c *Scheme) Intern(s string) *Symbol {
	return c.Shared.Intern(s)
}

func (c *SharedScheme) FindSymbolsByName(pattern string) []*Symbol {
	syms := []*Symbol{}
	c.oblist.Range(func(key, value any) bool {
		if strings.Contains(key.(string), pattern) {
			sym := value.(*Symbol)
			if sym.Value != c.UndefinedVal {
				syms = append(syms, sym)
			}
		}
		return true
	})
	return syms
}

func (c *Scheme) FindSymbolsByName(pattern string) []*Symbol {
	return c.Shared.FindSymbolsByName(pattern)
}

func (c *SharedScheme) Gensym(s string) *Symbol {
	n := atomic.AddInt32(&c.nextGensym, 1)
	name := ".G" + strconv.Itoa(int(n)) + "." + s
	return &Symbol{Name: name, Value: c.UndefinedVal}
}

func (c *Scheme) Gensym(s string) *Symbol {
	return c.Shared.Gensym(s)
}

func (c *Scheme) AllocateTlsKey() int32 {
	// TODO: Overflow checking
	return atomic.AddInt32(&c.Shared.nextTlsKey, 1)
}

func (c *Scheme) GetTlsValue(key int32) Val {
	if v, ok := c.tlsValues[key]; ok {
		return v
	}
	return c.UnspecifiedVal
}

func (c *Scheme) SetTlsValue(key int32, v Val) {
	c.tlsValues[key] = v
}

// When we unwind, eval() returns a (unwind-package, EvalUnwind) where the unwind-package
// propagates information to whomever stops the unwinding.  The primitives have the same
// convention.  However, EvalUnwind is internal to the evaluator and primitives, and
// client code will instead usually receive values embedded in the unwind-package.
//
// It is a system-wide invariant that unwinding carries an unwind-package; if
// the `number-of-values`` return value is EvalUnwind then the `value`` return
// value must be such a package.

const (
	EvalUnwind = -1
)

// This is a package holding an error while we're in parts of the system that can't easily
// signal the error using the standard mechanism.  It is converted to an unwinding error by
// SignalWrappedError.

type WrappedError struct {
	message   string
	irritants []Val
}

// Returns (unwind-object, EvalUnwind) for use in the standard error
// signalling protocol.

func (c *Scheme) Error(message string, irritants ...Val) (Val, int) {
	return c.SignalWrappedError(&WrappedError{message: message, irritants: irritants})
}

// Returns an unwind-object for use by the caller in its internal error
// signalling protocol.  That the payload is a list here is to be compatible
// with how call/cc does it; this will be cleaned up.

func (c *Scheme) WrapError(message string, irritants ...Val) *WrappedError {
	return &WrappedError{message: message, irritants: irritants}
}

// TODO: Eventually this will attempt to invoke the error handler, which will itself
// invoke an escape continuation; it will not just start an unwind with an error message.
// That will be the fallback behavior.

func (c *Scheme) SignalWrappedError(we *WrappedError) (Val, int) {
	is := c.NullVal
	for idx := len(we.irritants) - 1; idx >= 0; idx-- {
		is = &Cons{Car: we.irritants[idx], Cdr: is}
	}
	return c.NewUnwindPackage(c.FalseVal, &Cons{Car: &Str{Value: we.message}, Cdr: is}), EvalUnwind
}

// Returns an unwind-object wrapping the key and the payload.

func (c *Scheme) NewUnwindPackage(key Val, vs Val) Val {
	return &UnwindPkg{Key: key, Payload: vs}
}

// Returns (values, nil) on success, otherwise (nil, unwind-package)

func (c *Scheme) EvalToplevel(expr Code) ([]Val, Val) {
	return c.captureValues(c.eval(expr, nil))
}

func (c *Scheme) DefineToplevel(name string, v Val) {
	c.Intern(name).Value = v
}

// Returns (values, nil) on success, otherwise (nil, unwind-package)

func (c *Scheme) Invoke(proc Val, args []Val) ([]Val, Val) {
	v, k := c.invokeInternal(proc, args)
	if k == EvalUnwind {
		return nil, v
	}
	return c.captureValues(v, k)
}

// Returns nil on success, otherwise an uwind-package.

func (c *Scheme) InvokeConcurrent(proc Val) *WrappedError {
	// This is always (sint:go thunk) and there are "no" nullary primitive
	// procedures, so let's keep it simple and ban primitive procedures from
	// being used here.
	newCode, newEnv, prim, unw := c.invokeSetup(proc, []Val{})
	if unw != nil {
		return unw
	}
	if prim != nil {
		return c.WrapError("Primitive procedures cannot be invoked concurrently")
	}
	go NewScheme(c, c.UnwindReporter).evalWithUnwindReport(newCode, newEnv)
	return nil
}

// Returns nothing, but will report unwinding on the installed unwind reporter

func (c *Scheme) evalWithUnwindReport(expr Code, env *lexenv) {
	v, unw := c.eval(expr, env)
	if unw == EvalUnwind {
		if c.UnwindReporter != nil {
			c.UnwindReporter(c, v.(*UnwindPkg))
		} else {
			panic("No unwind reporter installed")
		}
	}
}

func (c *Scheme) InvokeWithUnwindHandler(filterKey Val, thunkProc *Procedure, handleProc *Procedure) (Val, int) {
	v, k := c.invokeInternal(thunkProc, []Val{})
	if k != EvalUnwind {
		return v, k
	}
	pkg := v.(*UnwindPkg)
	if filterKey == c.FalseVal || filterKey == pkg.Key {
		return c.invokeInternal(handleProc, []Val{pkg.Key, pkg.Payload})
	}
	return v, k
}

func (c *Scheme) captureValues(v Val, numVal int) ([]Val, Val) {
	if numVal == EvalUnwind {
		return nil, v
	}
	vs := make([]Val, numVal)
	vs[0] = v
	copy(vs[1:], c.MultiVals[:numVal-1])
	return vs, nil
}

func (c *Scheme) invokeSetup(proc Val, args []Val) (theCode Code, newEnv *lexenv, thePrim func(*Scheme, Val, Val, []Val) (Val, int), theErr *WrappedError) {
	if p, ok := proc.(*Procedure); ok {
		if len(args) < p.Lam.Fixed {
			theErr = c.WrapError("Not enough arguments") // TODO msg
			return
		}
		if len(args) > p.Lam.Fixed && !p.Lam.Rest {
			theErr = c.WrapError("Too many arguments") // TODO msg
			return
		}
		if p.Lam.Body == nil {
			thePrim = p.Primop
			return
		}
		// args (really the underlying vals) is freshly allocated,
		// so it's OK to use that storage here.
		if !p.Lam.Rest {
			newEnv = &lexenv{slots: args, link: p.Env}
		} else {
			// TODO: I think we can do better than this.  Since the storage
			// is fresh, we can store the rest argument in the slot after the
			// slice, if it exists, in which case we avoid copying the
			// array in the append() below.  If there is no extra slot then there's
			// at least a chance that the append() will use capacity that is there.
			newSlots := args[:p.Lam.Fixed]
			var l *Cons
			var last *Cons
			for i := p.Lam.Fixed; i < len(args); i++ {
				x := &Cons{Car: args[i], Cdr: c.NullVal}
				if l == nil {
					l = x
				}
				if last != nil {
					last.Cdr = x
				}
				last = x
			}
			if l == nil {
				newSlots = append(newSlots, c.NullVal)
			} else {
				newSlots = append(newSlots, l)
			}
			newEnv = &lexenv{slots: newSlots, link: p.Env}
		}
		theCode = p.Lam.Body
		return
	}
	theErr = c.WrapError("Invoke: Not a procedure" /*+ e.Exprs[0].String() + "\n" + proc.String()*/)
	return
}

func (c *Scheme) invokeInternal(proc Val, args []Val) (Val, int) {
	newCode, newEnv, prim, unw := c.invokeSetup(proc, args)
	if unw != nil {
		return c.SignalWrappedError(unw)
	}
	if prim != nil {
		var a0, a1 Val = c.UndefinedVal, c.UndefinedVal
		var an []Val
		if len(args) > 0 {
			a0 = args[0]
			if len(args) > 1 {
				a1 = args[1]
				if len(args) > 2 {
					an = args[2:]
				}
			}
		}
		return prim(c, a0, a1, an)
	}
	return c.eval(newCode, newEnv)
}

type lexenv struct {
	slots []Val
	link  *lexenv
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Scheme) eval(expr Code, env *lexenv) (Val, int) {
again:
	switch instr := expr.(type) {
	case *Quote:
		return instr.Value, 1
	case *If:
		test, testRes := c.eval(instr.Test, env)
		if testRes == EvalUnwind {
			return test, testRes
		}
		if test != c.FalseVal {
			expr = instr.Consequent
		} else {
			expr = instr.Alternate
		}
		goto again
	case *Begin:
		if len(instr.Exprs) == 0 {
			return c.UnspecifiedVal, 1
		}
		for _, e := range instr.Exprs[:len(instr.Exprs)-1] {
			r, nres := c.eval(e, env)
			if nres == EvalUnwind {
				return r, nres
			}
		}
		expr = instr.Exprs[len(instr.Exprs)-1]
		goto again
	case *Call:
		// We do things differently depending on the procedure, so evaluate that first
		// separately.
		maybeProc, procErr := c.eval(instr.Exprs[0], env)
		if procErr == EvalUnwind {
			return maybeProc, procErr
		}

		// Fast path for primitives to avoid allocating a temp slice in common cases.
		// Primitives take two fixed arguments and a slice, which can be nil.  If
		// an argument is not present, the value is Undefined, which is a value that
		// cannot be produced by Scheme code.
		if proc, ok := maybeProc.(*Procedure); ok && proc.Primop != nil {
			numActuals := len(instr.Exprs) - 1
			if numActuals < proc.Lam.Fixed || numActuals > proc.Lam.Fixed && !proc.Lam.Rest {
				return c.Error("Wrong number of arguments to procedure")
			}
			var arg0, arg1 Val = c.UndefinedVal, c.UndefinedVal
			var res0, res1 int
			var vals []Val
			if numActuals > 0 {
				arg0, res0 = c.eval(instr.Exprs[1], env)
				if res0 == EvalUnwind {
					return arg0, res0
				}
				if numActuals > 1 {
					arg1, res1 = c.eval(instr.Exprs[2], env)
					if res1 == EvalUnwind {
						return arg1, res1
					}
					if numActuals > 2 {
						var eUnw Val
						vals, eUnw = c.evalExprs(instr.Exprs[3:], env)
						if eUnw != nil {
							return eUnw, EvalUnwind
						}
					}
				}
			}
			return proc.Primop(c, arg0, arg1, vals)
		}
		// For now, unavoidable use of evalExprs.  Mostly this is being reused for the rib,
		// so it's OK.
		args, eUnw := c.evalExprs(instr.Exprs[1:], env)
		if eUnw != nil {
			return eUnw, EvalUnwind
		}
		newCode, newEnv, prim, iUnw := c.invokeSetup(maybeProc, args)
		if iUnw != nil {
			return c.SignalWrappedError(iUnw)
		}
		if prim != nil {
			panic("Should not happen")
		}
		expr = newCode
		env = newEnv
		goto again
	case *Apply:
		proc, procRes := c.eval(instr.Proc, env)
		if procRes == EvalUnwind {
			return proc, procRes
		}
		argList, argRes := c.eval(instr.Args, env)
		if argRes == EvalUnwind {
			return argList, argRes
		}
		args := []Val{}
		for {
			if argList == c.NullVal {
				break
			}
			a, ok := argList.(*Cons)
			if !ok {
				return c.Error("sint:apply: Not a list", argList)
			}
			args = append(args, a.Car)
			argList = a.Cdr
		}
		newCode, newEnv, prim, unw := c.invokeSetup(proc, args)
		if unw != nil {
			return c.SignalWrappedError(unw)
		}
		if prim != nil {
			var a0, a1 Val = c.UndefinedVal, c.UndefinedVal
			var an []Val
			if len(args) > 0 {
				a0 = args[0]
				if len(args) > 1 {
					a1 = args[1]
					if len(args) > 2 {
						an = args[2:]
					}
				}
			}
			return prim(c, a0, a1, an)
		}
		expr = newCode
		env = newEnv
		goto again
	case *Lambda:
		return &Procedure{Lam: instr, Env: env, Primop: nil}, 1
	case *Let:
		// evalExprs is OK here because the allocation is reused for the rib, which will always
		// have one.
		vals, unw := c.evalExprs(instr.Exprs, env)
		if unw != nil {
			return unw, EvalUnwind
		}
		newEnv := &lexenv{slots: vals, link: env}
		expr = instr.Body
		env = newEnv
		goto again
	case *LetValues:
		// Basically, create a rib with the required number of slots
		// Then evaluate the exprs in order in the old env and assign values to slots, throwing if
		// an expression returns the wrong number of values for the corresponding binding
		// Then evaluate the body in that environment
		panic("LetValues not implemented")
	case *LetStar:
		slotvals := make([]Val, len(instr.Exprs))
		for i := range slotvals {
			slotvals[i] = c.UnspecifiedVal
		}
		newEnv := &lexenv{slots: slotvals, link: env}
		for i, e := range instr.Exprs {
			v, unw := c.eval(e, newEnv)
			if unw == EvalUnwind {
				return v, unw
			}
			slotvals[i] = v
		}
		expr = instr.Body
		env = newEnv
		goto again
	case *Letrec:
		slotvals := make([]Val, len(instr.Exprs))
		for i := range slotvals {
			slotvals[i] = c.UnspecifiedVal
		}
		newEnv := &lexenv{slots: slotvals, link: env}
		// Specialize this to avoid allocating a temp slice for evalExprs().  This evaluates
		// expressions in reverse order but that's OK, the order is unspecified.
		if len(slotvals) < 4 {
			var t0, t1, t2, t3 Val
			var r0, r1, r2, r3 int
			if len(slotvals) > 0 {
				if len(slotvals) > 1 {
					if len(slotvals) > 2 {
						if len(slotvals) > 3 {
							t3, r3 = c.eval(instr.Exprs[3], newEnv)
							if r3 == EvalUnwind {
								return t3, r3
							}
						}
						t2, r2 = c.eval(instr.Exprs[2], newEnv)
						if r2 == EvalUnwind {
							return t2, r2
						}
					}
					t1, r1 = c.eval(instr.Exprs[1], newEnv)
					if r1 == EvalUnwind {
						return t1, r1
					}
				}
				t0, r0 = c.eval(instr.Exprs[0], newEnv)
				if r0 == EvalUnwind {
					return t0, r0
				}
			}
			if len(slotvals) > 0 {
				if len(slotvals) > 1 {
					if len(slotvals) > 2 {
						if len(slotvals) > 3 {
							slotvals[3] = t3
						}
						slotvals[2] = t2
					}
					slotvals[1] = t1
				}
				slotvals[0] = t0
			}
		} else {
			// evalExprs and the intermediate allocation are unavoidable here, though if it
			// becomes really hot we can consider adding more cases above.
			vals, unw := c.evalExprs(instr.Exprs, newEnv)
			if unw != nil {
				return unw, EvalUnwind
			}
			copy(slotvals, vals)
		}
		expr = instr.Body
		env = newEnv
		goto again
	case *Lexical:
		rib := env
		for levels := instr.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		return rib.slots[instr.Offset], 1
	case *Setlex:
		rhs, rhsRes := c.eval(instr.Rhs, env)
		if rhsRes == EvalUnwind {
			return rhs, rhsRes
		}
		rib := env
		for levels := instr.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		rib.slots[instr.Offset] = rhs
		return c.UnspecifiedVal, 1
	case *Global:
		val := instr.Name.Value
		if val == c.UndefinedVal {
			return c.Error("Undefined global variable", instr.Name)
		}
		return val, 1
	case *Setglobal:
		rhs, rhsRes := c.eval(instr.Rhs, env)
		if rhsRes == EvalUnwind {
			return rhs, rhsRes
		}
		instr.Name.Value = rhs
		return c.UnspecifiedVal, 1
	default:
		panic("Bad expression: " + expr.String())
	}
}

// Returns either (values, nil) or (nil, unwind-object)

func (c *Scheme) evalExprs(es []Code, env *lexenv) ([]Val, Val) {
	vs := make([]Val, len(es))
	for i, e := range es {
		r, nres := c.eval(e, env)
		if nres == EvalUnwind {
			return nil, r
		}
		vs[i] = r
	}
	return vs, nil
}
