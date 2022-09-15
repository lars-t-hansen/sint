;; -*- indent-tabs-mode: nil; fill-column: 100 -*-
;;
;; R7RS 6.7 "Strings"
;;
;; Note that strings in sint are Go strings: immutable byte arrays holding (mostly) UTF8-encoded
;; Unicode.  String lengths and string indices are *byte* lengths and indices.

;; This would be more efficient as a primitive, but it's likely used for very short strings and only
;; infrequently.

(define (string . cs)
  (sint:list->string cs))

;; Nobody cares about the perf of this

(define (make-string k . rest)
  (let ((fill (if (null? rest) #\space (car rest))))
    (sint:list->string (make-list k fill))))

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

(define (list->string cs)
  (if (not (list? cs))
      (error "list->string: not a proper list: " cs))
  (sint:list->string cs))
