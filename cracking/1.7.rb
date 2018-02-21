def display(matrix, rows)
  0.upto(rows-1) do |m|
    puts matrix[m].join " "
  end
end

row_count = 10
col_count = 10
matrix = []

0.upto(row_count-1) do |m|
  matrix[m] = []
  0.upto(col_count-1) do |n|
    matrix[m][n] = rand(9)
  end
end

puts "INITIAL"
display(matrix, row_count)

solution = []
matrix.each_with_index do |row, row_index|
  row.each_with_index do |value, col_index|
    if value == 0
      0.upto(row_count-1) do |m|
        solution[m] ||= []
        solution[m][col_index] = 0
      end
      0.upto(col_count-1) do |n|
        solution[row_index][n] = 0
      end
    else
      solution[row_index] ||= []
      solution[row_index][col_index] ||= value
    end
  end
end

puts "FINAL"
display(solution, row_count)
