;; Test the I/O subsystem in a superficial way.

(display "Testing: io.sch\n")

(define (assert x msg)
  (if (not x)
      (error msg)))

;; Test the port predicates

(assert (input-port? (current-input-port)) "current-input-port is an input port")
(assert (textual-port? (current-input-port)) "current-input-port is a textual port")
(assert (not (binary-port? (current-input-port))) "current-input-port is not a binary port")
(assert (input-port-open? (current-input-port)) "current-input-port is open")

(assert (output-port? (current-output-port)) "current-output-port is an output port")
(assert (textual-port? (current-output-port)) "current-output-port is a textual port")
(assert (not (binary-port? (current-output-port))) "current-output-port is not a binary port")
(assert (output-port-open? (current-output-port)) "current-output-port is open")

(assert (output-port? (current-error-port)) "current-error-port is an output port")
(assert (textual-port? (current-error-port)) "current-error-port is a textual port")
(assert (not (binary-port? (current-error-port))) "current-error-port is not a binary port")
(assert (output-port-open? (current-error-port)) "current-error-port is open")

;; Test call-with-input-file, eof-object?, read-char
;; foo.txt contains exactly "foo<newline>"

(call-with-input-file "tests/foo.txt"
  (lambda (p)
    (assert (char=? (read-char p) #\f) "read-char 1")
    (assert (char=? (read-char p) #\o) "read-char 2")
    (assert (char=? (read-char p) #\o) "read-char 3")
    (assert (char=? (read-char p) #\newline) "read-char 4")
    (assert (eof-object? (read-char p)) "eof 1")
    (assert (eof-object? (read-char p)) "eof 2")))

(call-with-input-file "tests/foo.txt"
  (lambda (p)
    (let ((x (read p)))
      (assert (eq? x 'foo) "read foo")
      (assert (eof-object? (read p)) "read foo eof"))))
      
(display "OK\n")

