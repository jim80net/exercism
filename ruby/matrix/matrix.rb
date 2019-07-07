class Matrix
  attr_reader :rows
  def initialize(string)
    @rows = string.split("\n").map {|v| v.split(' ').map(&:to_i) }
  end

  def columns
    @columns ||= @rows.transpose
  end
end
