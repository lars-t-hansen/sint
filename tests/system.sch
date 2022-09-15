(display "Testing: system.sch\n")

(assert (not (null? (memq 'sint (features)))) "features")

(load "tests/define-x.sch")
(assert (= defined-x 37) "load")

(display "OK\n")
