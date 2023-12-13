# [Advent of Code](/)

  * [[About]](/2023/about)
  * [[Events]](/2023/events)
  * [[Shop]](https://teespring.com/stores/advent-of-code)
  * [[Settings]](/2023/settings)
  * [[Log Out]](/2023/auth/logout)

agimenez [(AoC++)](/2023/support "Advent of Code Supporter") 22*

#        λy.[2023](/2023)

  * [[Calendar]](/2023)
  * [[AoC++]](/2023/support)
  * [[Sponsors]](/2023/sponsors)
  * [[Leaderboard]](/2023/leaderboard)
  * [[Stats]](/2023/stats)

Our [sponsors](/2023/sponsors) help make Advent of Code possible:

[Best
Buy](https://jobs.bestbuy.com/bby?id=career_area&content=technology&career_site=Digital%20and%20Technology,Data%20and%20Analytics&spa=1&s=req_id_num)
\- Our purpose is to enrich lives through technology. Join us!

## \--- Day 13: Point of Incidence ---

With your help, the hot springs team locates an appropriate spring which
launches you neatly and precisely up to the edge of *Lava Island*.

There's just one problem: you don't see any *lava*.

You *do* see a lot of ash and igneous rock; there are even what look like gray
mountains scattered around. After a while, you make your way to a nearby
cluster of mountains only to discover that the valley between them is
completely full of large *mirrors*. Most of the mirrors seem to be aligned in
a consistent way; perhaps you should head in that direction?

As you move through the valley of mirrors, you find that several of them have
fallen from the large metal frames keeping them in place. The mirrors are
extremely flat and shiny, and many of the fallen mirrors have lodged into the
ash at strange angles. Because the terrain is all one color, it's hard to tell
where it's safe to walk or where you're about to run into a mirror.

You note down the patterns of ash (`.`) and rocks (`#`) that you see as you
walk (your puzzle input); perhaps by carefully analyzing these patterns, you
can figure out where the mirrors are!

For example:

[code]

    #.##..##.
    ..#.##.#.
    ##......#
    ##......#
    ..#.##.#.
    ..##..##.
    #.#.##.#.
    
    #...##..#
    #....#..#
    ..##..###
    #####.##.
    #####.##.
    ..##..###
    #....#..#
    
[/code]

To find the reflection in each pattern, you need to find a perfect reflection
across either a horizontal line between two rows or across a vertical line
between two columns.

In the first pattern, the reflection is across a vertical line between two
columns; arrows on each of the two columns point at the line between the
columns:

[code]

    123456789
        ><   
    #.##..##.
    ..#.##.#.
    ##......#
    ##......#
    ..#.##.#.
    ..##..##.
    #.#.##.#.
        ><   
    123456789
    
[/code]

In this pattern, the line of reflection is the vertical line between columns 5
and 6. Because the vertical line is not perfectly in the middle of the
pattern, part of the pattern (column 1) has nowhere to reflect onto and can be
ignored; every other column has a reflected column within the pattern and must
match exactly: column 2 matches column 9, column 3 matches 8, 4 matches 7, and
5 matches 6.

The second pattern reflects across a horizontal line instead:

[code]

    1 #...##..# 1
    2 #....#..# 2
    3 ..##..### 3
    4v#####.##.v4
    5^#####.##.^5
    6 ..##..### 6
    7 #....#..# 7
    
[/code]

This pattern reflects across the horizontal line between rows 4 and 5. Row 1
would reflect with a hypothetical row 8, but since that's not in the pattern,
row 1 doesn't need to match anything. The remaining rows match: row 2 matches
row 7, row 3 matches row 6, and row 4 matches row 5.

To *summarize* your pattern notes, add up *the number of columns* to the left
of each vertical line of reflection; to that, also add *100 multiplied by the
number of rows* above each horizontal line of reflection. In the above
example, the first pattern's vertical line has `5` columns to its left and the
second pattern's horizontal line has `4` rows above it, a total of `*405*`.

Find the line of reflection in each of the patterns in your notes. *What
number do you get after summarizing all of your notes?*

To begin, [get your puzzle input](13/input).

Answer:

You can also [Shareon
[Twitter](https://twitter.com/intent/tweet?text=%22Point+of+Incidence%22+%2D+Day+13+%2D+Advent+of+Code+2023&url=https%3A%2F%2Fadventofcode%2Ecom%2F2023%2Fday%2F13&related=ericwastl&hashtags=AdventOfCode)
[Mastodon](javascript:void\(0\);)] this puzzle.

