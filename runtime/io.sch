;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7RS 6.13, Input and output
;; Most operations are implemented as primitives.

(define (call-with-input-file fn proc)
  (let ((port (open-input-file fn)))
    (dynamic-wind
        (lambda () #t)
        (lambda () (proc port))
        (lambda () (close-input-port port)))))

(define (call-with-output-file fn proc)
  (let ((port (open-output-file fn)))
    (dynamic-wind
        (lambda () #t)
        (lambda () (proc port))
        (lambda () (close-output-port port)))))
