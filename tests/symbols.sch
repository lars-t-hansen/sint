(display "Testing: symbols.sch\n")

(assert (symbol? 'foo) "symbol type #1")
(assert-not (symbol? "foo") "symbol type #2")

(assert (eq? 'foo 'foo) "symbol equality #1")
(assert (eq? 'foo (string->symbol (symbol->string 'foo))) "symbol equality #2")

(assert-not (eq? (gensym) (gensym)) "gensym")

(assert (symbol=? 'foo 'foo (string->symbol "foo")) "symbol=? #1")
(assert-not (symbol=? 'foo 'foo (string->symbol "Foo")) "symbol=? #2")

(display "OK\n")
