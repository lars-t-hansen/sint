;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; (sint:compile-toplevel-phrase x) interprets the datum `x` as top-level source code and returns a
;; thunk that evaluates that code and returns its result.

(define (eval x)
  ((sint:compile-toplevel-phrase x)))

;; (sint:apply fn l) is a primitive that applies the procedure `fn` to the values
;; in the proper list `l` in a properly tail-recursive manner.

(define apply
  (letrec ((build-apply-args
            (lambda (fst rest)
              (cond ((null? rest)
                     (if (not (list? fst))
                         (error "apply: expected list: " fst))
                     fst)
                    ((null? (cdr rest))
                     (if (not (list? (car rest)))
                         (error "apply: expected list: " (car rest)))
                     (cons fst (car rest)))
                    (else
                     (let ((rest (reverse rest)))
                       (if (not (list? (car rest)))
                           (error "apply: expected list: " (car rest)))
                       (cons x (reverse-append (cdr rest) (car rest)))))))))
    (lambda (fn x . rest)
      (if (not (procedure? fn))
          (error "apply: expected procedure"))
      (sint:apply fn (build-apply-args x rest)))))

;; (sint:receive-values thunk) is a primitive that invokes the procedure `thunk` and then returns a
;; proper list of the values it returns.

(define (call-with-values thunk receiver)
  (if (not (procedure? thunk))
      (error "call-with-values: expected procedure: " thunk))
  (if (not (procedure? receiver))
      (error "call-with-values: expected procedure: " receiver))
  (sint:apply receiver (sint:receive-values thunk)))

;; TODO:
;; filter
;; for-each
;; every?
;; some?

;; TODO: what's the appropriate termination condition for multi-argument map?
;; Here it's "first list", but spec may have "shortest list".
;;
;; TODO: detect non-list arguments.

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
