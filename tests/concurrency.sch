(display "Testing: concurrency.sch\n")

(define (send c n)
  (if (not (zero? n))
      (begin
	(channel-send c n)
	(send c (- n 1)))))

(define (receive c n)
  (if (zero? n)
      '()
      (let ((x (channel-receive c)))
	(cons x (receive c (- n 1))))))

(define *chan-size* 5)
(define chan (make-channel *chan-size*))

(assert (channel? chan) "channel? #1")
(assert-not (channel? "hi there") "channel? #2")

(assert (= (channel-length chan) 0) "channel-length #1")
(assert (= (channel-capacity chan) *chan-size*) "channel-capacity")

(define *item* "hello")
(channel-send chan *item*)
(assert (= (channel-length chan) 1) "channel-length #2")
(let ((x (channel-receive chan)))
  (assert (= (channel-length chan) 0) "channel-length #3")
  (assert (eq? x *item*) "channel-send / channel-receive"))

(go (send chan 10))
(assert (equal? '(10 9 8 7 6 5 4 3 2 1) (receive chan 10)) "go + channel")

(close-channel chan)
(assert (eq? (unspecified) (channel-receive chan)) "channel-receive on closed channel")

(assert (exact-integer? (goroutine-id)) "goroutine-id")

(display "OK\n")
