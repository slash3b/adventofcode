#! /usr/bin/env lua

--
-- Important note on difference between : and . notation. From stackoverflow:
-- "The colon is for implementing methods that pass self as the first parameter. So x:bar(3,4)should be the same as x.bar(x,3,4)."
-- 

function split_by(dimensions, char)

    x = string.find(dimensions, char)
    y = string.find(dimensions, char, x + 1)

    -- keep it simple and stupid.. and not flexible
    dm = {}
    dm.l = dimensions:sub(1, x -1)
    dm.w = dimensions:sub(x+1, y-1)
    dm.h = dimensions:sub(y+1, #dimensions)

    return dm

end

total_material = 0

for line in io.lines('./input') do
    dm = split_by(line, 'x')

    -- 2*l*w + 2*w*h + 2*h*l plus one of the smallest side

    a_side = dm.l * dm.w
    b_side = dm.w * dm.h
    c_side = dm.h * dm.l

    total =  math.min(a_side, b_side, c_side) + (2*a_side) + (2*b_side) + (2*c_side)
    total_material = total_material + total

end

print(total_material)




