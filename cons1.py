#!/usr/bin/env python

# cons function as defined in problem statement
def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

def car(pair):
    def first(a,b):
        return a
    return pair(first)
def cdr(pair):
    def second(a,b):
        return b
    return pair(second)

# Try out the functions
print("should be 3: ", car(cons(3, 4)))
print("should be 4: ", cdr(cons(3, 4)))

