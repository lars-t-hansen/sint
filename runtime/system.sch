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
