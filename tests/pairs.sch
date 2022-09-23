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
(let ((x (cons '() '())))
  (set-cdr! x x)
  (assert-not (list? x) "list? #5"))
(let ((x (list 1 2 3)))
  (set-cdr! (list-tail x 2) x)
  (assert-not (list? x) "list? #6"))
(let ((x (list 1 2 3 4)))
  (set-cdr! (list-tail x 3) x)
  (assert-not (list? x) "list? #7"))

(assert (equal? (make-list 2) (list (unspecified) (unspecified))) "make-list #1")
(assert (equal? (make-list 3 'a) '(a a a)) "make-list #2")

(assert (equal? (append '(a b c) '(d) '(e)) '(a b c d e)) "append #1")

(assert (equal? (list-tail '(a b c) 1) '(b c)) "list-tail #1")

(assert (eq? 'b (list-ref '(a b c) 1)) "list-ref #1")
(let ((l (list 'a 'b 'c)))
  (list-set! l 1 'd)
  (assert (equal? l '(a d c)) "list-set! #1"))

(assert (equal? (list-copy '(a b c)) '(a b c)) "list-copy #1")
(assert (equal? (list-copy 'a) 'a) "list-copy #2")
(assert (equal? (list-copy '(a b . c)) '(a b . c)) "list-copy #3")

(assert (eq? 'a (car (memq 'a '(b c a b)))) "memq #1")
(assert-not (memq 'd '(b c a b)) "memq #2")

(assert (eqv? 1 (car (memv 1 '(2 3 1 4)))) "memv #1")
(assert-not (memv 7 '(2 3 1 4)) "memv #2")

(assert (equal? '(a) (car (member '(a) '((c) (d) (a) (b))))) "member #1")
(assert-not (member '(e) '((c) (d) (a) (b))) "member #2")
(define (same? x y)
  (equal? x (reverse y)))
(assert (same? '(a b) (car (member '(a b) '((x c) (y d) (b a) (c b)) same?))) "member #3")
(assert-not (member '(a c) '((x c) (y d) (b a) (c b)) same?) "member #4")

(assert (equal? '(a b) (assq 'a '((a b) (c x) (b a) (d b)))) "assq #1")
(assert-not (assq 'a '((y b) (c x) (b a) (d b))) "assq #2")

(assert (equal? '(1 2) (assv '1 '((1 2) (2 3) (3 4) (4 5)))) "assv #1")
(assert-not (assv '5 '((1 2) (2 3) (3 4) (4 5))) "assv #2")

(assert (equal? '((a b) 3) (assoc '(a b) '(((a b) 3) ((c x) 4) ((b a) 5) ((d b) 7)))) "assoc #1")
(assert-not (assoc '(b b) '(((a b) 3) ((c x) 4) ((b a) 5) ((d b) 7))) "assoc #2")

(assert (= 0 (length '())) "length #1")
(assert (= 5 (length '(a b c d e))) "length #2")

(assert (equal? '(e d c b a) (reverse '(a b c d e))) "reverse #1")

(assert (equal? '(a b c d e) (reverse-append '(c b a) '(d e))) "reverse-append #1")

(display "OK\n")
