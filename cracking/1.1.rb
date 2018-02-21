str = "abcdefg"
chars = Hash.new
unique = true
str.each_char do |c|
  if chars[c]
    unique = false
    break
  end
  chars[c] = true
end

puts "unique? #{unique}"
