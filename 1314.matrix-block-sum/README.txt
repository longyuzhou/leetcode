another solution:

1  2    1  3
     =>
3  4    4 10
f1(y, x) = sum(m[0][0], ..., m[y][x])
         = m[y][x] + f1(y-1, x) + f1(y, x-1) - f1(y-1, x-1)

1  2  3    12 21 16
4  5  6 => 27 45 33
7  8  9    24 39 28
f2(y, x, k) = sum(m[y-k][x-k], ..., m[y+k][x+k])
            = f1(y+k, x+k) - f1(y-k-1, x) - f1(y, x-k-1) + f1(y-k-1, x-k-1)
