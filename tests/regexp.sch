(display "Testing: regexp.sch\n")

(assert (regexp? #/abc/) "regexp? #1")
(assert-not (regexp? "abc") "regexp? #2")

(assert (regexp? (string->regexp "abc")) "string->regexp #1")

(assert (string=? (regexp->string #/abc/) "abc") "regexp->string #1")

(assert (equal? (regexp-find-all #/a./ "abracadabra") '("ab" "ac" "ad" "ab")) "regexp-find-all #1")
(assert (null? (regexp-find-all #/x/ "abracadabra")) "regexp-find-all #2")

(display "OK\n")
