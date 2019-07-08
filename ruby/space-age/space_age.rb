class SpaceAge

  EARTH_YEARS_TO_PLANET_YEARS = {
    mercury: 0.2408467,
    venus: 0.61519726,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }

  def initialize(time_in_seconds)
    @time_in_seconds = time_in_seconds.to_f
  end

  def on_earth
    @time_in_seconds / 31557600
  end

  EARTH_YEARS_TO_PLANET_YEARS.each do |planet, orbital_period|
    define_method("on_#{planet}") do
      on_earth / orbital_period
    end
  end

end
