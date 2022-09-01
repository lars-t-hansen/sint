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
                      (eval item)
                      (loop p)))))))
    (lambda (fn)
      (call-with-input-file fn
        (lambda (p)
          (loop p))))))
