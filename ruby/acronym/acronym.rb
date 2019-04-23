module Acronym
  def self.abbreviate(text)
    text
      .split(/[ -]/)
      .delete_if {|v| ['','-'].include? v}
      .map {|v| v[0].upcase}
      .join
  end
end
