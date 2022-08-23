(define (caar x) (car (car x)))
(define (cadr x) (car (cdr x)))
(define (cdar x) (cdr (car x)))
(define (cddr x) (cdr (cdr x)))

(define (caaar x) (car (car (car x))))
(define (caadr x) (car (car (cdr x))))
(define (cadar x) (car (cdr (car x))))
(define (caddr x) (car (cdr (cdr x))))
(define (cdaar x) (cdr (car (car x))))
(define (cdadr x) (cdr (car (cdr x))))
(define (cddar x) (cdr (cdr (car x))))
(define (cdddr x) (cdr (cdr (cdr x))))

(define (caaaar x) (car (car (car (car x)))))
(define (caaadr x) (car (car (car (cdr x)))))
(define (caadar x) (car (car (cdr (car x)))))
(define (caaddr x) (car (car (cdr (cdr x)))))
(define (cadaar x) (car (cdr (car (car x)))))
(define (cadadr x) (car (cdr (car (cdr x)))))
(define (caddar x) (car (cdr (cdr (car x)))))
(define (cadddr x) (car (cdr (cdr (cdr x)))))
(define (cdaaar x) (cdr (car (car (car x)))))
(define (cdaadr x) (cdr (car (car (cdr x)))))
(define (cdadar x) (cdr (car (cdr (car x)))))
(define (cdaddr x) (cdr (car (cdr (cdr x)))))
(define (cddaar x) (cdr (cdr (car (car x)))))
(define (cddadr x) (cdr (cdr (car (cdr x)))))
(define (cdddar x) (cdr (cdr (cdr (car x)))))
(define (cddddr x) (cdr (cdr (cdr (cdr x)))))

(define (list . xs) xs)

;;(define (list? xs) ...)
;; list-ref
;; filter
;; for-each
;; every?
;; some?
;; append

(define length
  (letrec ((loop
	    (lambda (l k)
	      (if (null? l)
		  k
		  (loop (cdr l) (+ k 1))))))
    (lambda (l)
      (loop l 0))))

(define map
  (letrec ((map1
            (lambda (fn l0)
              (if (null? l0)
                  '()
                  (cons (fn (car l0))
			(map1 fn (cdr l0))))))
           (map2
            (lambda (fn l0 l1)
              (if (null? l0)
                  '()
                  (cons (fn (car l0) (car l1))
                        (map2 fn (cdr l0) (cdr l1))))))
	   (mapn
	    (lambda (fn ls)
	      (if (null? (car ls))
		  '()
		  (cons (apply fn (map1 car ls))
			(mapn fn (map1 cdr ls)))))))
    (lambda (fn l0 . rest)
      (if (null? rest)
          (map1 fn l0)
          (if (null? (cdr rest))
              (map2 fn l0 (car rest))
	      (mapn fn (cons l0 rest)))))))

(define reverse
  (letrec ((loop
	    (lambda (l r)
	      (if (null? l)
		  r
		  (loop (cdr l) (cons (car l) r))))))
    (lambda (l)
      (loop l '()))))

(define reverse-append
  (letrec ((loop
	    (lambda (xs l)
	      (if (null? xs)
		  l
		  (loop (cdr xs) (cons (car xs) l))))))
    (lambda (xs l)
      (loop xs l))))
