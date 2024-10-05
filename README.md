# Filebased

**FileBased** is a file-based router for the [Echo](https://echo.labstack.com/) web framework written in Go.
This router allows you to define routes through structured files,
making it easy to manage and modify large routing configurations without digging into code.
It's designed for flexibility and scalability,
especially for teams or projects that prefer keeping routing logic out of the main application code.

## Installation

```sh
go install github.com/ChaotenHG/filebased/cmd/filebased@latest
```

## Usage

To get started, you can either clone the [template project](https://github.com/ChaotenHG/filebased-template/) or create a new one from scratch. The recommended file structure can be found in the provided example.

Once everything is set up, simply run the following command to generate the necessary files:

```sh
filebased
```

To enable the routes, **make sure to call** the `registerRoutes(e *echo.Echo)` function in your project.

### Command options

```
  -h, --help             Show usage summary
  -i, --input string     directory of your route files (default "./routes")
  -o, --output string    Output file name (without .go) (default "routes")
  -p, --package string   the name of your package where the server is defined (default "main")
```

### File Structure

```
routes
├── article
│   └── :user
│       └── id:.go  -> example.com/article/bob/123
├── user
│   ├── new.go      -> example.com/user/new
│   └── user:.go    -> example.com/user/bob
├── hello.go        -> example.com/hello
└── world.go        -> example.com/world
```

To register a path parameters append the file name with Colon `user:.go` or prefix a directory with a Colon to do the same `:article_id/`.

## Caveats

- This tool has to be extcute everytime you changed a file in the routes directory (use [air](https://github.com/air-verse/air))
- This tool could merge only files from one directory using the specified package name.
- This tool does not check for syntax errors, so make sure that your code is syntactically correct before merging the files.
- The tool does not check for **yet** conflicts between files, so you need to make sure that there are no conflicts manually.

## Credits

A big thank you to **mrsombre** for their incredible work on the [original project](https://github.com/mrsombre/codingame-golang-merger),
which served as the foundation for this tool. Without their contributions,
this project wouldn’t have been possible.

## Contributing

I welcome and appreciate contributions of all kinds!
If you have any ideas, suggestions, or find any issues with the project,
please don't hesitate to submit an Issue or to create a Pull Request

This project follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

Thank you!
