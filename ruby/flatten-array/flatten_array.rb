module FlattenArray
  def self.flatten(array)
    result = []
    array.each do |element|
      if element.is_a? Array
        result += flatten(element)
      elsif element === nil
        break
      else
        result.push element
      end
    end

    return result
  end
end
