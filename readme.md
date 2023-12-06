# Advent of Code

My practice for [Advent of Code](https://adventofcode.com/). I am starting this for the year 2020.

## Getting the session cookie

* Log into Advent of Code at adventofcode.com as normal with preferred identity provider (I usually prefer GitHub)
* with the browser network tab open, make a request to an input at `/YEAR/day/DAY/input` (e.g. `adventofcode.com/2020/day/1/input`)
* take note of the cookie sent in the request headers and copy the __entire__ value
* paste that value in `get_data` for the var `COOKIE`

## Getting question and input

* run `get_data` at the root level with year and day desired

```sh
$ ./get_data 2023 1
```
