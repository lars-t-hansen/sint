;; Demonstrates the use of the built-in goroutine-based generator facility.

;; A tree is a number or a cell (t1 . t2)

(define *next-leaf* 1)

(define (make-tree depth)
  (if (zero? depth)
      (let ((l *next-leaf*))
	(set! *next-leaf* (+ l 1))
	l)
      (cons (make-tree (- depth 1)) (make-tree (- depth 1)))))

;; Yield the leaves of tree t in depth-first order

(define (dfs-tree t yield)
  (if (number? t)
      (yield t)
      (begin
	(dfs-tree (car t) yield)
	(dfs-tree (cdr t) yield))))

(let ((t (make-tree 5)))
  (let ((g (make-generator (lambda (yield)
			     (dfs-tree t yield)))))
    (letrec ((loop
	      (lambda ()
		(let ((x (g)))
		  (if (number? x)
		      (begin
			(display (string-append "Receiving " (number->string x) "\n"))
			(loop)))))))
      (loop))))
