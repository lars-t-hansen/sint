;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; (sint:compile-toplevel-phrase x) interprets the datum `x` as top-level source code and returns a
;; thunk that evaluates that code and returns its result.

(define (eval x)
  ((sint:compile-toplevel-phrase x)))

;; (sint:raw-apply fn l count) is a primitive that applies the procedure `fn` to the `count` first
;; values in the list `l` in a properly tail-recursive manner.

;; (define apply
;;   (letrec ((construct-apply-args
;;             (lambda (x rest)
;;               (if (null? rest)
;;                   (if (list? x)
;;                       x
;;                       (error "apply: expected list"))
;;                   (if (null? (cdr rest))
;;                       (if (list? (car rest))
;;                           (cons x (car rest))
;;                           (error "apply: expected list"))
;;                       (let ((rest (reverse rest)))
;;                         (if (list? (car rest))
;;                             (cons x (reverse-append (cdr rest) (car rest)))
;;                             (error "apply: expected list"))))))))
;;     (lambda (fn x . rest)
;;       (if (not (procedure? fn))
;;           (error "apply: expected procedure")
;;           (sint:raw-apply fn (construct-apply-args x rest))))))
                     
;; filter
;; for-each
;; every?
;; some?


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
