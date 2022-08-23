;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; TODO: This may not terminate
;; TODO: strings and vectors, eventually

(define (equal? a b)
  (or (eqv? a b)
      (and (pair? a)
           (pair? b)
           (equal? (car a) (car b))
           (equal? (cdr a) (cdr b)))))

