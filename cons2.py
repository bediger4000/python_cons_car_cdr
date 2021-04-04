#!/usr/bin/env python

# cons function as defined in problem statement
def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

def first(a,b):
    return a
def second(a,b):
    return b
def apply(f1):
    def applyinner(f2):
        return f2(f1)
    return applyinner
car = apply(first)
cdr = apply(second)

# Try out the functions
print("should be 3: ", car(cons(3, 4)))
print("should be 4: ", cdr(cons(3, 4)))

