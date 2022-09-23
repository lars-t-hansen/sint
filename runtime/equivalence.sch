;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7R7 6.1 "Equivalence predicates"
;;
;; TODO: vectors and bytevectors, eventually

;; TODO: This may not terminate

(define (equal? a b)
  (or (eqv? a b)
      (and (pair? a)
           (pair? b)
           (equal? (car a) (car b))
           (equal? (cdr a) (cdr b)))
      (and (string? a)
           (string? b)
           (string=? a b))))


