(display "Testing: numbers.sch\n")

(define *inf* (/ 1 0))

(assert (finite? 37) "finite? #1")
(assert (finite? 42.5) "finite? #2")
(assert-not (finite? *inf*) "finite? #3")

(assert (infinite? *inf*) "infinite? #1")
(assert-not (infinite? 37) "infinite? #2")
(assert-not (infinite? 42.5) "infinite? #3")

(assert (number? 1) "number? #1")
(assert (number? 3.4) "number? #2")
(assert (number? *inf*) "number? #3")
(assert-not (number? #\a) "number? #4")

(assert (complex? 1) "complex? #1")
(assert (complex? 3.4) "complex? #2")
(assert (complex? *inf*) "complex? #3")
(assert-not (complex? #\a) "complex? #4")

(assert (real? 1) "real? #1")
(assert (real? 3.4) "real? #2")
(assert (real? *inf*) "real? #3")
(assert-not (real? #\a) "real? #4")

(assert (rational? 1) "rational? #1")
(assert (rational? 3.4) "rational? #2")
;; Issue #29
;;(assert-not (rational? *inf*) "rational? #3")
(assert-not (rational? #\a) "rational? #4")

(assert (integer? 1) "integer? #1")
(assert-not (integer? 1.5) "integer? #2")
(assert-not (integer? *inf*) "integer? #3")
;; Issue #28
;;(assert (integer? 1.0) "integer? #4")

(assert (= (real-part 1) 1) "real-part #1")
(assert (= (imag-part 1) 0) "imag-part #1")

(assert (exact? 1) "exact? #1")
(assert-not (exact? 1.0) "exact? #2")
(assert-not (exact? *inf*) "exact? #3")

(assert (exact-integer? 1) "exact-integer? #1")
(assert-not (exact-integer? 1.0) "exact-integer? #2")
(assert-not (exact-integer? *inf*) "exact-integer? #3")

(assert (inexact? 1.0) "inexact? #1")
(assert (inexact? *inf*) "inexact? #2")
(assert-not (inexact? 3) "inexact? #3")

(assert (let ((x (exact 1.0))) (and (= x 1) (exact? x))) "exact #1")
(assert (let ((x (exact 1))) (and (= x 1) (exact? x))) "exact #2")

(assert (let ((x (inexact 1.0))) (and (= x 1) (inexact? x))) "inexact #1")
(assert (let ((x (inexact 1))) (and (= x 1) (inexact? x))) "inexact #2")

(assert (= 4 (square 2)) "square")

;; There is no NaN value, so it's always false
(assert-not (nan? 1) "nan? #1")
(assert-not (nan? *inf*) "nan? #2")

(assert (zero? 0) "zero? #1")
(assert-not (zero? *inf*) "zero? #2")
(assert (zero? 0.0) "zero? #3")
(assert-not (zero? 1) "zero? #4")

(assert (positive? 1) "positive? #1")
(assert-not (positive? 0) "positive? #2")

(assert (negative? (- 1)) "negative? #1")
(assert-not (negative? (- 0)) "negative? #2")

;; Issue #4
;;(assert (odd? 1) "odd? #1")
;;(assert-not (odd? 2) "odd? #2")
;;(assert (even? 2) "even? #1")
;;(assert-not (even? 1) "even? #2")

(assert (let ((m (max 1 1.5 2 (- 7)))) (and (= m 2) (inexact? m))) "max #1")
(assert (let ((m (max 1 2 (- 7)))) (and (= m 2) (exact? m))) "max #2")

(assert (let ((m (min 1 1.5 2 (- 7)))) (and (= m (- 7)) (inexact? m))) "min #1")
(assert (let ((m (min 1 2 (- 7)))) (and (= m (- 7)) (exact? m))) "min #2")

(display "OK\n")
