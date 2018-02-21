str = "abcd"

puts str.reverse

# or

rev = Array.new
i = 0

str.each_char do |c|
  rev[str.length - i] = c
  i = i + 1
end

puts rev.join("")
