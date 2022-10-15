# Data munging with Go

[Coding kata task](http://codekata.com/kata/kata04-data-munging/)

---
## Part One: Weather Data
In [weather.dat](http://codekata.com/data/04/weather.dat) you’ll find daily weather data for Morristown, NJ for June 2002. Download this text file, then write a program to output the day number (column one) with the **smallest temperature spread** (the maximum temperature is the second column, the minimum the third column).

> No TDD this time ☹️
### Pseudocode
- Read values from multidimensional slice of three columns
- Calculate difference between max and min (2nd and 3rd column)
- Compare differences and return the smallest one with the corresponding day

---

## Part Two: Soccer League Table
The file [football.dat](http://codekata.com/data/04/football.dat) contains the **league table** for the English Premier League for 2001/2. The columns labeled ‘F’ and ‘A’ contain the total number of **goals scored** for and against each team in that season (so Arsenal scored 79 goals against opponents, and had 36 goals scored against them). Write a program to **print the name of the team with the smallest difference** in ‘for’ and ‘against’ goals.

### Pseudocode
- Refactor previous code to be more generic and reusable
- Previous implementation was too specific to the weather data and read only first three columns
- Modify ParseData to read columns matching the first row of the file `parsedata(twoDimStrSlice[][], sliceOfImportantColumns[]) sliceOfStructs[]`
  - Can I make a dynamic slice of structs? `[]struct{}`?

> Might **really** need TDD this time 🤔

---