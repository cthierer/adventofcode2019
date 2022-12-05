# Rucksack reorganization

Determines overlapping item types in rucksacks, according to
[the prompt](https://adventofcode.com/2022/day/3):

> Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer 
> to different types of items).
>
> The list of items for each rucksack is given as characters all on a single line. A given rucksack 
> always has the same number of items in each of its two compartments, so the first half of the 
> characters represent items in the first compartment, while the second half of the characters 
> represent items in the second compartment.

The file is read in from `stdin` up until `EOF`.

## Sample usage

```bash
go run cmd/day03/main.go < input-day03.txt
```
