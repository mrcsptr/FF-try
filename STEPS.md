# Build the software...

We'll design this software in a maintainable way. It means good practices. It means documentation.
It means code structure.

This is your bible.
https://www.python.org/dev/peps/pep-0020/#id3

## 1. Create the folder architecture

In $GOPATH, create  the mrcsptr/riley folder (or clone it from github).

_Note: we'll call $PROJECT the `$GOPATH/mrcsptr/riley` directory_

Then, create two folders: `$PROJECT/bin`, and `$PROJECT/cmd/fight`
_Note: `bin` will contain the software after compilation, and the `config.json` file. `cmd/fight` will contain the source code of the CLI software. It is important to namespace it under a cmd directory as you may need another software attached to this project later._

Create a `$PROJECT/.gitignore` file containing this
```
bin/*
!bin/*.json
```

This will prevent the git repository to store the compiled binary, but will keep the configuration files.
_WARNING: configuration files stored in csv (control source version) *must NEVER* contain secret or specific data, such as password or IPs._

In `$PROJECT/bin`, create a first `config.json` file that contains an empty json object.
In `$PROJECT/cmd/fight`, create a `main.go` file with a simple hello world.

### Libs
- fmt

### Success Criterion

```sh
cd $PROJECT
go build -o bin/fight cmd/fight
bin/fight
# prints hello world
```

## 2. Parse CLI flag

With the standard flag library, parse the CLI to fetch the configuration filepath from the terminal.
Open and display the content of the configuration file in the terminal.
If the config file cannot be loaded, crash with `log.Fatal` and display an explicit error.
_Note: a call to os.Exit() also crashes the software, but log displays a message in the stderr output_

### Libs
- fmt
- flag
- io/ioutil
- log

### Success Criterion

The software prints the exact content of the file given as c flag.
If the file does not exist, the software crashes with a pretty error logged.

## 3. Build the configuration

Create a `$PROJECT/cmd/fight/config.go` that declares a `Config` type.
Implement a function that creates this config type and puts the data from the input file into it.
Implement a method that displays the content of the config, in a pretty way.
Change the main implementation to create a new config object from the given filepath, and print the object.

```go

// Config of the software.
type Config struct{
  // declare parameters here
}

// NewConfig loads a new configuration created with the values declared in the json of the filepath.
NewConfig(filepath string) (Config, error) {
  // implementation...
}

// String implements the fmt.Stringer method. It allows, amongst other things, a pretty display of the receiver when passed
// as parameter to fmt.Print functions.
func (c Config) String() string{
  // display the json encoded values of the config.
}
```

### Libs
- encoding/json
- fmt
- flag
- io/ioutil

### Success Criterion

The software parses the content of the configuration located the c flag.
The software crashes if the config is invalid.
If the file does not exist, the software crashes with a pretty error logged.
The exact output can be different from the input, but the values are the same: space or line feed will be ignored, all attributes will be CamelCased, etc...

## 4. Define the Dude domain type

A dude is defined by a name, and a list of previous fights.
Declare a dude type at the domain (project's root) level, with the proper attributes.
The dude type must implement three methods that calculate its grade regarding his history.
To share the maximum amount of code, those three *exported* methods should rely on a unique *non exported method*.
You will need to declare a dude result type, or embed it in the dude type.
Since the results are splitted into 3 distinct types (A, M, D), think on the best primitive to store them.

Once the dude type is declared, you need to instanciate dudes with their files. The configuration gives the location
of the dude directory. Each dude from this directory must be loaded and stored in memory. Taking the content of a file
to instanciate a new dude is called `parsing`. Parse all files within the dude folder to instanciate the dudes.
This parsing must be declared outside of the domain (project's root). create a `$PROJECT/data/filesystem` directory, and
implement the dude parsing in the `dude.go` file. Because the number of dudes and the size of their history is limitless,
use an `io.Reader` instead of `ioutil.ReadFile`.

```go
func AllDudes(location string) ([]riley.Dude, error) {
  // ... load dudes.
}

// String implements the fmt.Stringer interface. It displays the name of the dude, it's current grades, and the number of fights he did.
func (d Dude) String() string{
}
```

Finaly, print all dudes on the terminal after loading them in memory.

_Note: because you have internal dependancies, you need to use the riley package within your source code. The riley package is the name of the `domain` package. Because the principal repository is `moxar/riley`, you need to rename it this way *on your computer*. Otherwise, the imports won't work._

_Note: put some dude files in a `$PROJECT/samples/dudes` repository._

### Libs
- fmt
- strings
- io
- os
- moxar/riley

## 5. Define the Team domain type

A team is composed of several dudes.
Declare a team type at the domain (projet's root) level, with the proper attributes.
A team is composed of a collection of dudes.

Once the team type is declared, you need to instanciate it with its file. The configuration gives the location of the team's directory.
Just as the dudes, each team must be parsed and loaded in memory.

The team type must implement:
- methods to calculate the sum of score of it's dudes, for each position.
- methods to calculate the result of a given position against another team.
- a method to determine, when opposing another team, the score: VTD.

Finaly, print all teams on the terminal after loading them in memory.

_Note: put some team files in a `$PROJECT/samples/teams` repository._

## 6. Putting things together...

Using te configuration, load the two opposing teams and their dudes, oppose them and define the winning team.
Add the result to the dudes of each team for the next encounter.

## 7. SQL

TODO

## 8. HTTP

TODO

## 9. Unit tests

TODO
