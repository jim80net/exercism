class SpaceAge

  SPACE_AGE_EARTH_YEARS_TO_PLANET_YEARS = {
    mercury: 0.2408467,
    venus: 0.61519726,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }

  def initialize(seconds)
    @seconds = seconds.to_f
  end

  def on_earth
    @seconds / 31557600
  end

  SPACE_AGE_EARTH_YEARS_TO_PLANET_YEARS.keys.each { |key|
    define_method :"on_#{key.to_s}" do
      on_earth / SPACE_AGE_EARTH_YEARS_TO_PLANET_YEARS[key]
    end
  }

end
