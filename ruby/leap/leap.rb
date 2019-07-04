module Year
  def self.leap?(year)
    output = false
    output = true if year % 4 == 0
    output = false if year % 100 == 0
    output = true if year % 400 == 0
    output
  end

end
