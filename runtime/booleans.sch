;; -*- indent-tabs-mode: nil; fill-column: 100 -*-
;;
;; R7RS 6.3 "Booleans"

(define (boolean? x)
  (or (eq? x #t) (eq? x #f)))

(define (not x)
  (if x #f #t))

(define boolean=?
  (letrec ((check
            (lambda (x)
              (if (not (boolean? x))
                  (error "boolean=?: not a boolean: " x))))
           (loop
            (lambda (a xs)
              (cond ((null? xs)
                     #t)
                    ((eq? a (car xs))
                     (loop a (cdr xs)))
                    (else
                     (check (car xs))
                     #f)))))
    (lambda (a b . xs)
      (check a)
      (check b)
      (and (eq? a b) (loop a xs)))))
