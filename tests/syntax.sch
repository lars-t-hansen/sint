(display "Testing: syntax.sch\n")

(let ((i 37)
      (j 42))
  (let* ((i (+ i 1))
	 (j (+ i 1 j)))
    (assert (= i 38) "let* #1")
    (assert (= j 81) "let* #2")))

(display "OK\n")
