;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; note map, for-each, et al are in control.sch

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

;; FIXME: Totally insufficient.

(define (list? x)
  (cond ((null? x) #t)
        ((pair? x) (list? (cdr x)))
        (else      #f)))

;; list?
;; make-list
;; append
;; list-tail
;; list-ref
;; list-set!
;; assq
;; assv
;; assoc
;; list-copy

(define (memq obj list)
  (cond ((null? list) #f)
        ((eq? obj (car list)) list)
        (else (memq obj (cdr list)))))

(define (memv obj list)
  (cond ((null? list) #f)
        ((eqv? obj (car list)) list)
        (else (memv obj (cdr list)))))

(define member
  (letrec ((loop
            (lambda (obj list same?)
              (cond ((null? list) #f)
                    ((same? obj (car list)) list)
                    (else (loop obj (cdr list) same?))))))
    (lambda (obj list . rest)
      (loop obj list (if (null? rest) equal? (car rest))))))

(define length
  (letrec ((loop
            (lambda (l k)
              (if (null? l)
                  k
                  (loop (cdr l) (+ k 1))))))
    (lambda (l)
      (loop l 0))))

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
