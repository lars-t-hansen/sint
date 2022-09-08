// Evaluation engine core

package core

import (
	"math/big"
	"strconv"
	"sync"
	"sync/atomic"
)

// Well-known tls key values.  The ones below 100 are embedded in Scheme code as
// constants.
const (
	CurrentInputPort  = 1
	CurrentOutputPort = 2
	CurrentErrorPort  = 3
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
	Zero *big.Int

	GoroutineId *big.Int

	///////////////////////////////////////////////////////////////////////////////
	//
	// Per-thread mutable state

	// This is interpreted in the context of the number-of-values flag passed back
	// in the evaluator
	MultiVals []Val

	// The tls store is used for parameter values primarily, but can also be used
	// for other things.  The key is global; see comments in SharedScheme.
	// The store is initialized from the parent's store when the thread is forked.
	//  Threads never merge.
	tlsValues map[int32]Val
}

// oldScheme can be nil, in which case we create a new globally shared
// SharedScheme instance.  For new goroutines, oldScheme must never be nil.
func NewScheme(oldScheme *Scheme) *Scheme {
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
		tlsValues:      make(map[int32]Val),
		GoroutineId:    big.NewInt(atomic.AddInt64(&ss.nextGoroutineId, 1)),
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
	sym := &Symbol{Name: s, Value: c.UndefinedVal}
	c.oblist.Store(s, sym)
	return sym
}

func (c *Scheme) Intern(s string) *Symbol {
	return c.Shared.Intern(s)
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

// Returns (unwind-object, EvalUnwind) for use in the standard error
// signalling protocol.
//
// TODO: Eventually this will invoke the error handler, which will itself
// invoke an escape continuation; it will not just start an unwind with an
// error message.
func (c *Scheme) Error(message string) (Val, int) {
	return c.WrapError(message), EvalUnwind
}

// Returns an unwind-object for use by the caller in its internal error
// signalling protocol.
func (c *Scheme) WrapError(message string) Val {
	return c.NewUnwindPackage(c.FalseVal, []Val{&Str{Value: message}})
}

// Returns an unwind-object wrapping the key and the values; the values
// are first consed into a list.  The unwinding is not necessarily for
// an error, it could be for invoking a captured continuation.
func (c *Scheme) NewUnwindPackage(key Val, vs []Val) Val {
	l := c.NullVal
	for i := len(vs) - 1; i >= 0; i-- {
		l = &Cons{Car: vs[i], Cdr: l}
	}
	return &UnwindPkg{Key: key, Payload: l}
}

// Returns (values, nil) on success, otherwise (nil, unwind-object)
func (c *Scheme) EvalToplevel(expr Code) ([]Val, Val) {
	return c.captureValues(c.eval(expr, nil))
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
func (c *Scheme) InvokeConcurrent(proc Val) Val {
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
	go NewScheme(c).eval(newCode, newEnv)
	return nil
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
	vs := []Val{v}
	if numVal > 1 {
		vs = append(vs, c.MultiVals[:numVal-1]...)
	}
	return vs, nil
}

func (c *Scheme) invokeSetup(proc Val, args []Val) (theCode Code, newEnv *lexenv, thePrim func(*Scheme, []Val) (Val, int), theErr Val) {
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
		return unw, EvalUnwind
	}
	if prim != nil {
		return prim(c, args)
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
		_, unw := c.evalExprs(instr.Exprs[:len(instr.Exprs)-1], env)
		if unw != nil {
			return unw, EvalUnwind
		}
		expr = instr.Exprs[len(instr.Exprs)-1]
		goto again
	case *Call:
		vals, eUnw := c.evalExprs(instr.Exprs, env)
		if eUnw != nil {
			return eUnw, EvalUnwind
		}
		maybeProc := vals[0]
		args := vals[1:]
		newCode, newEnv, prim, iUnw := c.invokeSetup(maybeProc, args)
		if iUnw != nil {
			return iUnw, EvalUnwind
		}
		if prim != nil {
			return prim(c, args)
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
				return c.Error("sint:apply: Not a list") // TODO: msg
			}
			args = append(args, a.Car)
			argList = a.Cdr
		}
		newCode, newEnv, prim, unw := c.invokeSetup(proc, args)
		if unw != nil {
			return unw, EvalUnwind
		}
		if prim != nil {
			return prim(c, args)
		}
		expr = newCode
		env = newEnv
		goto again
	case *Lambda:
		return &Procedure{Lam: instr, Env: env, Primop: nil}, 1
	case *Let:
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
	case *Letrec:
		// TODO: Probably there's a more efficient way to do this?  Note we need
		// fresh storage, so at a minimum we need to copy out of a master slice of
		// undefined values.
		slotvals := []Val{}
		for i := 0; i < len(instr.Exprs); i++ {
			slotvals = append(slotvals, c.UnspecifiedVal)
		}
		newEnv := &lexenv{slots: slotvals, link: env}
		vals, unw := c.evalExprs(instr.Exprs, newEnv)
		if unw != nil {
			return unw, EvalUnwind
		}
		for i, v := range vals {
			slotvals[i] = v
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
			return c.Error("Undefined global variable '" + instr.Name.Name + "'")
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
	vs := []Val{}
	for _, e := range es {
		r, nres := c.eval(e, env)
		if nres == EvalUnwind {
			return nil, r
		}
		vs = append(vs, r)
	}
	return vs, nil
}
