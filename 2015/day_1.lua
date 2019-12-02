#! /usr/bin/env lua

if arg[1] == nil then
  error 'You can not run this script without argument'
end

floors = arg[1]

total = 0

for i = 1, #floors do

  char = floors:sub(i,i)

  if char == '(' then
    total = total + 1
  else
    total = total - 1
  end
end

print( total )
