#! /usr/bin/env lua

if arg[1] == nil then
  error 'You can not run this script without argument'
end

floors = arg[1]

total = 0
basement_entered = false

-- #floor means length of a string floor
-- so for loop goes from 1 up to length of floor string
for i = 1, #floors do

  char = floors:sub(i,i)

  if char == '(' then
    total = total + 1
  else
    total = total - 1
  end

  if total < 0 and basement_entered == false then
    basement_entered = true
    print('Basement was entered on ' .. i .. ' index')
  end
end

print('Santa\'s floor is ' ..  total )
