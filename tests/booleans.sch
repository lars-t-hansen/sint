(display "Testing: booleans.sch\n")

(define (assert x msg)
  (if (not x)
      (error msg)))

(define (assert-not x msg)
  (if x
      (error msg)))

(assert (boolean? #t) "boolean? #1")
(assert (boolean? #f) "boolean? #2")
(assert-not (boolean? 37) "boolean? #3")

(assert (boolean=? #t #t) "boolean=? #1")
(assert-not (boolean=? #t #f) "boolean=? #2")
(assert (boolean=? #f #f) "boolean=? #3")
(assert (boolean=? #f #f #f #f #f) "boolean=? #4")
(assert-not (boolean=? #f #f #f #t #f) "boolean=? #5")

(display "OK\n")
