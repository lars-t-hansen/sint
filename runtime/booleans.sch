;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

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
              (if (null? xs)
                  #t
                  (if (eq? a (car xs))
                      (loop a (cdr xs))
                      (begin
                        (check (car xs))
                        #f))))))
    (lambda (a b . xs)
      (check a)
      (check b)
      (if (eq? a b)
          (loop b xs)
          #f))))
