module Gigasecond
  @@gigasecond = 10**9
  def self.from(time)
    raise "Invalid input" unless time.is_a? Time
    return time + @@gigasecond
  end
end
