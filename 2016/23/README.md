## 2016 - Day 23
```
curl https://adventofcode.com/2016/day/23
```

## IMPORTANT NOTE
The provided code is calculating something with the factorial of the contents of register "a".
The original code (and asembunny interpreter) does not have the "mul" instruction, which means that it's **very** slow.

On an X1 Carbon, my naive, string-based solution took 20 minutes for part 2.
After including the "mul" instruction (and modifying my input to `test.txt`), it takes <2ms

To execute the normal (slow) code, `just a` (5ms p1, 20m p2)
To execute the "optimised" code, `just t` (1ms p1 and p2)

