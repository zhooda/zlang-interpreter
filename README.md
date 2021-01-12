<h1 align="center">zlang-interpreter</h1>
<p align="center"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/zhooda/zlang-interpreter/Go">
<img alt="go report card" src="https://goreportcard.com/badge/github.com/zhooda/zlang-interpreter"></p>

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

After you've compiled the program with the instructions above, you can start the repl/shell with the following command:

```bash
$ ./bin/z
hello bob, welcome to z
type some commands
z>
```

## License
zlang-interpreter uses the [MIT](https://choosealicense.com/licenses/mit/) license `:~)`
