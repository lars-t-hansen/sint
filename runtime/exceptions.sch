;; -*- indent-tabs-mode: nil; fill-column: 100 -*-
;;
;; R7RS 6.11 "Exceptions"

(define (error msg . irritants)
  (if (not (string? msg))
      (error "error: the first argument must be a string" msg)
      (sint:report-error msg irritants)))

