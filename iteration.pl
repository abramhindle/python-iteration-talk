# while loop

my $condition = 1;
while ($condition) {
    # ...
    $condition = !$condition;
}
# condition is False

#how many lines ?
my $count = 0;
$count++ while(<STDIN>);

my $x = 10;
while ($x > 0) {
    $x--;
}
# x is 0

# maybe you're not sure how many
# iterations you need?
my $x = 100.0;
while ($x > 1) {
    $x /= 3;
}

# for loop
my $sum = 0;
for my $i (1..9) {
    $sum += $i;
}
# sum is 45

# for loop
my $sum = 0;
for ( my $i = 1 ; $i < 100000; $i++ ) {
    $sum += $i;
}
# sum is 4 999 950 000

my $s = "";
for my $elm ("a","b","c") {
    $s .= $elm;
}
# s is abc

# The OO way
package OnlyEvens;
use Moose;
has seq => (is=>'rw', default=>sub{[]});
has index => (is => 'rw', default => 0);

sub has_next {
    my ($self) = @_;
    return $self->index < scalar(@{$self->seq});
}
sub next {
    my ($self) = @_;
    my $v = $self->seq->[$self->index];
    $self->index($self->index + 2);
    return $v;
}
1;

my $iter = OnlyEvens->new( seq => [1..10] );
while ($iter->has_next()) {
    print $iter->next().$/;
}

# add 1 to a list
my @v = (1..30);
my @u = map { $_ + 1 }  @v;
# u is now [2,3,4,..,31], v is still [1,2,3]

use File::Basename;

my @v = ("/home","/file","/usr/local");
my @u = map { basename $_ } @v;
my @u = map(uc,@v);
#['/HOME', '/FILE', '/USR/LOCAL']

import urllib2
urls = ["http://cbc.ca","http://gc.ca","http://alberta.ca"]
def get_url(url):
    return urllib2.urlopen(url).read()

pages = map(get_url, urls)


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


def recsum(l,i=0):
    if (i < len(l)):
        return l[i] + recsum(l,i+1)
    else:
        return 0

recsum(range(1,10))


p = multi.Pool( processes=2 )
pages = p.map(get_url, urls)
