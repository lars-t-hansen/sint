(display "Testing: numbers.sch\n")

(define *inf* (/ 1 0))

(assert (= (+) 0) "+ #1")
(assert (= (+ 1) 1) "+ #2")
(assert (= (+ 1 2) 3) "+ #3")
(assert (= (+ 1 2 3 4 5 6 7) 28) "+ #4")
(assert (= (+ 1 2 3 4 5 6.0 7) 28) "+ #5")
(assert (inexact? (+ 1 2 3 4 5 6.0 7)) "+ #6")
(assert (exact? (+ 1 2 3 4 5 6 7)) "+ #7")

(assert (= (- 1) -1) "- #1")
(assert (= (- 1 3) -2) "- #2")
(assert (= (- 1 3 4 5 6) -17) "- #3")
(assert (exact? (- 1 3 4 5 6)) "- #4")
(assert (= (- 1 3 4 5.0 6) -17) "- #5")
(assert (inexact? (- 1 3 4 5.0 6)) "- #6")

(assert (= (*) 1) "* #1")
(assert (= (* 3) 3) "* #2")
(assert (= (* 3 4) 12) "* #3")
(assert (= (* 1 2 3 4 5 6 7) 5040) "* #4")
(assert (= (* 1 2 3 4 5 6.0 7) 5040) "* #5")
(assert (inexact? (* 1 2 3 4 5 6.0 7)) "* #6")
(assert (exact? (* 1 2 3 4 5 6 7)) "* #7")

(assert (= (/ 2) 0.5) "/ #1")
(assert (= (/ 2 4) 0.5) "/ #2")
(assert (= (/ 4 2 2) 1) "/ #3")
;; Not really a bug as we have no exact rationals, though one could argue that
;; the integer answer here ought to be exact.
;;(assert (exact? (/ 4 2 2)) "/ #4")
(assert (= (/ 4 2 2.0) 1) "/ #5")
(assert (inexact? (/ 4 2 2.0)) "/ #6")

(assert (= 1 1) "= #1")
(assert-not (= 1 2) "= #2")
(assert (= 1 1.0) "= #3")
(assert (= 1 1 1 1 1 1.0 1 1 1) "= #4")
(assert-not (= 1 1 1 1 1 1.5 1 1 1) "= #5")

(assert (< 1 2) "< #1")
(assert-not (< 2 2) "< #2")
(assert-not (< 3 2) "< #3")
(assert (< 1 2.0) "< #4")
(assert-not (< 2 2.0) "< #5")
(assert-not (< 3 2.0) "< #6")
(assert (< 1 2 3 4 5) "< #7")
(assert-not (< 1 2 3 1 5) "< #8")
	   
(assert (<= 1 2) "<= #1")
(assert (<= 2 2) "<= #2")
(assert-not (<= 3 2) "<= #3")
(assert (<= 1 2.0) "<= #4")
(assert (<= 2 2.0) "<= #5")
(assert-not (<= 3 2.0) "<= #6")
(assert (<= 1 2 3 4 5) "<= #7")
(assert (<= 1 2 3 3 5) "<= #8")
(assert-not (<= 1 2 3 1 5) "<= #9")
	   
(assert (> 2 1) "> #1")
(assert-not (> 2 2) "> #2")
(assert-not (> 2 3) "> #3")
(assert (> 2.0 1) "> #4")
(assert-not (> 2 2.0) "> #5")
(assert-not (> 2.0 3) "> #6")
(assert (> 5 4 3 2 1) "> #7")
(assert-not (> 5 1 3 2 1) "> #8")

(assert (>= 2 1) ">= #1")
(assert (>= 2 2) ">= #2")
(assert-not (>= 2 3) ">= #3")
(assert (>= 2.0 1) ">= #4")
(assert (>= 2 2.0) ">= #5")
(assert-not (>= 2.0 3) ">= #6")
(assert (>= 5 4 3 2 1) ">= #7")
(assert-not (>= 5 1 3 2 1) ">= #8")

(assert (= (string->number "37.5") 37.5) "number->string #1")
(assert (= (string->number "3.75e+1") 37.5) "number->string #2")
(assert (= (string->number "-31415") -31415) "number->string #3")
(assert (= (string->number "+375e-1") 37.5) "number->string #4")
;; Issue #32
;;(assert (infinite? (string->number "+inf.0")) "number->string #5")

(assert (string=? (number->string -1234) "-1234") "number->string #1")
(assert (string=? (number->string 1234.5) "1234.5") "number->string #2")

(assert (= (inexact 37) 37.0) "inexact #1")
(assert (inexact? (inexact 37)) "inexact #2")
(assert (inexact? (inexact 37.0)) "inexact #3")

(assert (= (exact 37.0) 37) "exact #1")
(assert (exact? (exact 37.0)) "exact #2")
(assert (exact? (exact 37)) "exact #3")

(assert (= 37 (abs -37)) "abs #1")
(assert (= 37.5 (abs -37.5)) "abs #2")
(assert (= 37.5 (abs 37.5)) "abs #3")

(assert (= -37 (floor -36.7)) "floor #1")
(assert (= 12 (floor 12.3)) "floor #2")

(assert (= -36 (ceiling -36.7)) "ceiling #1")
(assert (= 13 (ceiling 12.3)) "ceiling #2")

(assert (= -36 (truncate -36.7)) "truncate #1")
(assert (= 12 (truncate 12.3)) "truncate #2")

(assert (= -36 (round -36.5)) "round #1")
(assert (= 12 (round 12.5)) "round #2")
;; Issue #2
;;(assert (= -36 (round -35.5)) "round #3")
;;(assert (= 12 (round 11.5)) "round #4")

(assert (= 6 (bitwise-and 7 -2)) "bitwise-and #1")

(assert (= -2 (bitwise-or -4 2)) "bitwise-or #1")

(assert (= -2 (bitwise-xor -4 2)) "bitwise-xor #1")
(assert (= 4 (bitwise-xor 7 3)) "bitwise-xor #2")

(assert (= 9 (bitwise-and-not 13 6)) "bitwise-and-not #1") 

(assert (= -2 (bitwise-not 1)) "bitwise-not #1")

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
(assert-not (exact-integer? 1.5) "exact-integer? #4")

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
(assert (zero? -0.0) "zero? #5")

(assert (positive? 1) "positive? #1")
(assert-not (positive? 0) "positive? #2")

(assert (negative? -1) "negative? #1")
(assert-not (negative? -0.0) "negative? #2")

;; Issue #4
;;(assert (odd? 1) "odd? #1")
;;(assert-not (odd? 2) "odd? #2")
;;(assert (even? 2) "even? #1")
;;(assert-not (even? 1) "even? #2")

(assert (let ((m (max 1 1.5 2 -7))) (and (= m 2) (inexact? m))) "max #1")
(assert (let ((m (max 1 2 -7))) (and (= m 2) (exact? m))) "max #2")

(assert (let ((m (min 1 1.5 2 -7))) (and (= m -7) (inexact? m))) "min #1")
(assert (let ((m (min 1 2 -7))) (and (= m -7) (exact? m))) "min #2")

(display "OK\n")
