# Build the software...

We'll design this software in a maintainable way. It means good practices. It means documentation.
It means code structure.

This is your bible.
https://www.python.org/dev/peps/pep-0020/#id3

## 1. Create the folder architecture

In $GOPATH, create  the mrcsptr/riley folder (or clone it from github).

Then, create two folders: `$GOPATH/bin`, and `$GOPATH/cmd/fight`
_Note: `bin` will contain the software after compilation, and the `config.json` file. `cmd/fight` will contain the source code of the CLI software. It is important to namespace it under a cmd directory as you may need another software attached to this project later._

Create a `$GOPATH/.gitignore` file containing this
```
bin/*
!bin/*.json
```

This will prevent the git repository to store the compiled binary, but will keep the configuration files.
_WARNING: configuration files stored in csv (control source version) *must NEVER* contain secret or specific data, such as password or IPs._

In `$GOPATH/bin`, create a first `config.json` file that contains an empty json object.
In `$GOPATH/cmd/fight`, create a `main.go` file with a simple hello world.

### Libs
- fmt

### Success Criterion

```sh
cd $GOPATH
go build -o bin/fight cmd/fight
bin/fight
# prints hello world
```

## 2. Parse CLI flag

With the standard flag library, parse the CLI to fetch the configuration filepath from the terminal.
Open and display the content of the configuration file in the terminal.
If the config file cannot be loaded, crash with `log.Fatal` and display an explicit error.
_NOTE: a call to os.Exit() also crashes the software, but log displays a message in the stderr output_

### Libs
- fmt
- flag
- io/ioutil
- log

### Success Criterion

The software prints the exact content of the file given as c flag.
If the file does not exist, the software crashes with a pretty error logged.

## 3. Build the configuration

Create a `$GOPATH/cmd/fight/config.go` that declares a `Config` type.
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
