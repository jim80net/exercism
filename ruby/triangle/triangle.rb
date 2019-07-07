class Triangle
  def initialize(sides)
    @a = sides[0]
    @b = sides[1]
    @c = sides[2]
  end

  def equilateral?
    is_a_triangle && @a == @b && @b == @c
  end

  def isosceles?
    is_a_triangle && ( @a == @b || @a == @c || @b == @c )
  end

  def scalene?
    is_a_triangle && ( @a != @b && @a != @c && @b != @c )
  end

  def is_a_triangle
    @a > 0 &&
      @b > 0 &&
      @c > 0 &&
      @a + @b >= @c &&
      @a + @c >= @b &&
      @b + @c >= @a
  end

end
    ys_a_triangle && !isosceles?
