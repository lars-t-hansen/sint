// Initialize code that has been compiled from Scheme and included in sint.
// These are generally pairs of .sch/.go files in the runtime/ directory.

package runtime

import . "sint/core"

func InitCompiled(c *Scheme) {
	// Fundamental stuff.  These should not reference each other during
	// initialization and can be in alpha order.
	initBooleans(c)
	initControl(c)
	initEquivalence(c)
	initExceptions(c)
	initNumbers(c)
	initPairs(c)
	initStrings(c)
	initSymbols(c)
	initSystem(c)

	// Higher-level stuff.  These can reference definitions from the previous set
	// during initialization.
	initIo(c)
}
