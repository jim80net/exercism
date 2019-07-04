module Complement
  DNA = 'CATG'
  RNA = 'GUAC'

  def self.of_dna(nucleotides)
    nucleotides.tr(DNA, RNA)
  end
end
