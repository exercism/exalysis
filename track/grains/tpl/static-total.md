- You are calculating the total number of grains by summing the values for each square, which is straightforward, but there's a short-cut. Because the answer is always the same, assuming the chessboard has 64 squares (which most do), you can return a constant expression for the value of `Total`. The mathematical formula for the sum of N consecutive powers of 2 is 2^(N+1) - 1. Since square 1 contains 2^0 grains, square 2 contains 2^1 grains, and so on, we are summing the series 2^0 + 2^1 + ... + 2^63. The formula for this sum is 2^64 - 1. Can you see how to use the bit-shift operator (`<<`) to implement this formula?

(If you're wondering why this is the case, consider the following binary numbers, which represent the series 2^0, 2^1, etc:

    0001
    0010
    0100
    1000

What's the sum of this series? That's easy:

    1111

If you add one to this, you get:

    10000

So the sum of the first N powers of 2 is one less than the next highest power of 2. Does that help?)
