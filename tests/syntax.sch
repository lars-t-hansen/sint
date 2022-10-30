(display "Testing: syntax.sch\n")

(let ((i 37)
      (j 42))
  (let* ((i (+ i 1))
	 (j (+ i 1 j))
	 (i (+ i 1)))
    (assert (= i 39) "let* #1")
    (assert (= j 81) "let* #2")))

(display "OK\n")
