# Daily Coding Problem: Problem #846 [Medium]

This problem was asked by Jane Street.

`cons(a, b)` constructs a pair, and `car(pair)` and `cdr(pair)`
returns the first and last element of that pair.
For example, `car(cons(3, 4))` returns 3,
and `cdr(cons(3, 4))` returns 4.

Given this implementation of cons:

```python
def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
```

Implement `car` and `cdr`.

## Analysis

I may not be implementing an idiomatic version of this,
a problem about Lisp-style primitives in Python.
I'm don't have current experience with either family of languages.
I've never understood the fine points of the Lisp-advocate
arguments about how `cons` should only deal with pairs
of elements, rather than arbitrary lists.

My reasoning goes like this:
`cons` returns a partially evaluated function.
Arguments a, b of `cons` get used in another function
invocation by nested function `pair`.
That invoked function has to have 2 arguments.
`car` and `cdr` have a single argument,
the partially evaluated function returned by `cons`.
I wrote `car` to take an argument `pair`.
It has a nested function `first` that requires
two argument, returning the first.
`cdr` is similar.

My [first version](cons1.py) spells it all out:

```python
def car(pair):
    def first(a,b):
        return a
    return pair(first)
def cdr(pair):
    def second(a,b):
        return b
    return pair(second)
```

I wrote a [second version](cons2.py) to show the similarity
of `car` and `cdr`.

I also did a [Go language](cons2.go) implementation.
It's similar to the Python version,
except it specifies argument and return types.

## Interview Analysis

This strikes me as a question from the
"pyton is cool, so do python to attract the best programmers"
era.
It also probably seperates the functional programming sheep
from the procedural programming goats.
Python has features that can be used in either modality,
but my guess is people who succeed at this question
would mostly be in the functional programming group.
I also like it as a "medium" question for up to mid-level candidates.
After that, a developer should probably be able to argue
the finer points of the lisp-advocate viewpoint about `cons`
doing only 2-cell lists.
