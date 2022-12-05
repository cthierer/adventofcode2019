# Calorie reporter

Reports information about calories, read in from a file matching the format described in [the
prompt](https://adventofcode.com/2022/day/1):

> The Elves take turns writing down the number of Calories contained by the various meals, snacks, 
> rations, etc. that they've brought with them, one item per line. Each Elf separates their own 
> inventory from the previous Elf's inventory (if any) by a blank line.

The file is read in from `stdin` up until `EOF`.

## Sample usage

```bash
go run cmd/day01/main.go < input-day01.txt
```
