;; The traditional doubly-recursive fibonnaci function.

(define (fib n)
  (if (< n 2)
      n
      (+ (fib (- n 1)) (fib (- n 2)))))

(let* ((then (current-jiffy))
       (r (fib 30))
       (now (current-jiffy)))
  (values r (/ (- now then) (jiffies-per-second))))

