class Matrix
  attr_reader :rows
  def initialize(string)
    @rows = string.split("\n").map {|v| v.split(' ').map(&:to_i) }
  end

  def gen_columns
    output = Array.new(@rows.length) { Array.new(@rows.length) }
    @rows.each_with_index { |row, index_r|
      row.each_with_index { |element, index_c|
        #STDERR.printf("output[%s][%s] = [%s]\n", index_c, index_r, element)
        output[index_c][index_r] = element
      }
    }
    return output
  end

  def columns
    @columns ||= gen_columns
  end
end
