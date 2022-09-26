;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7RS 6.14, System interface

(define (features)
  (list 'sint 'sint-0.1))

(define load
  (letrec ((loop
            (lambda (p)
              (let ((item (read p)))
                (if (not (eof-object? item))
                    (begin
                      (call-with-values
                          (lambda () (eval item))
                        (lambda results
                          (if (not (and (= 1 (length results))
                                        (eq? (unspecified) (car results))))
                              (for-each (lambda (x)
                                          (display x)
                                          (newline))
                                        results))))
                      (loop p)))))))
    (lambda (fn)
      (call-with-input-file fn
        loop))))

;; (doc obj) => documentation about obj
;; When we have more debugging information, this will use it opportunistically.

(define (doc obj)
  (cond ((procedure? obj)
         (letrec ((args (lambda (k)
                          (if (zero? k)
                              (if (inexact? k)
                                  'rest
                                  '())
                              (cons (string-append "p" (number->string (exact k)))
                                    (args (- k 1)))))))
           (list 'procedure (procedure-name obj) (list 'lambda (args (procedure-arity obj)) '...))))
        ((and (symbol? obj) (symbol-has-value? obj))
         (list 'symbol obj (symbol-value obj)))
        (else
         (list 'datum obj))))
