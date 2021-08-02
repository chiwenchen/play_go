require 'pry'
def getcombination(target, max) # return [[]]
	return [[]] if target == 0
	return nil if target < 0

	result = []

	(1..max).map do |i|
		remain = target - i
		getcombination(remain, i)&.map do |e|
			# e -> array
			e&.push(i)
		end
	end
end
