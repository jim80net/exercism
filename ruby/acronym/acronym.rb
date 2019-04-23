module Acronym
  def self.abbreviate(text)
    text
      .scan(/\b[a-zA-Z]/)
      .map {|v| v[0].upcase}
      .join
  end
end
