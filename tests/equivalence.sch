(display "Testing: equivalence.sch\n")

(assert (equal? "a" "a") "equal? #1")
(assert-not (equal? "a" "b") "equal? #2")
(assert (equal? '(a b c) '(a b c)) "equal? #3")
(assert-not (equal? '(a b c) '(a b c d)) "equal? #4")

(display "OK\n")
