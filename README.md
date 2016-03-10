## PoC Raw Printing of a ZPL string via Win32 system calls.

Prints to the system default printer.

Binary available here: https://github.com/rjocoleman/raw-print-go-poc/releases/tag/0.0.1

Usage - double click or run `raw-print.exe` from the command line (best for stdout and stderr).

Windows only.
Printers that support ZPL only (though the raw string could be changed to anything).


## Build from source (requires Go)

```
$ go build raw-print.go
$ raw-print.exe

```
