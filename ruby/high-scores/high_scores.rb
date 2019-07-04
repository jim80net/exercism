class HighScores
  attr_accessor :scores
  def initialize(scores=[])
    @scores = Array.new(scores)
  end

  def latest
    @scores.last
  end

  def personal_best
    @scores.sort.last
  end

  def personal_top_three
    start_index = @scores.length - 3
    start_index = 0 if start_index < 0
    @scores.sort[start_index..-1].reverse
  end
end
