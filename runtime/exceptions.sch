;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

(define error
  (letrec ((loop
            (lambda (irritants s)
              (if (null? irritants)
                  s
                  (loop (cdr irritants)
                        (string-append s " " (fmt (car irritants)))))))
           (fmt
            ;; Dumb pretty-printer
            ;; FIXME: We need something much better
            (lambda (x)
              (cond ((string? x))
                    ((number? x) (number->string x))
                    ((symbol? x) (symbol->string x))
                    ((char? x) (string x))
                    ((eq? x #t) "#t")
                    ((eq? x #f) "#f")
                    (else "#<weird>")))))
    (lambda (msg . irritants)
      (sint:throw-string (loop irritants msg)))))

