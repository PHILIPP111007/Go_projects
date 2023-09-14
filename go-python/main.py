#!/usr/bin/python3
# -*- coding: utf-8 -

#####################################################
#
# To compile Go modules with C types to work with Python run:
# go build -o functions.so -buildmode=c-shared functions.go
#
#####################################################

import ctypes

Lib = ctypes.CDLL('go-module/functions.so')
GoSum = Lib.go_sum  # or Lib.__getattr__('go_sum')
GoLoop = Lib.go_loop  # or Lib.__getattr__('go_loop')
GoSum.argtypes = [ctypes.c_longlong, ctypes.c_longlong]
GoSum.restype = ctypes.c_longlong


def time_counter(func):
    """Custom decorator."""
    import time

    def wrapper(*args, **kwargs):
        t_1 = time.time()
        func(*args, **kwargs)
        t_2 = time.time()
        z = t_2 - t_1
        print(f'{func.__name__} = {z}')
    return wrapper


@time_counter
def PyLoop():
    for _ in range(1_000_000):
        pass

PyLoop() # PyLoop = 0.028534650802612305
time_counter(GoLoop)() # go_loop = 0.000476837158203125

s = GoSum(10, 1)
print(s) # 11
