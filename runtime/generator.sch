;; -*- fill-column: 100 -*-

;; Simple generator abstraction, based on goroutines.

;; TODO: We might want another one here, make-serialized-generator, which uses some kind of latch to
;; ensure that there are no concurrency issues.  Only a call to the generator would trigger the
;; generation of the next item.  We could use a second channel for this signal.

(define (make-generator p . rest)

  "make-generator takes a procedure `p` of one argument, and optionally a value `end`, and returns a
thunk, the generator.  The generator is invoked to obtain the values yielded by `p`.  `p` will be
called once with a a function of one argument that is used by `p` to yield its values.  Once `p`
returns, the value `end` is yielded by the generator once, followed by a stream of #!unspecified.

`p` is run on a concurrent thread, and care should be taken by both `p` and the consumer when
updating shared state.  The communication channel is unbuffered, so `p` and the consumer run
somewhat in lockstep, but they are not synchronized: `p` is working on the next item while the
consumer is processing the previous one.
"

  (let ((chan (make-channel))
	(end  (if (not (null? rest)) (car rest) (unspecified))))
    (go ((lambda ()
	   (p (lambda (x)
		(channel-send chan x)))
	   (channel-send chan end)
	   (close-channel chan))))
    (lambda ()
      (channel-receive chan))))

