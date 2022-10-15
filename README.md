# Data munging with Go

[Coding kata task](http://codekata.com/kata/kata04-data-munging/)

## Part One: Weather Data
In [weather.dat](http://codekata.com/data/04/weather.dat) you’ll find daily weather data for Morristown, NJ for June 2002. Download this text file, then write a program to output the day number (column one) with the **smallest temperature spread** (the maximum temperature is the second column, the minimum the third column).

> No TDD this time ☹

### Pseudocode

- Read values from multidimensional slice of three columns
- Calculate difference between max and min (2nd and 3rd column)
- Compare differences and return the smallest one with the corresponding day

