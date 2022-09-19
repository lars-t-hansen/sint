;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7RS 6.4 "Pairs and lists", and some related non-standard procedures.
;; cons, car, cdr, set-car!, set-cdr!, null? and pair? are primitive.

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

;; FIXME: Issue #26: This is totally insufficient, it needs to deal properly with circular lists.

(define (list? x)
  (cond ((null? x) #t)
        ((pair? x) (list? (cdr x)))
        (else      #f)))

(define make-list
  (letrec ((loop
            (lambda (k fill l)
              (if (<= k 0)
                  l
                  (loop (- k 1) fill (cons fill l))))))
    (lambda (k . rest)
      (loop k (if (null? rest) (unspecified) (car rest)) '()))))

;; FIXME: Issue #33: Needs to check that each of its arguments is a list, though arguably for backward
;; compatibility the last argument should not be checked.

(define append
  (letrec ((loop
            (lambda (result lists)
              (if (null? lists)
                  result
                  (loop (loop2 (reverse (car lists)) result)
                        (cdr lists)))))
           (loop2
            (lambda (rev result)
              (if (null? rev)
                  result
                  (loop2 (cdr rev) (cons (car rev) result))))))
    (lambda lists
      (if (null? lists)
          '()
          (let ((r (reverse lists)))
            (loop (car r) (cdr r)))))))

(define list-tail
  (letrec ((loop
            (lambda (l k)
              (if (<= k 0)
                  l
                  (loop (cdr l) (- k 1))))))
    (lambda (l k)
      (loop l k))))

(define (list-ref l k)
  (car (list-tail l k)))

(define (list-set! l k v)
  (set-car! (list-tail l k) v))

;; FIXME: Issue #34: list-copy is a lot weirder than this

(define (list-copy l)
  (append l '()))

;; TODO: For efficiency, the list? test could be rolled into the loops below so that we don't need
;; to loop across the list twice.

(define memq
  (letrec ((loop
            (lambda (obj list)
              (cond ((null? list) #f)
                    ((eq? obj (car list)) list)
                    (else (loop obj (cdr list)))))))
    (lambda (obj list)
      (if (not (list? list))
          (error "memq: not a list: " list))
      (loop obj list))))

(define memv
  (letrec ((loop
            (lambda (obj list)
              (cond ((null? list) #f)
                    ((eqv? obj (car list)) list)
                    (else (loop obj (cdr list)))))))
    (lambda (obj list)
      (if (not (list? list))
          (error "memv: not a list: " list))
      (loop obj list))))

(define member
  (letrec ((loop
            (lambda (obj list same?)
              (cond ((null? list) #f)
                    ((same? obj (car list)) list)
                    (else (loop obj (cdr list) same?))))))
    (lambda (obj list . rest)
      (if (not (list? list))
          (error "member: not a list: " list))
      (loop obj list (if (null? rest) equal? (car rest))))))

(define assq
  (letrec ((loop
            (lambda (obj alist)
              (cond ((null? alist) #f)
                    ((eq? obj (caar alist)) (car alist))
                    (else (loop obj (cdr alist)))))))
    (lambda (obj alist)
      (if (not (list? alist))
          (error "assq: not a list: " alist))
      (loop obj alist))))

(define assv
  (letrec ((loop
            (lambda (obj alist)
              (cond ((null? alist) #f)
                    ((eqv? obj (caar alist)) (car alist))
                    (else (loop obj (cdr alist)))))))
    (lambda (obj alist)
      (if (not (list? alist))
          (error "assv: not a list: " alist))
      (loop obj alist))))

(define assoc
  (letrec ((loop
            (lambda (obj alist same?)
              (cond ((null? alist) #f)
                    ((same? obj (caar alist)) (car alist))
                    (else (loop obj (cdr alist) same?))))))
    (lambda (obj alist . rest)
      (if (not (list? alist))
          (error "assoc: not a list: " alist))
      (loop obj alist (if (null? rest) equal? (car rest))))))

;; FIXME: Issue #35: Arguably this needs to test for list? while it's computing the length.

(define length
  (letrec ((loop
            (lambda (l k)
              (if (null? l)
                  k
                  (loop (cdr l) (+ k 1))))))
    (lambda (l)
      (loop l 0))))

;; FIXME: Issue #36: Arguably this needs to test for list?

(define reverse
  (letrec ((loop
            (lambda (l r)
              (if (null? l)
                  r
                  (loop (cdr l) (cons (car l) r))))))
    (lambda (l)
      (loop l '()))))

;; FIXME: Issue #36: Arguably this needs to test for list?

(define reverse-append
  (letrec ((loop
            (lambda (xs l)
              (if (null? xs)
                  l
                  (loop (cdr xs) (cons (car xs) l))))))
    (lambda (xs l)
      (loop xs l))))
