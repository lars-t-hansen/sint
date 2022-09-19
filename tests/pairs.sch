(display "Testing: pairs.sch\n")

(assert (null? '()) "null? #1")
(assert-not (null? #f) "null? #2")

(assert (pair? '(a . b)) "pair? #1")
(assert-not (pair? #t) "pair? #2")

(let ((x (cons 'a 'b)))
  (assert (eq? 'a (car x)) "cons/car")
  (assert (eq? 'b (cdr x)) "cons/cdr")
  (set-car! x 'c)
  (set-cdr! x 'd)
  (assert (eq? 'c (car x)) "set-car!")
  (assert (eq? 'd (cdr x)) "set-cdr!"))

(assert (equal? '(a b c) (list 'a 'b 'c)) "list #1")
(assert (null? (list)) "list #2")

(assert (list? (list 'a 'b 'c)) "list? #1")
(assert (list? '()) "list? #2")
(assert-not (list? "hi") "list? #3")
(assert-not (list? (cons 'a 'b)) "list? #4")
;; Issue #36
;; (let ((x (cons '() '())))
;;   (set-cdr! x x)
;;   (assert-not (list? x) "list? #5"))

(assert (equal? (make-list 2) (list (unspecified) (unspecified))) "make-list #1")
(assert (equal? (make-list 3 'a) '(a a a)) "make-list #2")

(assert (equal? (append '(a b c) '(d) '(e)) '(a b c d e)) "append #1")

(assert (equal? (list-tail '(a b c) 1) '(b c)) "list-tail #1")

(assert (eq? 'b (list-ref '(a b c) 1)) "list-ref #1")
(let ((l (list 'a 'b 'c)))
  (list-set! l 1 'd)
  (assert (equal? l '(a d c)) "list-set! #1"))

(assert (equal? (list-copy '(a b c)) '(a b c)) "list-copy #1")

(display "OK\n")
