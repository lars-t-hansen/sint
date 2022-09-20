;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7RS 6.2 - Numbers

(define (number? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (complex? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (real? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (rational? obj)
  (or (sint:exact-integer? obj) (sint:inexact-float? obj)))

(define (integer? obj)
  (sint:exact-integer? obj))
  
(define (real-part z)
  (if (not (number? z))
      (error "real-part: not a number: " z))
  z)

(define (imag-part z)
  (if (not (number? z))
      (error "imag-part: not a number: " z))
  0)

(define (exact? z)
  (if (not (number? z))
      (error "exact?: not a number: " z))
  (sint:exact-integer? z))

(define (inexact? z)
  (if (not (number? z))
      (error "inexact?: not a number: " z))
  (sint:inexact-float? z))

(define (exact-integer? z)
  (if (not (number? z))
      (error "exact-integer?: not a number: " z))
  (sint:exact-integer? z))

(define exact->inexact inexact)
(define inexact->exact exact)

(define (square z)
  (* z z))

(define (nan? z)
  (if (not (number? z))
      (error "nan?: not a number: " z))
  #f)

(define (zero? z)
  (= z 0))

(define (positive? z)
  (> z 0))

(define (negative? z)
  (< z 0))

(define (odd? n)
  (if (not (exact-integer? n))
      (error "odd?: not an exact integer: " n))
  (not (zero? (remainder n 2))))

(define (even? n)
  (if (not (exact-integer? n))
      (error "even?: not an exact integer: " n))
  (zero? (remainder n 2)))

(define max
  (letrec ((loop
            (lambda (isInexact max xs)
              (if (null? xs)
                  (if isInexact
                      (inexact max)
                      max)
                  (let ((x (car xs)))
                    (let ((isInexact (or isInexact (inexact? x))))
                      (if (> x max)
                          (loop isInexact x (cdr xs))
                          (loop isInexact max (cdr xs)))))))))
    (lambda (x . xs)
      (loop (inexact? x) x xs))))

(define min
  (letrec ((loop
            (lambda (isInexact min xs)
              (if (null? xs)
                  (if isInexact
                      (inexact min)
                      min)
                  (let ((x (car xs)))
                    (let ((isInexact (or isInexact (inexact? x))))
                      (if (< x min)
                          (loop isInexact x (cdr xs))
                          (loop isInexact min (cdr xs)))))))))
    (lambda (x . xs)
      (loop (inexact? x) x xs))))
  
;; k^1 = k
;; k^n = k^n-1 * k^1 if n odd
;; k^n = (k^(n/2))^2 if n even

(define expt
  (letrec ((expt
            (lambda (x y)
              (cond ((= y 1)
                     x)
                    ((even? y)
                     (square (expt x (quotient y 2))))
                    (else
                     (* (expt x (- y 1)) x))))))
    (lambda (x y)
      (if (not (and (exact-integer? y) (>= y 0)))
          (error "expt: Exponent must be exact nonnegative integer" y))
      (if (not (rational? x))
          (error "expt: Base must be rational number" x))
      (cond ((= x 0)
             (if (= y 0) 1 0))
            ((= y 0)
             1)
            (else
             (expt x y))))))
