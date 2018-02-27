# SPECS

- The sofware is a CLI
- The sofwares uses yaml configuration
- The software uses data stored in filesystem

## CLI

* flag -c (default: config.json) defines the location of the configuration file

## Configuration

* uses json format
* defines the location of the directory containing the list of teams
* defines the location of the directory containing the list of dudes
* define the name of the first team
* define the name of the second team

## DATA

* one directory for the teams
* one file per team, containing the list of dudes:
  * identified by name
  * one dude per row
  * character # indicates comment
  * blank lines are ignored
  * blank within a line are ignored

```
teams/
  blues
  red
  green
```

```
johny
manuel
```

* one directory for the dudes
* one file per dude
  * one combat log per line
  * each line contains the position (case sensitive) and score
  * character # indicates comment
  * blank lines are ignored
  * blank within a line are ignored

```
dudes/
  johny.txt
  manuel.txt
```

```
# this line is a comment
A V # A for attack, V for victory
M D # M for middle, D for defeat
D T # D for defense, T for tie
```

## Behaviour

* The software exits with code 1 and a nice error message if
  * the configuration is invalid
  * a team or a dude cannot be reached
  * any processing fails

* process position's grade A, M and D for each dude: 
  * to calculate the grade of a dude on a given position, take it's historic. A victory at the position earns 3 points, a tie earns 2, a defeat earns 1.
  * oppose the sum of grades A of team 1 (A1) to grades D of team 2 (D2). Same with A2 vs D1, and M1 vs M2
  * winning a contest earns 1 point, a tie earns none, a defeat costs 1 point. The team with the most points wins the match.
  * After the match is played, the software registers the score on each player's history.
