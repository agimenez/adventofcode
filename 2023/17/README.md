# [Advent of Code](/)

  * [[About]](/2023/about)
  * [[Events]](/2023/events)
  * [[Shop]](https://teespring.com/stores/advent-of-code)
  * [[Settings]](/2023/settings)
  * [[Log Out]](/2023/auth/logout)

agimenez [(AoC++)](/2023/support "Advent of Code Supporter") 26*

#    0xffff&[2023](/2023)

  * [[Calendar]](/2023)
  * [[AoC++]](/2023/support)
  * [[Sponsors]](/2023/sponsors)
  * [[Leaderboard]](/2023/leaderboard)
  * [[Stats]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[it-jobs.de](https://it-jobs.de/?sta_cmp=advent_2023) \- designed with love
for IT ..""".....""" ."...."."...." ."....."....." .."........." ..."......."
....."..." ......."

## \--- Day 17: Clumsy Crucible ---

The lava starts flowing rapidly once the Lava Production Facility is
operational. As you leave, the reindeer offers you a parachute, allowing you
to quickly reach Gear Island.

As you descend, your bird's-eye view of Gear Island reveals why you had
trouble finding anyone on your way up: half of Gear Island is empty, but the
half below you is a giant factory city!

You land near the gradually-filling pool of lava at the base of your new
*lavafall*. Lavaducts will eventually carry the lava throughout the city, but
to make use of it immediately, Elves are loading it into large
[crucibles](https://en.wikipedia.org/wiki/Crucible) on wheels.

The crucibles are top-heavy and pushed by hand. Unfortunately, the crucibles
become very difficult to steer at high speeds, and so it can be hard to go in
a straight line for very long.

To get Desert Island the machine parts it needs as soon as possible, you'll
need to find the best way to get the crucible *from the lava pool to the
machine parts factory*. To do this, you need to minimize *heat loss* while
choosing a route that doesn't require the crucible to go in a *straight line*
for too long.

Fortunately, the Elves here have a map (your puzzle input) that uses traffic
patterns, ambient temperature, and hundreds of other parameters to calculate
exactly how much heat loss can be expected for a crucible entering any
particular city block.

For example:

[code]

    2413432311323
    3215453535623
    3255245654254
    3446585845452
    4546657867536
    1438598798454
    4457876987766
    3637877979653
    4654967986887
    4564679986453
    1224686865563
    2546548887735
    4322674655533
    
[/code]

Each city block is marked by a single digit that represents the *amount of
heat loss if the crucible enters that block*. The starting point, the lava
pool, is the top-left city block; the destination, the machine parts factory,
is the bottom-right city block. (Because you already start in the top-left
block, you don't incur that block's heat loss unless you leave that block and
then return to it.)

Because it is difficult to keep the top-heavy crucible going in a straight
line for very long, it can move *at most three blocks* in a single direction
before it must turn 90 degrees left or right. The crucible also can't reverse
direction; after entering each city block, it may only turn left, continue
straight, or turn right.

One way to *minimize heat loss* is this path:

[code]

    2 *>* *>* 34 *^* *>* *>* *>* 1323
    32 *v* *>* *>* *>* 35 *v* 5623
    32552456 *v* *>* *>* 54
    3446585845 *v* 52
    4546657867 *v* *>* 6
    14385987984 *v* 4
    44578769877 *v* 6
    36378779796 *v* *>*
    465496798688 *v*
    456467998645 *v*
    12246868655 *<* *v*
    25465488877 *v* 5
    43226746555 *v* *>*
    
[/code]

This path never moves more than three consecutive blocks in the same direction
and incurs a heat loss of only `*102*`.

Directing the crucible from the lava pool to the machine parts factory, but
not moving more than three consecutive blocks in the same direction, *what is
the least heat loss it can incur?*

To begin, [get your puzzle input](17/input).

Answer:

You can also [Shareon
[Twitter](https://twitter.com/intent/tweet?text=%22Clumsy+Crucible%22+%2D+Day+17+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F17&related=ericwastl&hashtags=AdventOfCode)
[Mastodon](javascript:void\(0\);)] this puzzle.

