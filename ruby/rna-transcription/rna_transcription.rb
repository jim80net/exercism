module Complement
  def self.of_dna(nucleotide)
    nucleotide.split('').map do |v|
      case v
      when 'G'
        'C'
      when 'C'
        'G'
      when 'A'
        'U'
      when 'T'
        'A'
      else
        v
      end
    end.join('')
  end
end
