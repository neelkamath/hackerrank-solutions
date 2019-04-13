#!/bin/python3

import math
import os
import random
import re
import sys


def minimumAbsoluteDifference(arr):
    arr.sort()
    minimum = 1e9
    for index, num in enumerate(arr[:len(arr) - 1]):
        newMin = abs(num - arr[index + 1])
        if newMin < minimum:
            minimum = newMin
    return minimum


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = minimumAbsoluteDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
