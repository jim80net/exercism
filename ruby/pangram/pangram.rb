module Pangram
  def self.pangram?(arg)
    test_string = arg.downcase.split('').uniq.sort.join('')
    test_string.include? "abcdefghijklmnopqrstuvwxyz"
  end
end
