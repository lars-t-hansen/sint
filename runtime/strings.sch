;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; Probably most of these should be written in Go.

;; TODO: This is just plain slow
(define (string->vector s)
  (list->vector (string-map (lambda (c) c) s)))

;; TODO: This conses a lot
(define (vector->string v)
  (list->string (vector->list v)))

;; TODO: This conses a lot
(define (string-for-each fn . args)
  (apply string-map fn args)
  (unspecified))
