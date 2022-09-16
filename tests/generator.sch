(display "Testing: generator.sch\n")

(let ((g (make-generator (lambda (yield)
			   (for-each yield '(1 2 3 4 5 6)))
			 'end)))
  (letrec ((receive
	       (lambda ()
		 (let ((x (g)))
		   (if (eq? x (unspecified))
		       (list x)
		       (cons x (receive)))))))
    (let ((r (receive)))
      (assert (equal? r (list 1 2 3 4 5 6 'end (unspecified)))
	      "generator"))))
		
(display "OK\n")
