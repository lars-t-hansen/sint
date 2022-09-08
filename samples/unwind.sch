(define (unw-test n k)
  (if (zero? n)
      (k "hi")
      (begin
	(writeln (string-append "before " (number->string n)))
	(dynamic-wind
	    (lambda () (set! *inc* (+ *inc* 1)))
	    (lambda () (unw-test (- n 1) k))
	    (lambda () (set! *inc* (- *inc* 1))))
	(writeln (string-append "after " (number->string n))))))

(define *inc* 0)

(call-with-current-continuation
 (lambda (k)
   (unw-test 20 k)))

*inc*
