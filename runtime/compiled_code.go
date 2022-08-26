// Initialize code that has been compiled from Scheme and included in sint.
// These are generally pairs of .sch/.go files in the runtime/ directory.

package runtime

import . "sint/core"

func InitCompiled(c *Scheme) {
	// It's fine for these to be in alpha order, I think; they had better
	// not reference each other during initialization.
	initBooleans(c)
	initControl(c)
	initEquivalence(c)
	initExceptions(c)
	initNumbers(c)
	initSymbols(c)
}
