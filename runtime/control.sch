;; compile-toplevel-phrase returns a thunk.

(define (eval x)
  ((compile-toplevel-phrase x)))

;; @raw-apply@ is a built-in function that uses a special instruction
;; to invoke the function on the arguments with proper tail recursion.

(define apply
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
          (@raw-apply@ fn (construct-apply-args x rest))))))
                     
