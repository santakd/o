# o [![Build Status](https://travis-ci.com/xyproto/o.svg?branch=master)](https://travis-ci.com/xyproto/o) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/o)](https://goreportcard.com/report/github.com/xyproto/o) [![License](https://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/o/master/LICENSE)

`o` is a text editor that is limited to the VT100 standard. It launches instantly and syntax highlights code instantly.

It's a good fit for:
* Editing git commit messages (using `EDITOR=o git commit`).
* Editing Markdown.
* Quickly editing a file when programming.

For a more feature complete editor that is also written in Go, check out [micro](https://github.com/zyedidia/micro).

## Packaging status

[![Packaging status](https://repology.org/badge/vertical-allrepos/o-editor.svg)](https://repology.org/project/o-editor/versions) [![Packaging status](https://repology.org/badge/vertical-allrepos/o.svg)](https://repology.org/project/o/versions)

## Quick start

You can install `o` with Go 1.10 or later:

    go get -u github.com/xyproto/o

## Setting `o` as the default editor for `git`

To set:

    git config --global core.editor o

To unset:

    git config --global --unset core.editor

## Features that are unique to `o`

These features are unique to `o`, as far as I am aware:

* Smart cursor movement, trying to maintain the X position when moving up and down.
* Press `ctrl-v` once to paste one line, press `ctrl-v` again to paste the rest.
* Press `ctrl-c` once to copy one line, press `ctrl-v` again to copy until the next blank line.
* Open or close a portal with `ctrl-r`. If a portal is open, paste across files (or within the same file) with `ctrl-v`.
* Build code with `ctrl-space` and format code with `ctrl-w`, for a wide range of programming languages (suitable for smaller projects and quick edits).
* Cycle git rebase keywords with `ctrl-r`, when in an interactive git rebase session.
* Multiple uses of `ctrl-l`: Enter a number to jump to a line or press `return` to jump to the top. Press `ctrl-l` and `return` again to jump to the end.

## Other features and limitations

* Loads up instantly.
* Reasonable executable size (around 500k, when built with `gccgo` + `upx`).
* Configuration-free, for better and for worse.
* Provides syntax highlighting for Go, C++, Markdown, Bash and several other languages. There is generic syntax highlighting.
* The syntax highlighting is instant.
* The [`NO_COLOR`](https://no-color.org) environment variable can be set to disable all colors.
* Rainbow parentheses makes lines with many parentheses easier to read.
* Limited to the VT100 standard, so hotkeys like `ctrl-a` and `ctrl-e` must be used instead of `Home` and `End`.
* Compiles with either `go` or `gccgo`.
* Tested with `st`, `urxvt`, `konsole` and `xfce4-terminal`.
* Tested on Arch Linux, Debian and FreeBSD.
* Loads faster than both `vim` and `emacs`.
* Never asks before saving or quitting. Be careful!
* Will strip trailing whitespace whenever it can.
* Must be given a filename at start.
* May provide smart indentation.
* Requires `/dev/tty` to be available.
* `xclip` (for X) or `wl-clipboard` (for Wayland) must be installed if the system clipboard is to be used.
* May take a line number as the second argument, with an optional `+` prefix.
* All the text will be red if a loaded file is read-only.
* If the filename is `COMMIT_EDITMSG`, the look and feel will be adjusted for git commit messages.
* Supports `UTF-8`, but some runes may be displayed incorrectly.
* Can render text to PDF.
* Only UNIX-style line endings are supported (`\n`).
* Will convert DOS/Windows line endings (`\r\n`) to UNIX line endings (just `\n`), whenever possible.
* Will replace non-breaking space (`0xc2 0xa0`) with a regular space (`0x20`) whenever possible.
* Will jump to the last visited line when opening a recent file.
* If tab completion on the terminal went wrong and you are trying to open a `main.` file that does not exist, but `main.cpp` and `main.o` does exists, then `main.cpp` will be opened.
* If interactive rebase is launched with `git rebase -i`, then either `ctrl-r` or `ctrl-w` will cycle the keywords for the current line (`fixup`, `drop`, `edit` etc).
* When editing Markdown, checkboxes can be toggled with `ctrl-w`.
* `o` makes it easy to write `"Hello, World!"` in one of many languages, then simply compile it with `ctrl-space`. The idea is that it should be very quick to write and compile a short program, without any hassle.
* `ctrl-t` can toggle between a C++ header and source file, if searching for the file in the parent directories is quick enough.
* If the editor executable is named to `red`, the default theme will be red/white/gray.

## Known bugs

* Some unicode runes may disrupt the text flow. This is generally not a problem for editing code and configuration files, but may be an issue when editing files that contains text in many languages.
* The cursor may be misplaced when moving around on lines longer than the terminal width and then inserting or deleting text.
* The smart indentation is not always smart.

## Hotkeys

* `ctrl-s` - Save.
* `ctrl-q` - Quit.
* `ctrl-r` - Open or close a portal. Text can be pasted from the portal into another file with `ctrl-v`.
             For "git interactive rebase" mode, cycle the rebase keywords.
* `ctrl-w` - Format the current file (see the table below).
* `ctrl-a` - Go to start of text, then start of line and then to the previous line.
* `ctrl-e` - Go to end of line and then to the next line.
* `ctrl-p` - Scroll up 10 lines, or go to the previous match if a search is active.
* `ctrl-n` - Scroll down 10 lines, or go to the next match if a search is active.
* `ctrl-k` - Delete characters to the end of the line, then delete the line.
* `ctrl-g` - Toggle a status line at the bottom for displaying: filename, line, column, Unicode number and word count.
* `ctrl-d` - Delete a single character.
* `ctrl-t` - Render the current document to a PDF file. For C++, toggle between header and implementation files.
* `ctrl-o` - Open a command menu with actions that can be performed. The first menu item is always `Save and quit`.
* `ctrl-x` - Cut the current line. Press twice to cut a block of text (to the next blank line).
* `ctrl-c` - Copy one line. Press twice to copy a block of text.
* `ctrl-v` - Paste one trimmed line. Press twice to paste multiple untrimmed lines.
* `ctrl-space` - Build (see table below)
* `ctrl-j` - Join lines (or jump to the bookmark, if set).
* `ctrl-u` - Undo (`ctrl-z` is also possible, but may background the application).
* `ctrl-l` - Jump to a specific line number. Follows by `return` to jump to the top. If at the top, press `return` to jump to the bottom.
* `ctrl-f` - Search for a string. The search wraps around and is case sensitive.
* `esc` - Redraw the screen and clear the last search.
* `ctrl-b` - Toggle a bookmark for the current line, or if set: jump to a bookmark on a different line.
* `ctrl-\` - Comment in or out a block of code.
* `ctrl-~` - Jump to a matching parenthesis.

## Updating PKGBUILD files

When editing `PKGBUILD` files, it is possible to press `ctrl-w` to update the `pkgver=` and `source=` fields, by a combination of guesswork and online searching.

[`guessica`](https://github.com/xyproto/guessica) is the utility that is used for performing the guesswork, and must be installed for this feature to work.

## Build and format

* At the press of `ctrl-space`, `o` will try to build or export the current file.
* At the press of `ctrl-w`, `o` will try to format the current file, in an opinionated way.

| Programming language                            | File extensions                                           | Jump to error | Build command                                     | Format command ($filename is a temporary file)                                                                 |
|-------------------------------------------------|-----------------------------------------------------------|---------------|---------------------------------------------------|----------------------------------------------------------------------------------------------------------------|
| Go                                              | `.go`                                                     | yes           | `go build`                                        | `goimports -w -- $filename`                                                                                    |
| C++                                             | `.cpp`, `.cc`, `.cxx`, `.h`, `.hpp`, `.c++`, `.h++`, `.c` | yes           | `cxx`                                             | `clang-format -fallback-style=WebKit -style=file -i -- $filename`                                              |
| Rust                                            | `.rs`                                                     | yes           | `rustc $filename`                                 | `rustfmt $filename`                                                                                            |
| Rust, if `Cargo.toml` or `../Cargo.toml` exists | `.rs`                                                     | yes           | `cargo build`                                     | `rustfmt $filename`                                                                                            |
| Zig                                             | `.zig`                                                    | yes           | `zig build`                                       | `zig fmt $filename`                                                                                            |
| V                                               | `.v`                                                      | yes           | `v build`                                         | `v fmt $filename`                                                                                              |
| Haskell                                         | `.hs`                                                     | yes           | `ghc -dynamic $filename`                          | `brittany --write-mode=inplace $filename`                                                                      |
| Python                                          | `.py`                                                     | yes           | `python -m py_compile $filename`                  | `autopep8 -i --maxline-length 120 $filename`                                                                   |
| Crystal                                         | `.cr`                                                     | yes           | `crystal build --no-color $filename`              | `crystal tool format $filename`                                                                                |
| Kotlin                                          | `.kt`                                                     | yes           | `kotlinc $filename -include-runtime -d $name.jar` | `ktlint`                                                                                                       |
| Java                                            | `.java`                                                   | yes           | `javac` + `jar`, see details below                | `google-java-format -i $filename`                                                                              |
| Lua                                             | `.lua`                                                    | yes           | `luac`                                            | `lua-format -i --no-keep-simple-function-one-line --column-limit=120 --indent-width=2 --no-use-tab $filename`  |
| Object Pascal                                   | `.pas`, `.pp`, `.lpr`                                     | yes           | `fpc`                                             | WIP                                                                                                            |
| Nim                                             | `.nim`                                                    | WIP           | `nim c`                                           | WIP                                                                                                            |

* `o` will try to jump to the location where the error is and otherwise display `Success`.
* For regular text files, `ctrl-w` will word wrap the lines to a length of 99.

CXX can be downloaded here: [GitHub project page for CXX](https://github.com/xyproto/cxx).

| File type | File extensions  | Build or export command                                           |
|-----------|------------------|-------------------------------------------------------------------|
| AsciiDoc  | `.adoc`          | `asciidoctor -b manpage` (writes to `out.1`)                      |
| scdoc     | `.scd`, `.scdoc` | `scdoc` (writes to `out.1`)                                       |
| Markdown  | `.md`            | `pandoc -N --toc -V geometry:a4paper` (writes to `$filename.pdf`) |

If [`guessica`](https://github.com/xyproto/guessica) is installed, `PKGBUILD` files will be updated at the press of `ctrl-w`. The `guessica` utility tries to guess the latest project version, tag and git commit hash for a `PKGBUILD` file and may or may not succeed.

## Manual installation

On Linux:

    git clone https://github.com/xyproto/o
    cd o
    go build -mod=vendor
    sudo install -Dm755 o /usr/bin/o
    gzip o.1
    sudo install -Dm644 o.1.gz /usr/share/man/man1/o.1.gz

## Dependencies

C++

* For building code with `ctrl-space`, [`cxx`](https://github.com/xyproto/cxx) must be installed.
* For formatting code with `ctrl-w`, `clang-format` must be installed.

Go

* For building code with `ctrl-space`, The `go` compiler must be installed.
* For formatting code with `ctrl-w`, [`goimports`](https://godoc.org/golang.org/x/tools/cmd/goimports) must be installed.

Zig

* For building and formatting Zig code, only the `zig` command is needed.

V

* For building and formatting V code, only the `v` command is needed.

Rust

* For building code with `ctrl-space`, `Cargo.toml` must exist and `cargo` must be installed.
* For formatting code with `ctrl-w`, `rustfmt` must be installed.

Haskell

* For building the current file with `ctrl-space`, the `ghc` compiler must be installed.
* For formatting code with `ctrl-w`, [`brittany`](https://github.com/lspitzner/brittany) must be installed.

Python

* `ctrl-space` only checks the syntax, without executing. This only requires `python` to be available.
* For formatting the code with `ctrl-w`, `autopep8` must be installed.

Crystal

* For building and formatting Crystal code, only the `crystal` command is needed.

Kotlin

* For building code with `ctrl-space`, `kotlinc` must be installed. A `.jar` file is created if the compilation succeeded.
* For formatting code with `ctrl-w`, `ktlint` must be installed.

Java

* For building code with `ctrl-space`, `javac` and `jar` must be installed. A `.jar` file is created if the compilation succeeded.
* For formatting code with `ctrl-w`, `google-java-format` must be installed.

JSON

* The JSON formatter is built-in.

## A note on Java

Since `kotlinc $filename -include-runtime -d` builds to a `.jar`, I though I should do the same for Java. The idea is to easily build a single or a small collection of `.java` files, where one of the file has a `main` function.

If you know an easier way to build a `.jar` file from `*.java` without using something like gradle, please let me know by submitting a pull request. This is pretty verbose...

```sh
javaFiles=$(find . -type f -name '*.java')
for f in $javaFiles; do
  grep -q 'static void main' "$f" && mainJavaFile="$f"
done
className=$(grep -oP '(?<=class )[A-Z]+[a-z,A-Z,0-9]*' "$mainJavaFile" | head -1)
packageName=$(grep -oP '(?<=package )[a-z,A-Z,0-9,.]*' "$mainJavaFile" | head -1)
if [[ $packageName != "" ]]; then
  packageName="$packageName."
fi
mkdir -p _o_build/META-INF
javac -d _o_build $javaFiles
cd _o_build
echo "Main-Class: $packageName$className" > META-INF/MANIFEST.MF
classFiles=$(find . -type f -name '*.class')
jar cmf META-INF/MANIFEST.MF ../main.jar $classFiles
cd ..
rm -rf _o_build
```

## List of optional runtime dependencies

* `go` / `golang`
* `goimports`
* [`cxx`](https://github.com/xyproto/cxx)
* `g++` / `base-devel`
* `clang-format`
* `rustc`
* `rustfmt`
* `cargo`
* `zig`
* `v`
* `ghc`
* `brittany`
* `python`
* `autopep8`
* `kotlin`
* `ktlint`
* `javac`
* `jar`
* `google-java-format`

## Size

* The `o` executable is only **508k** when built with GCC 9.3 (for 64-bit Linux) and compressed with `upx`.
* This isn't as small as [e3](https://sites.google.com/site/e3editor/), an editor written in assembly (which is **234k**), but it's reasonably lean.

One way of building with `gccgo` and `upx`:

    go build -mod=vendor -gccgoflags '-Os -s' && upx o

It's **5.2M** when built with Go 1.14 and no particular build flags are given.

## Jumping to a specific line when opening a file

These four ways of opening `file.txt` at line `7` are supported:

* `o file.txt 7`
* `o file.txt +7`
* `o file.txt:7`
* `o file.txt+7`

This also means that filenames containing `+` or `:`, and then followed by a number, are not supported.

## The first keypress

If the very first keypress after starting is `G` or `/`, it will trigger the following vi-compatible behavior:

* `/` - enter search-mode (same as when pressing `ctrl-f`)
* `G` - go to the end of the file

The reason for adding these is to make using `o` easier to use for long-time vi/vim/neovim users.

## Spinner

When loading files that are large or from a slow disk, an animated spinner will appear. The loading operation can be interrupted by pressing `esc`, `q` or `ctrl-q`.

![progress](img/progress.gif)

## Find and open

This shell function works in `zsh` and `bash` and may be useful for both searching for and opening a file at the given line number (works best if there is only one matching file, if not it will open several files in succession):

```bash
fo() { find . -type f -wholename "*$1" -exec /usr/bin/o {} $2 \;; }
```

Example use:

```sh
fo somefile.cpp 123
```

## Easter eggs

Pressing `ctrl-space` twice will render Markdown files to PDF using `pandoc` (as opposed to `ctrl-r`, which will save the text directly to a PDF, without using `pandoc`).

## Suggested settings

### Konsole

* Try the `Breeze` color scheme (but with a black background) and the wonderful (and open source) `JetBrains Mono NL` font.
* Untick the `Flow control` option in the profile settings, to ensure that `ctrl-s` will never freeze the terminal.

## General info

* Version: 2.32.5
* License: 3-clause BSD
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
