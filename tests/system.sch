(display "Testing: system.sch\n")

(assert (not (null? (memq 'sint (features)))) "features")

(load "tests/define-x.sch")
(assert (= defined-x 37) "load")

(letrec ((fib
	  (lambda (n)
	    (if (< n 2)
		n
		(+ (fib (- n 1)) (fib (- n 2)))))))
  (let* ((then (current-jiffy))
	 (r (fib 10))
	 (now (current-jiffy)))
    (assert (= r 55) "fib #1")
    (assert (exact-integer? then) "current-jiffy #1")
    (assert (exact-integer? now) "current-jiffy #2")
    (assert (> now then) "current-jiffy #3")
    (assert (exact-integer? (jiffies-per-second)) "jiffies-per-second #1")
    (assert (> (jiffies-per-second) 0) "jiffies-per-second #2")))

(display "OK\n")
