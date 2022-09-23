(define (assert x msg)
  (if (not x)
      (error msg))
;  (display msg) (newline)
  )

(define (assert-not x msg)
  (if x
      (error msg))
;  (display msg) (newline)
  )

