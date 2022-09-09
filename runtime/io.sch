;; -*- indent-tabs-mode: nil; fill-column: 100 -*-

;; R7RS 6.13, Input and output
;; Most operations are implemented as primitives.

;; The key values for the standard port parameters are fixed and defined in core/engine.go.

(define *current-input-port-key* 1)
(define *current-output-port-key* 2)
(define *current-error-port-key* 3)

;; The flag values for the ports are fixed and defined in core/engine.go.
;;
;; Once closed, a port stays closed, but in the presence of goroutines the flag can flip from open
;; to closed at any time.  Thus if input-port-open? or output-port-open? returns #f the result may
;; become invalid before it is inspected due to action by another thread.

(define *input-port-flag* 1)
(define *output-port-flag* 2)
(define *textual-port-flag* 4)
(define *binary-port-flag* 8)
(define *closed-port-flag* 16)

(define (input-port? obj)
  (and (port? obj)
       (not (zero? (bitwise-and (sint:port-flags obj) *input-port-flag*)))))

(define (output-port? obj)
  (and (port? obj)
       (not (zero? (bitwise-and (sint:port-flags obj) *output-port-flag*)))))

(define (textual-port? obj)
  (and (port? obj)
       (not (zero? (bitwise-and (sint:port-flags obj) *textual-port-flag*)))))

(define (binary-port? obj)
  (and (port? obj)
       (not (zero? (bitwise-and (sint:port-flags obj) *binary-port-flag*)))))

(define (input-port-open? p)
  (if (not (input-port? p))
      (error "input-port-open?: Not an input port: " p))
  (zero? (bitwise-and (sint:port-flags obj) *closed-port-flag*)))

(define (output-port-open? p)
  (if (not (output-port? p))
      (error "output-port-open?: Not an output port: " p))
  (zero? (bitwise-and (sint:port-flags obj) *closed-port-flag*)))

(define current-input-port
  (sint:make-parameter-function
   *current-input-port-key* (lambda (p)
                              (if (not (input-port? p))
                                  (error "Cannot set current input port: " p))
                              p)))

(define current-output-port
  (sint:make-parameter-function
   *current-output-port-key* (lambda (p)
                               (if (not (output-port? p))
                                   (error "Cannot set current output port: " p))
                               p)))

(define current-error-port
  (sint:make-parameter-function
   *current-error-port-key* (lambda (p)
                              (if (not (output-port? p))
                                  (error "Cannot set current error port: " p))
                              p)))

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

;; call-with-port always closes the port on escape because it is not possible to re-enter the
;; continuation.

(define (call-with-port port proc)
  (if (not (port? port))
      (error "call-with-port: Not a port: " port))
  (dynamic-wind
      (lambda () #t)
      (lambda () (proc port))
      (lambda () (close-port port))))
