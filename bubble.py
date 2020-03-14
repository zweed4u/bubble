#!/usr/bin/python3
# move the biggest item in a pass to its proper rightmost position
import time
import random

length_of_array = 10  # generate 10 numbers 
arr = [random.randint(0,10) for i in range(length_of_array)]  # 10 random numbers between 0 and 10

# [3, 3, 8, 8, 2, 2, 4, 4, 2, 5]
print(arr)

passes = 0
start = time.time()
while passes < length_of_array:
    # flag used to check if we've altered anything in a pass
    array_altered = False

    # subtract passes from how far into array index needs to be checked
    # bubble sort pushes everything to the right so each pass should need
    # to go less into the array
    for i in range(length_of_array-passes):
        # prevent out of bound error - accessing len(array)+1 index
        if i == length_of_array - 1:
            break
        v1 = arr[i]
        v2 = arr[i+1]
        # elements are in wrong order - swap and mark that we have altered the list
        if v1 > v2:
            array_altered = True
            arr[i+1] = v1
            arr[i] = v2
    # we made it through the whole array without swapping any elements - list is sorted - break
    if not array_altered:
        break
    passes+=1
end = time.time()
print(arr)

print(f"sort took {(end-start)*1000} milliseconds to sort array of {length_of_array} elements")
