;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; (sint:exact-integer? obj) => bool is a primitive
;; (sint:inexact-float? obj) => bool is a primitive

(define (number? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (complex? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (real? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

;; The floats are technically always rational though this is not
;; a useful fact.
(define (rational? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (integer? obj)
  (sint:exact-integer? obj))
  
(define (real-part z)
  (if (not (number? z))
      (error "real-part: not a number: " z)
      z))

(define (imag-part x)
  (if (not (number? z))
      (error "imag-part: not a number: " z)
      z))

(define (exact? z)
  (if (not (number? z))
      (error "exact?: not a number: " z)
      (sint:exact-integer? z)))

(define (inexact? z)
  (if (not (number? z))
      (error "inexact?: not a number: " z)
      (sint:inexact-float? z)))

(define (exact-integer? z)
  (if (not (number? z))
      (error "exact-integer?: not a number: " z)
      (sint:exact-integer? z)))

;; finite?
;; infinite?
;; nan?

(define (zero? z)
  (= z 0))

(define (positive z)
  (> z 0))

(define (negative? z)
  (< z 0))

(define (odd? n)
  (if (not (exact-integer? n))
      (error "odd?: not an exact integer: " n)
      (not (zero? (remainder n 2)))))

(define (even? n)
  (if (not (exact-integer? n))
      (error "odd?: not an exact integer: " n)
      (zero? (remainder n 2))))

(define max
  (letrec ((loop
            (lambda (isInexact max xs)
              (if (null? xs)
                  (inexact max)
                  (let ((x (car xs)))
                    (if (> x max)
                        (loop (or isInexact (inexact? x)) x (cdr xs))
                        (loop isInexact max (cdr xs))))))))
    (lambda (x . xs)
      (loop (inexact? x) x xs))))

(define min
  (letrec ((loop
            (lambda (isInexact min xs)
              (if (null? xs)
                  (inexact min)
                  (let ((x (car xs)))
                    (if (< x min)
                        (loop (or isInexact (inexact? x)) x (cdr xs))
                        (loop isInexact min (cdr xs))))))))
    (lambda (x . xs)
      (loop (inexact? x) x xs))))
