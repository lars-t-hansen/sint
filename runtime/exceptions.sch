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
              (if (string? x)
                  x
                  (if (number? x)
                      (number->string x)
                      (if (symbol? x)
                          (symbol->string x)
                          (if (char? x)
                              (string x)
                              (if (eq? x #t)
                                  "#t"
                                  (if (eq? x #f)
                                      "#f"
                                      "#<weird>")))))))))
    (lambda (msg . irritants)
      (sint:throw-string (loop irritants msg)))))

