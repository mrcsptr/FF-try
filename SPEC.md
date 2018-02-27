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
  * position of a dude depends on the position on the file. Position blocks must be separated by blank lines.
  * blank within a line are ignored
  * it might be important to take into consideration the general team config: for example (PSG systematically uses a 4-3-3). Of course configurations exceeding 3 lines would be merged in a 3-line config (eg a 4-2-3-1 would be merged in 4-5-1). That implies that lines can contain from 1 to 5 players, which could also be taken into account (1 attacker should be in trouble facing 5 defenders)
```
teams/
  blue.txt
  red.txt
  green.txt
```

```
# This line is a comment. It shows the global position of dudes within the blue team.
# Attack - first block
Johny
Manuel
# Middle - second block
Keyran
# Defense - third block
Van
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
# this line is a comment. It show the history of matches for Johny.
A V # A for attack, V for victory
M D # M for middle, D for defeat
D T # D for defense, T for tie
```

## Behaviour

* The software exits with code 1 and a nice error message if
  * the configuration is invalid
  * a team or a dude cannot be reached
  * any processing fails
  * more than 3 blocks or less than 1 block is defined within each team

* process position's grade A, M and D for each dude: 
  * to calculate the grade of a dude on a given position, take it's historic. A victory at the position earns 3 points, a tie earns 2, a defeat earns 1.
  * oppose the sum of grades A of team 1 (A1) to grades D of team 2 (D2). Same with A2 vs D1, and M1 vs M2
  * In order to win, a team must collect two position vicory, or one and no defeat. (VTT, VVT, VVD) If a team wins, the other looses. Otherwise, it's a tie.
  * After the match is played, the software registers the score on each player's history. The V/T/D of a dude depends on the match's result, not his personnal score.
