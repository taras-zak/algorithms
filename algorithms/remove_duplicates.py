def removeDuplicates(arr, n):
    j = 0
    for i in range(0, n-1):
        if arr[i] != arr[i + 1]:
            arr[j] = arr[i]
            j += 1
    arr[j] = arr[n-1]
    return arr[:j+1]


arr = [1, 2, 2, 3, 4, 4, 4, 5, 5,5]
n = len(arr)

# removeDuplicates() returns
# new size of array.
print(removeDuplicates(arr, n))