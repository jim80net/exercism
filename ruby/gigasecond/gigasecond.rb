module Gigasecond
  def self.from(time)
    raise unless time.is_a? Time
    return time + 1000000000
  end
end
