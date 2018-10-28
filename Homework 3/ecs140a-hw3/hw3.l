;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; Your solution goes in this file ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defun element-count (a l)
    (if (null l)
        0
        (if (equal a (car l)) 
            ;;add 1
            (+ 1 (element-count a (cdr l)) )
            (element-count a (cdr l))
        ) 
        
    )
)

(defun maximum (l)
   (if (null l)
       0   ;;if end of list return null
       ;;beggining of list is greater than min mean max
        (if(> (car l) (maximum(cdr l))) 
        ;;return beggining of list
            (car l)
        ;;else return campared value
            (maximum (cdr l))
        )
    )
)

(defun minimum (l)
    ;;(+ (car l))
   ;;(setq x (maximum l) )
   (if (null l)
        1000000000000000000000000000000000000   ;;if end of list return null
        ;(list '())
        (if(> (car l) (minimum(cdr l))) 
        ;;return beggining of list
            (minimum (cdr l))
        ;;else return campared value
            (car l)
        )
    )
)

(defun sum (l)
    (if (null l) 
        0
        (+ (car l) (sum (cdr l)) )
    )
)

(defun mean (l)
    (if (null l) 
        0
        (/ (sum l) (list-length l) )
    )
)

(defun min-mean-max (l)
    (if (null l)
        nil
       (setq lis (list (minimum l) (mean l) (maximum l)))
    )
    
    ;;(setq lis (list (minimum l) (mean l) (maximum l)))
    ;;lis
)

(defun s (l)
    (append (cdr l) (list (car l)))
)

(defun shift (n l)

    (if (null l)
        nil
        
        (if (= n 0)
            l
            (shift (- n 1) (s l))
;;            (append (shift(- n 1) (cdr l)) (list(car l)))
        )
    )    
        
)

(defun my-filter (condp lst)
    (delq nil
        (mapcar (lambda (x) (and (funcall condp x) x)) lst)
    )
)


(defun small (n l)
    ;;(remove-if-not
    ;;    (lambda (x) (< x n)) l
    ;;)
   ;; (maplist (lambda (x) (< x n)) l)
    
    (if (or (null l) (null n)) 
        nil
        (if (<= n (car l))
            (small n (cdr l))
            (cons (car l) (small n (cdr l)))
        )
    )
)

(defun big (n l)
    ;;(remove-if-not
    ;;    (lambda (x) (>= x n)) l
    ;;)
    ;;(maplist (lambda (x) (>= x n)) l)
    (if (or (null l) (null n)) 
        nil
        (if (> n (car l)) 
            (big n (cdr l))
            (cons (car l) (big n (cdr l)))
        )
    )
)

(defun pivot (n l)
    (if (null l)
        (list '() '())
        (list(small n l) (big n l))
    )
)

(defun create (a l)
    (if (null l)
        nil
        (if (equal a (car l))
        
            (break-list a (cdr l) ) 
            ;(list (cdr l))
            ;()
            (cons (car l) (create a (cdr l)) )
        ) 
        
    )
)

;cons create break
(defun break-list (a l)
    
    (if (null l)
        nil
        ;(cons (create a l) ())
        (if (equal a (car l))
            ;()
            ;(cons () (create a (cdr l) ))
            (break-list a (cdr l))
            ;(create a (cdr l))

            ;(cons (create a l) ())
            (list (create a l))
        )
    )
)



(defun quicksort (l)
    (if (null l)
        nil
        (append (quicksort (small (car l) (cdr l))) (cons (car l) nil) (quicksort (big (car l) (cdr l))) )
    )
)

(defun same(l1 l2)
    (if (equal l1 l2)
        t
        nil
    )
)


(defun excl(l1)

    (if (> (element-count '! l1) 0)
        t
        nil
    )
)

(defun quest(l1 l2)
    ;(if (equal (length l1) (length l2) )
    (if (and (null l1) (null l2))
        t
        (if (equal (car l1) (car l2))
            (match (cdr l1) (cdr l2))
            
            (if (or (equal '? (car l1)) (equal '? (car l2)) )
                ;(or( (equal ('? (car l1))) (equal ('? (car l2))) ) )
                (match (cdr l1) (cdr l2))
                (if (or (equal '! (car l1)) (equal '! (car l2)) )
                    (match (cdr l1) (cdr l2))
                    nil
                )
            )
            
        )
    )
       ; nil
    ;)

)

;;part 2
(defun match (l1 l2)
    
    ;(mapcar #'equal l1 l2)
    
    (if (equal l1 l2)
        t
        (if (equal (length l1) (length l2) )
            (quest l1 l2)
            (if (< (length l1) (length l2))
                (excl l1)
                nil
            )
        )
    )

)


(defun add-row (r1 r2 row)
    (if (and (not (null r1)) (not (null r2)))
        (add-row (cdr r1) (cdr r2)
            (cons (+ (car r1) (car r2)) row)
        )
        (reverse row)
    )
)

(defun add-each (m1 m2 res)
    (if (and (not (null m1)) (not (null m2)))
        (let ((row (add-row (car m1) (car m2) ())))
            (add-each (cdr m1) (cdr m2) (cons row res))
        )
        (reverse res)
    )
)

(defun matrix-add (m1 m2)
    (if (and (null m1) (null m2) )
        nil
        (add-each m1 m2 ())   
    )
)

(defun col (matrix i)
    (mapcar #'(lambda (row) (elt row i)) matrix)
)

(defun row (matrix i)
    (elt matrix i)
)

(defun matrix-multiply (m1 m2)
    (loop for r from 0 below (length m1)
        collect (loop for c from 0 below (length (row m2 0))
            ;multiply row index by col index
            collect (apply #'+ (mapcar #'* (row m1 r) (col m2 c)))
            ;(+ (mapcar #'* (row m1 r) (col m2 c)))
        )
    )
)

(defun cars (matrix)
  ;Return a list with all the cars of the lists in matrix
    (if (null matrix)
        nil
        (cons (car (car matrix)) (cars (cdr matrix)))
    )
)

(defun cdrs (matrix)
  ;Return a list with all the cdrs of the lists in matrix
    (if (null matrix)
        nil
        (cons (cdr (car matrix)) (cdrs (cdr matrix)))
    )
)

(defun matrix-transpose (matrix)
  ;Transpose matrix
    (if (null matrix)
        nil
        (if (null (car matrix))
            nil
            (cons (cars matrix) (matrix-transpose (cdrs matrix)))
        )
    )
)




 
                        
                        
