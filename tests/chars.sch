;; Test the character subsystem

(display "Testing: chars.sch\n")

(define (assert x msg)
  (if (not x)
      (error msg)))

(define (assert-not x msg)
  (if x
      (error msg)))

;; Bug #16: Currently the relationals are limited to exactly two arguments

(assert (char? #\a) "char? #1")
(assert-not (char? '()) "char? #2")

(assert (char=? #\a #\a) "char=? #1")
(assert-not (char=? #\a #\b) "char=? #2")

(assert (char<? #\a #\b) "char<? #1")
(assert-not (char<? #\a #\a) "char<? #2")

(assert (char>? #\b #\a) "char>? #1")
(assert-not (char>? #\a #\a) "char>? #2")

(assert (char<=? #\a #\b) "char<=? #1")
(assert (char<=? #\a #\a) "char<=? #2")
(assert-not (char<=? #\b #\a) "char<=? #3")

(assert (char>=? #\b #\a) "char>=? #1")
(assert (char>=? #\a #\a) "char>=? #2")
(assert-not (char>=? #\a #\b) "char>=? #3")

(assert (= (char->integer #\a) 97) "char->integer")

(assert (char=? (integer->char 97) #\a) "integer->char")

(assert (char? (integer->char 1234)) "integer->char + char?")

(display "OK\n")
