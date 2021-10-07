def count(a,b)
    return a if a == b
    x = a
    y = b
    if x > y
        x = x - y
        count(x,y)
    elsif x < y
        y = y - x
        count(x,y)
    end
    puts x
    puts y
end


puts count(2437, 875)

