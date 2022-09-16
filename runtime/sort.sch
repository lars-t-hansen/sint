;; -*- fill-column: 100 -*-

;; list-sort! - in-place list sort.
;;
;; An in-place merge sort uses O(lg(n)) stack space and O(1) extra list space, and runs in
;; guaranteed O(n*lg(n)) time.
;;
;; Sorting (a b c d e f g), it creates in order (of course with necessary rearrangements):
;;
;;  (a) (b) (c d e f g)
;;  (a b) (c d e f g)
;;  (a b) (c) (d) (e f g)
;;  (a b) (c d) (e f g)
;;  (a b c d) (e f g)
;;  (a b c d) (e) (f) (g)
;;  (a b c d) (e f) (g)
;;  (a b c d) (e f g)
;;  (a b c d e f g)
;;
;; using the stack to hold the intermediate lists, never more than lg(n) of them at the same time.
;;
;; TODO: The inner lambdas in build-sorted! must be removed.
;;
;; TODO: Stability of this?  We should be able to have a stable sort, but I don't think this is.

(define list-sort!
  (letrec ((list-merge!
	    (lambda (<? xs ys tl)
	      ;; Given sorted lists xs and ys and a binary ordering predicate <?, and a cell tl,
	      ;; merge the xs and ys destructively into the tail of tl.
	      (cond ((null? xs)
		     (set-cdr! tl ys))
		    ((null? ys)
		     (set-cdr! tl xs))
		    ((<? (car xs) (car ys))
		     (let ((tail-xs (cdr xs)))
		       (set-cdr! xs '()) ; TODO: Technically redundant?
		       (set-cdr! tl xs)
		       (list-merge! <? tail-xs ys xs)))
		    (else
		     (let ((tail-ys (cdr ys)))
		       (set-cdr! ys '()) ; TODO: Technically redundant?
		       (set-cdr! tl ys)
		       (list-merge! <? xs tail-ys ys))))))
	   
	   (build-sorted!
	    (lambda (n <? xs tmp)
	      ;; Given a depth n, a predice <?, a list xs, and a pair tmp, return two values: a
	      ;; sorted list of the first 2^n elements of xs (or all of xs if it is shorter than
	      ;; that), and the rest of xs.
	      (cond ((null? xs)
		     (values '() '()))
		    ((null? (cdr xs))
		     (values xs '()))
		    ((= n 1)
		     (let ((rest-xs (cdr xs)))
		       (set-cdr! xs '())
		       (values xs rest-xs)))
		    (else
		     ;; TODO: We want something here that does not require creating the closures.
		     ;; let-values would do.
		     (call-with-values
			 (lambda ()
			   (build-sorted! (- n 1) <? xs tmp))
		       (lambda (as xs)
			 (if (null? xs)
			     (values as '())
			     (call-with-values 
				 (lambda ()
				   (build-sorted! (- n 1) <? xs tmp))
			       (lambda (bs xs)
				 (set-cdr! tmp '())
				 (list-merge! <? as bs tmp)
				 (values (cdr tmp) xs)))))))))))
    (lambda (<? xs)
      "Given a binary predicate `<?` and a list `xs`, sort the `xs` in-place and return the new list."

      ;; The use of "32" here means no more than 2^32 elements in the list; this is a minor hack.
      ;; Cleaner would be to compute the next power of 2 no less than the length of the list.
      (build-sorted! 32 <? xs (cons #f #f)))))

;; Return #t if the list xs is sorted according to the binary predicate <?

(define (list-sorted? <? xs)
  "Return #t iff the list `xs` are sorted according to the binary predicate `<?`."
  (or (null? xs) (null? (cdr xs))
      (and (not (<? (cadr xs) (car xs)))
	   (list-sorted? <? (cdr xs)))))
