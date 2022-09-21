;; -*- indent-tabs-mode: nil; fill-column: 100 -*-
;;
;; R7RS 6.5 "Symbols"

(define symbol=?
  (letrec ((check
            (lambda (x)
              (if (not (symbol? x))
                  (error "symbol=?: not a symbol: " x))))
           (loop
            (lambda (sa xs)
              (if (null? xs)
                  #t
                  (begin
                    (check (car xs))
                    (if (string=? sa (symbol->string (car xs)))
                        (loop sa (cdr xs))
                        #f))))))
    (lambda (a b . xs)
      (check a)
      (check b)
      (let ((sa (symbol->string a)))
        (if (string=? sa (symbol->string b))
            (loop sa xs)
            #f)))))

(define (apropos pattern)
  (let ((xs (list-sort! (lambda (a b)
                          (string<? (symbol->string a) (symbol->string b)))
                        (filter-global-variables pattern))))
    (for-each (lambda (x)
                (display x)
                (newline))
              xs)))
