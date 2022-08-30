;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; Note sint:string-compare checks that both operands are strings, we don't need to check that here.

(define string=?
  (letrec ((loop
            (lambda (a rest)
              (cond ((null? rest)
                     #t)
                    ((= (sint:string-compare a (car rest)) 0)
                     (loop a (cdr rest)))
                    (else
                     #f)))))
    (lambda (a b . rest)
      (loop a (cons b rest)))))

(define string<?
  (letrec ((loop
            (lambda (a rest)
              (cond ((null? rest)
                     #t)
                    ((< (sint:string-compare a (car rest)) 0)
                     (loop a (cdr rest)))
                    (else
                     #f)))))
    (lambda (a b . rest)
      (loop a (cons b rest)))))

(define string<=?
  (letrec ((loop
            (lambda (a rest)
              (cond ((null? rest)
                     #t)
                    ((<= (sint:string-compare a (car rest)) 0)
                     (loop a (cdr rest)))
                    (else
                     #f)))))
    (lambda (a b . rest)
      (loop a (cons b rest)))))

(define string>?
  (letrec ((loop
            (lambda (a rest)
              (cond ((null? rest)
                     #t)
                    ((> (sint:string-compare a (car rest)) 0)
                     (loop a (cdr rest)))
                    (else
                     #f)))))
    (lambda (a b . rest)
      (loop a (cons b rest)))))

(define string>=?
  (letrec ((loop
            (lambda (a rest)
              (cond ((null? rest)
                     #t)
                    ((>= (sint:string-compare a (car rest)) 0)
                     (loop a (cdr rest)))
                    (else
                     #f)))))
    (lambda (a b . rest)
      (loop a (cons b rest)))))

;; FIXME: Use let-values here, it's much cheaper
(define string->list
  (letrec ((loop
            (lambda (first last s i)
              (if (= i (string-length s))
                  first
                  (call-with-values
                      (lambda ()
                        (string-ref s i))
                    (lambda (ch sz) 
                      (let ((c (cons ch '())))
                        (if (null? first)
                            (set! first c)
                            (set-cdr! last c))
                        (loop first c s (+ i sz)))))))))
    (lambda (s . rest)
      (if (not (null? rest))
          (if (not (null? (cdr rest)))
              (set! s (substring s (car rest) (cadr rest)))
              (set! s (substring s (car rest)))))
      (loop '() '() s 0))))

(define (string-copy s . rest)
  (if (not (null? rest))
      (if (not (null? (cdr rest)))
          (substring s (car rest) (cadr rest))
          (substring s (car rest)))
      (substring s 0 (string-length s))))

;; (define (->string x)
;;   (cond ((string? x))
;;         ((number? x) (number->string x))
;;         ((symbol? x) (symbol->string x))
;;         ((procedure? x) "#<procedure>")
;;         ((char? x) (string x))
;;         ((eq? x #t) "#t")
;;         ((eq? x #f) "#f")
;;         ((eq? x '()) "()")
;;         ((eq? x (unspecified)) "#!unspecified")
;;         ((eof-object? x) "#!eof")
;;         ((pair? x) ...)
;;         (else "#<weird>")))
