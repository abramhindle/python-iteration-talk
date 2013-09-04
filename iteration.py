# while loop

condition = True
while condition:
    """block"""
    condition = not condition
# condition is False

x = 10
while x > 0:
    x = x - 1
# x is 0

# maybe you're not sure how many
# iterations you need?
x = 100.0
while x > 1:
    x = x / 3

# for loop
sum = 0
for i in range(1,10):
    sum = sum + i
# sum is 45
sum = 0
for i in xrange(1,100000):
    sum = sum + i
# sum is 4 999 950 000
s = ""
for elm in ["a","b","c"]:
    s = s + elm
# s is abc

# iterable
class OnlyEvens(object):
    def __init__(self,s):
        self.sequence = s
        self.index = 0

    def __iter__(self):
        return self

    def next(self):
        if self.index >= len(self.sequence):
            raise StopIteration
        v = self.sequence[self.index]
        self.index += 2
        return v

oe = OnlyEvens(range(1,10))
for even in oe:
    print(even)

# add 1 to a list
v = [1,2,3]
u = map((lambda x: x+1), v)
# u is now [2,3,4], v is still [1,2,3]

def basename(path):
    return path.split("/")[-1]

v = ["/home","/","/usr/local"]
u = map(basename, v)
#['home', '', 'local']

# this is why you want blocks with 
# few dependencies!
import multiprocessing as multi
def square(x):
    return x * x

p = multi.Pool( processes=8 )
u = p.map(square, range(1,1000000))
len(u)
#999999

import operator
l = range(1,1000000)
p = multi.Pool( processes=2 )
u = reduce(operator.add, p.map(square, l))
v = sum(p.map(square, l))

# parallel map
def parallel_square(l):
    p = multi.Pool( processes=2 )
    return p.map(square, l)

# parallel reduce
def parallel_sum(l):
    p = multi.Pool( processes=2 )
    return sum(p.map(sum, [ l[0:len(l)/2], l[len(l)/2:len(l)] ]))

parallel_sum( parallel_square(l))

