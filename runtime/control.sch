;; -*- fill-column: 100 -*-

;; (sint:compile-toplevel-phrase x) interprets the datum `x` as top-level source code and returns a
;; thunk that evaluates that code and returns its result.

(define (eval x)
  ((sint:compile-toplevel-phrase x)))

;; (sint:raw-apply fn l count) is a primitive that applies the procedure `fn` to the `count` first
;; values in the list `l` in a properly tail-recursive manner.

  (letrec ((construct-apply-args
            (lambda (x rest)
              (if (null? rest)
                  (if (list? x)
                      x
                      (error "apply: expected list"))
                  (if (null? (cdr rest))
                      (if (list? (car rest))
                          (cons x (car rest))
                          (error "apply: expected list"))
                      (let ((rest (reverse rest)))
                        (if (list? (car rest))
                            (cons x (reverse-append (cdr rest) (car rest)))
                            (error "apply: expected list"))))))))
    (lambda (fn x . rest)
      (if (not (procedure? fn))
          (error "apply: expected procedure")
          (sint:raw-apply fn (construct-apply-args x rest))))))
                     
