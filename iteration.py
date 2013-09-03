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
