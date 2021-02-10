<h1 align="center">zlang-interpreter</h1>
<p align="center">
<img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/zhooda/zlang-interpreter/Go">
<img alt="go report card" src="https://goreportcard.com/badge/github.com/zhooda/zlang-interpreter">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/zhooda/zlang-interpreter">
</p>

<div style="margin-bottom: 2%"></div>

A simple interpreter written in Go. 

<div style="margin-bottom: 2%"></div>

## Getting Started

These instructions *should* get you a copy of the project up and running
on your local machine. There are no deployment strategies right now `:~)`

### Prerequisites

To run zlang-interpreter you will need:

- [golang](https://golang.org/) v1.15+

On linux you will also need `build-essential` to be able to use the `make` command

### Installation

_Clone the repository or checkout with subversion to get a copy of zlang-interpreter_

<details open>
<summary><b>Linux/macOS</b></summary>
<br>

```bash
$ git clone https://github.com/zhooda/zlang-interpreter
$ cd zlang-interpreter
$ make
```
</details>

<details>
<summary><b>Windows</b></summary>
<br>

```powershell
PS> git clone https://github.com/zhooda/zlang-interpreter
PS> cd .\zlang-interpreter
PS> go build -v -o .\bin\z.exe main.go
```

When running abc2 on windows using the commands outlined below, replace `./bin/z` with `.\bin\z.exe` and you'll be good to go :)

</details>

---

### Usage

After you've compiled the program with the instructions above, you can start the repl/shell or you can run a program from a .z file!

File:
```
$ ./bin/z example.z
3 is greater than 2 
hello bob 
length of arr: 5. 
arr: [1, 2, 3, 4, 5]. 
arr is of type: ARRAY 

BOOLEAN 
STRING 
INTEGER 
bruh moment: wrong number of arguments. got=0, want=1
```

REPL:

```
$ ./bin/z
z 0.1.0 (v0.1.0:0x000047, 02/07/2020)
[go version go1.15.8]

hello zeeshanhooda, type some commands
▷ let x = 1 * 2 * 3 / 4 * 5 + 6 - 7
▷ x * y / 2 + 3 * 8 - 123
bruh moment: identifier not found: y
▷ let y = 5
▷ x * y / 2 + 3 * 8 - 123
-89
▷ true == false
false
▷ let x 12 * 3
we ran into some issues here!
 parser errors:
        expected next token to be =, got INT instead
```

## License
zlang-interpreter uses the [MIT](https://choosealicense.com/licenses/mit/) license `:~)`
