(display "Testing: sort.sch\n")

(assert (list-sorted? < '()) "list-sorted #0")
(assert (list-sorted? < '(1 1 1 1)) "list-sorted? #1")
(assert (list-sorted? < '(1 2 3 4)) "list-sorted? #2")
(assert-not (list-sorted? < '(1 2 3 2)) "list-sorted? #3")
(assert (list-sorted? < '(1)) "list-sorted? #4")

(assert (null? (list-sort! < '())) "list-sort! #1")
(assert (equal? (list-sort! < '(1)) '(1)) "list-sort! #2")
(assert (equal? (list-sort! < (list 3 1 2 4 7 8 1 2 4))
		'(1 1 2 2 3 4 4 7 8))
	"list-sort! #3")

(display "OK\n")
