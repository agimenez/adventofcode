# [Advent of Code](/)

  * [[About]](/2023/about)
  * [[Events]](/2023/events)
  * [[Shop]](https://teespring.com/stores/advent-of-code)
  * [[Settings]](/2023/settings)
  * [[Log Out]](/2023/auth/logout)

agimenez [(AoC++)](/2023/support "Advent of Code Supporter") 24*

#   {year=>[2023](/2023)}

  * [[Calendar]](/2023)
  * [[AoC++]](/2023/support)
  * [[Sponsors]](/2023/sponsors)
  * [[Leaderboard]](/2023/leaderboard)
  * [[Stats]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Smarty](https://www.smarty.com/advent-of-code) \- Join our private
leaderboard and solve our puzzles for BIG PRIZES!!! ----------------- Address
Validation and Autocomplete, and more!

## \--- Day 16: The Floor Will Be Lava ---

With the beam of light completely focused *somewhere* , the reindeer leads you
deeper still into the Lava Production Facility. At some point, you realize
that the steel facility walls have been replaced with cave, and the doorways
are just cave, and the floor is cave, and you're pretty sure this is actually
just a giant cave.

Finally, as you approach what must be the heart of the mountain, you see a
bright light in a cavern up ahead. There, you discover that the beam of light
you so carefully focused is emerging from the cavern wall closest to the
facility and pouring all of its energy into a contraption on the opposite
side.

Upon closer inspection, the contraption appears to be a flat, two-dimensional
square grid containing *empty space* (`.`), *mirrors* (`/` and `\`), and
*splitters* (`|` and `-`).

The contraption is aligned so that most of the beam bounces around the grid,
but each tile on the grid converts some of the beam's light into *heat* to
melt the rock in the cavern.

You note the layout of the contraption (your puzzle input). For example:

[code]

    .|...\....
    |.-.\.....
    .....|-...
    ........|.
    ..........
    .........\
    ..../.\\..
    .-.-/..|..
    .|....-|.\
    ..//.|....
    
[/code]

The beam enters in the top-left corner from the left and heading to the
*right*. Then, its behavior depends on what it encounters as it moves:

  * If the beam encounters *empty space* (`.`), it continues in the same direction.
  * If the beam encounters a *mirror* (`/` or `\`), the beam is *reflected* 90 degrees depending on the angle of the mirror. For instance, a rightward-moving beam that encounters a `/` mirror would continue *upward* in the mirror's column, while a rightward-moving beam that encounters a `\` mirror would continue *downward* from the mirror's column.
  * If the beam encounters the *pointy end of a splitter* (`|` or `-`), the beam passes through the splitter as if the splitter were *empty space*. For instance, a rightward-moving beam that encounters a `-` splitter would continue in the same direction.
  * If the beam encounters the *flat side of a splitter* (`|` or `-`), the beam is *split into two beams* going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a `|` splitter would split into two beams: one that continues *upward* from the splitter's column and one that continues *downward* from the splitter's column.

Beams do not interact with other beams; a tile can have many beams passing
through it at the same time. A tile is *energized* if that tile has at least
one beam pass through it, reflect in it, or split in it.

In the above example, here is how the beam of light bounces around the
contraption:

[code]

    >|<<<\....
    |v-.\^....
    .v...|->>>
    .v...v^.|.
    .v...v^...
    .v...v^..\
    .v../2\\..
    <->-/vv|..
    .|<<<2-|.\
    .v//.|.v..
    
[/code]

Beams are only shown on empty tiles; arrows indicate the direction of the
beams. If a tile contains beams moving in multiple directions, the number of
distinct directions is shown instead. Here is the same diagram but instead
only showing whether a tile is *energized* (`#`) or not (`.`):

[code]

    ######....
    .#...#....
    .#...#####
    .#...##...
    .#...##...
    .#...##...
    .#..####..
    ########..
    .#######..
    .#...#.#..
    
[/code]

Ultimately, in this example, `*46*` tiles become *energized*.

The light isn't energizing enough tiles to produce lava; to debug the
contraption, you need to start by analyzing the current situation. With the
beam starting in the top-left heading right, *how many tiles end up being
energized?*

To begin, [get your puzzle input](16/input).

Answer:

You can also [Shareon
[Twitter](https://twitter.com/intent/tweet?text=%22The+Floor+Will+Be+Lava%22+%2D+Day+16+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F16&related=ericwastl&hashtags=AdventOfCode)
[Mastodon](javascript:void\(0\);)] this puzzle.

