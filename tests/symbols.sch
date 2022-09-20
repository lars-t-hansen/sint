(display "Testing: symbols.sch\n")

(assert (symbol? 'foo) "symbol type #1")
(assert-not (symbol? "foo") "symbol type #2")

(assert (eq? 'foo 'foo) "symbol equality #1")
(assert (eq? 'foo (string->symbol (symbol->string 'foo))) "symbol equality #2")

(assert-not (eq? (gensym) (gensym)) "gensym")

(assert (symbol=? 'foo 'foo (string->symbol "foo")) "symbol=? #1")
(assert-not (symbol=? 'foo 'foo (string->symbol "Foo")) "symbol=? #2")

(let ((xs (apropos "char")))
  (assert (and (memq 'char->integer xs)
	       (memq 'integer->char xs)
	       (memq 'char? xs))
	  "apropos #1"))

(assert (null? (apropos 'fnord)) "apropos #2")

(display "OK\n")
