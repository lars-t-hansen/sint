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

(assert (equal? '(1 2 3) (map (lambda (x) (+ x 1)) '(0 1 2))) "map #1")
(assert (equal? '(12 23 34) (map (lambda (a b)
				   (+ (* a 10) b))
				 '(1 2 3)
				 '(2 3 4 5)))
	"map #2")

(assert (= 10 (let ((k 0))
		(for-each (lambda (x) (set! k (+ k x))) '(1 2 3 4))
		k))
	"for-each #1")

(assert (every? (lambda (x) (> x 0)) '(1 2 3 5)) "every? #1")
(assert-not (every? (lambda (x) (> x 0)) '(1 2 0 5)) "every? #2")

(assert (some? zero? '(1 2 0 5)) "some? #1")
(assert-not (some? zero? '(1 2 3 5)) "some? #2")

(assert (equal? '(0 0) (filter zero? '(1 2 3 0 1 2 0 1))) "filter #1")

;; TODO: make-parameter
;; TODO: call-with-current-continuation
;; TODO: dynamic-wind

(display "OK\n")