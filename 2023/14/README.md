# [Advent of Code](/)

  * [[About]](/2023/about)
  * [[Events]](/2023/events)
  * [[Shop]](https://teespring.com/stores/advent-of-code)
  * [[Settings]](/2023/settings)
  * [[Log Out]](/2023/auth/logout)

agimenez [(AoC++)](/2023/support "Advent of Code Supporter") 22*

#       /^[2023](/2023)$/

  * [[Calendar]](/2023)
  * [[AoC++]](/2023/support)
  * [[Sponsors]](/2023/sponsors)
  * [[Leaderboard]](/2023/leaderboard)
  * [[Stats]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Spotify](https://engineering.atspotify.com/) \- Follow our engineering blog
to see how our developers solve complex tech problems, at scale, every day.

## \--- Day 14: Parabolic Reflector Dish ---

You reach the place where all of the mirrors were pointing: a massive
[parabolic reflector dish](https://en.wikipedia.org/wiki/Parabolic_reflector)
attached to the side of another large mountain.

The dish is made up of many small mirrors, but while the mirrors themselves
are roughly in the shape of a parabolic reflector dish, each individual mirror
seems to be pointing in slightly the wrong direction. If the dish is meant to
focus light, all it's doing right now is sending it in a vague direction.

This system must be what provides the energy for the lava! If you focus the
reflector dish, maybe you can go where it's pointing and use the light to fix
the lava production.

Upon closer inspection, the individual mirrors each appear to be connected via
an elaborate system of ropes and pulleys to a large metal platform below the
dish. The platform is covered in large rocks of various shapes. Depending on
their position, the weight of the rocks deforms the platform, and the shape of
the platform controls which ropes move and ultimately the focus of the dish.

In short: if you move the rocks, you can focus the dish. The platform even has
a control panel on the side that lets you *tilt* it in one of four directions!
The rounded rocks (`O`) will roll when the platform is tilted, while the cube-
shaped rocks (`#`) will stay in place. You note the positions of all of the
empty spaces (`.`) and rocks (your puzzle input). For example:

[code]

    O....#....
    O.OO#....#
    .....##...
    OO.#O....O
    .O.....O#.
    O.#..O.#.#
    ..O..#O..O
    .......O..
    #....###..
    #OO..#....
    
[/code]

Start by tilting the lever so all of the rocks will slide *north* as far as
they will go:

[code]

    OOOO.#.O..
    OO..#....#
    OO..O##..O
    O..#.OO...
    ........#.
    ..#....#.#
    ..O..#.O.O
    ..O.......
    #....###..
    #....#....
    
[/code]

You notice that the support beams along the north side of the platform are
*damaged* ; to ensure the platform doesn't collapse, you should calculate the
*total load* on the north support beams.

The amount of load caused by a single rounded rock (`O`) is equal to the
number of rows from the rock to the south edge of the platform, including the
row the rock is on. (Cube-shaped rocks (`#`) don't contribute to load.) So,
the amount of load caused by each rock in each row is as follows:

[code]

    OOOO.#.O.. 10
    OO..#....#  9
    OO..O##..O  8
    O..#.OO...  7
    ........#.  6
    ..#....#.#  5
    ..O..#.O.O  4
    ..O.......  3
    #....###..  2
    #....#....  1
    
[/code]

The total load is the sum of the load caused by all of the *rounded rocks*. In
this example, the total load is `*136*`.

Tilt the platform so that the rounded rocks all roll north. Afterward, *what
is the total load on the north support beams?*

To begin, [get your puzzle input](14/input).

Answer:

You can also [Shareon
[Twitter](https://twitter.com/intent/tweet?text=%22Parabolic+Reflector+Dish%22+%2D+Day+14+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F14&related=ericwastl&hashtags=AdventOfCode)
[Mastodon](javascript:void\(0\);)] this puzzle.

