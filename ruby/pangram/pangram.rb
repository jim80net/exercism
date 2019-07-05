module Pangram
  def self.pangram?(arg)
    test_string = arg.downcase.chars
    ([*('a'..'z')] - test_string).empty?
  end
end
