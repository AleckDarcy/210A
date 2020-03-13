# Pre-requisite
## Install Golang

## Git clone the code
Make sure the path is:
```shell script
$GOPATH\src\github.com\AleckDarcy\210A
```

## Install & execute gradle
```shell script
gradle goBuild
```

## Install antlr4 (optional)
```shell script
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.2-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'
```

## Generate code for grammar (optional)
```shell script
antlr4 -Dlanguage=Go -package grammar zz/grammar/ZZ.g4
```

# Run the code
## Compile ZZ
### Windows
```shell script
cd $GOPATH\src\github.com\AleckDarcy\210A
go build -o main.exe
```

### Linux
```shell script
cd $GOPATH/src/github.com/AleckDarcy/210A
go build -o main
```

## Run ZZ
Now we demonstrate how to run ZZ in two different language types. The examples are run under Linux OS.

Input ($GOPATH/src/github.com/AleckDarcy/210A/example.zz) is a .zz file with some basic matrix operations.

Output ($GOPATH/src/github.com/AleckDarcy/210A/output) is a directory storing intermediate .go file and executable binary files.

### Interpreted Language
Execute:
```shell script
./main -f example.zz
```

The stdout will be:
```shell script
arguments: [-t i] [-f example.zz]
reading .zz file
transforming code
writing code
executing ZZ as an Interpreted Language
output:
[
        [5 11 17],
        [11 25 39],
        [17 39 61]
]
[
        [5 11 17],
        [11 25 39],
        [17 39 61]
]

finished
```

### Complied Language
Execute:
```shell script
./main -f example.zz -t c
```

The stdout will be:
```shell script
arguments: [-t c] [-f example.zz]
reading .zz file
transforming code
writing code
compiling ZZ as a Compiled Language
executable file: output/main
finished
```

Then we execute output/main:
```shell script
./output/main
```

The stdout will be:
```shell script
[
        [5 11 17],
        [11 25 39],
        [17 39 61]
]
[
        [5 11 17],
        [11 25 39],
        [17 39 61]
]

```