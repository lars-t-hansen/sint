;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; TODO: High-value control operations
;;
;; call/cc but only one-shot, same-goroutine, upward
;; dynamic-wind, it is used by the I/O system

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

;; A parameter function takes zero or one arguments.  When called with zero, it returns the current
;; value.  When called with one, it sets the value to its argument after applying the conversion
;; function.
;;
;; TODO: We should mark the parameter object as a parameter, somehow.  This could be done by keeping
;; it in a weak table, or by setting a flag on it.  Parameterize would then check that flag.

(define (make-parameter init . rest)
  (let ((conv (if (null? rest) (lambda (x) x) (car rest)))
        (key  (sint:new-tls-key)))
    (sint:write-tls-value key (conv init))
    (lambda args
      (cond ((null? args)
             (sint:read-tls-value key))
            ((null? cdr rest)
             (sint:write-tls-value key (conv (car rest))))
            (else
             (error "Invalid call to parameter function"))))))

(define (dynamic-wind before during after)
  (before)
  (sint:unwind-protect during after))

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
              (if (or (null? l0) (null? l1))
                  '()
                  (cons (fn (car l0) (car l1))
                        (map2 fn (cdr l0) (cdr l1))))))
           (mapn
            (lambda (fn ls)
              (if (some? null? ls)
                  '()
                  (cons (apply fn (map1 car ls))
                        (mapn fn (map1 cdr ls)))))))
    (lambda (fn l0 . rest)
      (cond ((null? rest)
             (map1 fn l0))
            ((null? (cdr rest))
             (map2 fn l0 (car rest)))
            (else
              (mapn fn (cons l0 rest)))))))

;; TODO: detect non-list arguments.

(define for-each
  (letrec ((each1
            (lambda (fn l0)
              (if (null? l0)
                  (unspecified)
                  (begin
                    (fn (car l0))
                    (each1 fn (cdr l0))))))
           (each2
            (lambda (fn l0 l1)
              (if (or (null? l0) (null? l1))
                  (unspecified)
                  (begin
                    (fn (car l0) (car l1))
                    (each2 fn (cdr l0) (cdr l1))))))
           (eachn
            (lambda (fn ls)
              (if (some? null? ls)
                  (unspecified)
                  (begin
                    (apply fn (each1 car ls))
                    (eachn fn (each1 cdr ls)))))))
    (lambda (fn l0 . rest)
      (cond ((null? rest)
             (each1 fn l0))
            ((null? (cdr rest))
             (each2 fn l0 (car rest)))
            (else
             (eachn fn (cons l0 rest)))))))

;; TODO: Multi-list version?
;; TODO: detect non-list argument

(define every?
  (letrec ((loop
            (lambda (p l)
              (cond ((null? l) #t)
                    ((not (p (car l))) #f)
                    (else (loop p (cdr l)))))))
    (lambda (p l)
      (loop p l))))

;; TODO: Multi-list version?
;; TODO: detect non-list argument

(define some?
  (letrec ((loop
            (lambda (p l)
              (cond ((null? l) #f)
                    ((p (car l)) #t)
                    (else (loop p (cdr l)))))))
    (lambda (p l)
      (loop p l))))

;; TODO: Multi-list version?
;; TODO: detect non-list argument

(define filter
  (letrec ((loop
            (lambda (p l)
              (cond ((null? l) '())
                    ((p (car l))
                     (cons (car l) (loop p (cdr l))))
                    (else
                     (loop p (cdr l)))))))
    (lambda (p l)
      (loop p l))))
