;; Test the I/O subsystem in a superficial way.

(display "Testing: io.sch\n")

;; Test the port predicates

(assert (input-port? (current-input-port)) "current-input-port is an input port")
(assert (textual-port? (current-input-port)) "current-input-port is a textual port")
(assert (not (binary-port? (current-input-port))) "current-input-port is not a binary port")
(assert (input-port-open? (current-input-port)) "current-input-port is open")
(assert (port? (current-input-port)) "current-input-port is a port")

(assert (output-port? (current-output-port)) "current-output-port is an output port")
(assert (textual-port? (current-output-port)) "current-output-port is a textual port")
(assert (not (binary-port? (current-output-port))) "current-output-port is not a binary port")
(assert (output-port-open? (current-output-port)) "current-output-port is open")
(assert (port? (current-output-port)) "current-output-port is a port")

(assert (output-port? (current-error-port)) "current-error-port is an output port")
(assert (textual-port? (current-error-port)) "current-error-port is a textual port")
(assert (not (binary-port? (current-error-port))) "current-error-port is not a binary port")
(assert (output-port-open? (current-error-port)) "current-error-port is open")
(assert (port? (current-error-port)) "current-error-port is a port")

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

(call-with-output-file "tests/out.txt"
  (lambda (p)
    (write-char #\b p)
    (write-char #\a p)
    (write-char #\z p)
    (write-char #\newline p)))
(call-with-input-file "tests/out.txt"
  (lambda (p)
    (assert (char=? (read-char p) #\b) "write-char 1")
    (assert (char=? (read-char p) #\a) "write-char 2")
    (assert (char=? (read-char p) #\z) "write-char 3")
    (assert (char=? (read-char p) #\newline) "write-char 4")
    (assert (eof-object? (read-char p)) "write-char 5")))

(call-with-input-file "tests/foo.txt"
  (lambda (p)
    (assert (char=? (peek-char p) #\f) "peek-char 1")
    (assert (char=? (peek-char p) #\f) "peek-char 2")
    (assert (char=? (read-char p) #\f) "peek-char 3")
    (assert (char=? (read-char p) #\o) "peek-char 4")
    (assert (char=? (read-char p) #\o) "peek-char 5")
    (assert (char=? (read-char p) #\newline) "peek-char 6")
    (assert (eof-object? (peek-char p)) "peek-char 7")
    (assert (eof-object? (peek-char p)) "peek-char 8")
    (assert (eof-object? (read-char p)) "peek-char 9")))

(call-with-input-file "tests/foo.txt"
  (lambda (p)
    (let ((x (read p)))
      (assert (eq? x 'foo) "read foo")
      (assert (eof-object? (read p)) "read foo eof"))))
      
(let ((p (open-input-file "tests/foo.txt")))
  (call-with-port p
    (lambda (p)
      (assert (eq? (read p) 'foo) "call-with-port / close-port #1")))
  (assert-not (input-port-open? p) "call-with-port / close-port #2"))

(let ((p (open-output-file "tests/out.txt")))
  (call-with-port p
    (lambda (p)
      (write 'hi p)))
  (assert-not (output-port-open? p) "call-with-port / close-port #3"))

(display "OK\n")

