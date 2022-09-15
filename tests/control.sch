(display "Testing control.sch\n")

(assert (procedure? assert) "procedure? #1")
(assert-not (procedure? #f) "procedure? #2")

(assert (string=? "bcd" (string-map (lambda (x)
				      (integer->char (+ 1 (char->integer x))))
				    "abc"))
	"string-map")
(assert (equal? '(#\c #\b #\a)
		(let ((l '()))
		  (string-for-each (lambda (x)
				     (set! l (cons x l))) "abc")
		  l))
	"string-for-each")

(assert (eq? (unspecified) (unspecified)) "unspecified #1")
(assert (eq? (unspecified) (if #f #t)) "unspecified #2")

(assert (= 37 (eval '(+ 12 25))) "eval expression")
(eval '(define evald-x 44))
(assert (= 44 evald-x) "eval definition")

(assert (= 0 (apply + '())) "apply #1")
(assert (= 1 (apply + 1 '())) "apply #2")
(assert (= 3 (apply + 1 '(2))) "apply #3")
(assert (= 10 (apply + '(1 2 3 4))) "apply #4")

(assert (= 10 (call-with-values
		  (lambda () (values 1 2 3 4))
		(lambda values
		  (apply + values)))) "call-with-values / values")

(display "OK\n")
