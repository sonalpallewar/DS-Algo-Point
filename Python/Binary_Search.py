# Binary Search = An efficent way to search in a sorted list
def binary_new_function(arr, low, high, x):
    print("NEW function in binary search")
    while low < high:
        mid = low + (high - 1) // 2
        if (arr[mid] == x):
            return mid
        elif (arr[mid] < x):
            high = mid - 1
        else:
            low = mid + 1
    return -1


def binary(arr, low, high, x):
    while low < high:
        mid = low + (high - 1) // 2
        if (arr[mid] == x):
            return mid
        elif (arr[mid] < x):
            low = mid + 1
        else:
            high = mid - 1
    return -1


print('Enter a count of number in a list or array')
arr = list(map(int, input().split()))
n = len(arr)

arr.sort()
x = int(input("Enter the number to find\n"))
result = binary(arr, 0, len(arr) - 1, x)
if (result == -1):
    print("Element not found")
else:
    print("Element found at ", result + 1, " position")

'''
Time complexity ==> O(logn)
Space Complexity ==> O(1)

I/O ==>

Enter a list of number
234 567 321 345 123 110 767 111
Enter the number to find
767
Element found at  8  position
'''
