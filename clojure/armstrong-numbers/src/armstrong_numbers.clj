(ns armstrong-numbers)

(defn armstrong? [num]
  "An armstrong number is equal to the sum of each character raised to the number of characters"
  (==  num (armstrong_calc num)))

(defn armstrong_calc [num]
  "Sum each character raised to the number of characters"
  (reduce + (map #(exp % (count (seq_num num))) (seq_num num))))

(defn exp [x n]
  "Raise x n times."
  (reduce * (repeat n x)))


(defn seq_num [num]
  "Turn an number into a seq of its characters"
  (map #(- (int %) 48) (str num)))




