(display "Testing: strings.sch\n")

(assert (string? "abcd") "string? #1")
(assert-not (string? 'abcd) "string? #2")

;; Remember, strings are byte slices holding UTF8 encodings of unicode
;; code points, and Norwegian letters such as "ø" occupy more than one
;; byte.

(assert (= 3 (string-length "abc")) "string-length #1")
(assert (= 4 (string-length "abø")) "string-length #2")

(assert (char=? #\b (string-ref "abc" 1)) "string-ref #1")
(assert (char=? #\ø (string-ref "abø" 2)) "string-ref #2")
(assert (= 2 (call-with-values
		 (lambda ()
		   (string-ref "abø" 2))
	       (lambda (c len)
		 len)))
	"string-ref #3")

(assert (string=? (string-append "ab" "cd" "ef") "abcdef") "string-append")

(assert (string=? (substring "abcd" 1 3) "bc") "substring #1")
(assert (string=? (substring "abød" 1 4) "bø") "substring #2")

(assert (string=? (string #\a #\b #\ø #\d) "abød") "string")

(assert (string=? (make-string 5 #\a) "aaaaa") "make-string #1")
(assert (string=? (make-string 5) "     ") "make-string #2")

(assert (string=? "ab" "ab" "ab") "string=? #1")
(assert-not (string=? "ab" "ab" "abc") "string=? #2")

(assert (string<? "ab" "ac" "ad") "string<? #1")
(assert-not (string<? "ab" "ac" "ac") "string<? #2")

(assert (string<=? "ab" "ac" "ac") "string<=? #1")
(assert-not (string<=? "ab" "ac" "ab") "string<=? #2")

(assert (string>? "ad" "ac" "ab") "string>? #1")
(assert-not (string>? "ac" "ab" "ab") "string>? #2")

(assert (string>=? "ad" "ac" "ab" "ab") "string>=? #1")
(assert-not (string>=? "ac" "ab" "ab" "ac") "string>=? #2")

(assert (equal? (string->list "abc") '(#\a #\b #\c)) "string->list #1")
(assert (equal? (string->list "abc" 1) '(#\b #\c)) "string->list #2")
(assert (equal? (string->list "abcd" 1 3) '(#\b #\c)) "string->list #3")

(assert (string=? "abc" (list->string '(#\a #\b #\c))) "list->string")

(assert (string=? (string-copy "abcd") "abcd") "string-copy #1")
(assert (string=? (string-copy "abcd" 1) "bcd") "string-copy #2")
(assert (string=? (string-copy "abcd" 1 3) "bc") "string-copy #3")

(display "OK\n")
