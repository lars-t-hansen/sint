(define (boolean? x)
  (or (eq? x #t) (eq? x #f)))

(define (not x)
  (if x #f #t))