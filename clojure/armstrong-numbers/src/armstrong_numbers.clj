(ns armstrong-numbers
  "Tells you whether a number is an armstrong number."
)

(defn exp
  "Raise x n times."
  [x n]
  (reduce * (repeat n x)))

(defn seq_num
  "Turn an number into a seq of its characters"
  [num]
  (map #(- (int %) 48) (str num)))


(defn armstrong_calc
  "Sum each character raised to the number of characters."
  [num]
  (reduce + (map #(exp % (count (seq_num num))) (seq_num num))))

(defn armstrong?
  "An armstrong number is equal to the sum of each character raised to the number of characters."
  [num]
  (==  num (armstrong_calc num)))
