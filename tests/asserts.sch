(define (assert x msg)
  (if (not x)
      (error msg)))

(define (assert-not x msg)
  (if x
      (error msg)))

