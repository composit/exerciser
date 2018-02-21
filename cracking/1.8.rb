orig = "waterbottle"
rota = "erbottlewat"
nota = "otherstring"

def is_rotated?(o, z)
  r = false

  if o.length == z.length
      r = (o+o).include?(z)
  end

  r
end

[[orig, rota], [orig, nota]].each do |pair|
  o, z = pair
  puts "#{o} and #{z}? #{is_rotated?(o, z)}"
end
