# Specifications

This document lists all the possible specifications. It should eventually be
dynamically filed from the code.

IDs are assigned as code to test for each specification is written.

> [!NOTE]
> "option" is prefered to "flag" as suggested by POSIX.1-2024.

| id  | regex / name          | example                                            | expected parsing                      | comments                        |
| --- | ---                   | ---                                                | ---                                   | ---                             |
|     | `-\w`                 | `-x`                                               | `{ "x": true }`                       |                                 |
|     | `(-\w)( \1)*`         | `-x -x -x`                                         | `{ "x": 3 }`                          |                                 |
|     | `-(\w)(\1)*`          | `-xxx`                                             | `{ "x": 3 }`                          |                                 |
|     | `-\w+`                | `-xyz`                                             | `{ "x": true, "y": true, "z": true }` |                                 |
|     | `-\w.+`               | `-xabc`                                            | `{ "x": "abc" }`                      |                                 |
|     | `-\w=.+`              | `-x=abc`                                           | `{ "x": "abc" }`                      |                                 |
|     | `-\w .+`              | `-x abc`                                           | `{ "x": "abc" }`                      |                                 |
|     | `(-\w).+ (\1.+)+`     | `-xa -xb -xc`                                      | `{ "x": ["a", "b", "c" }`             |                                 |
|     | `(-\w)=.+ (\1=.+)+`   | `-x=a -x=b -x=c`                                   | `{ "x": ["a", "b", "c" }`             |                                 |
|     | `(-\w) .+ (\1 .+)+`   | `-x a -x b -x c`                                   | `{ "x": ["a", "b", "c" }`             |                                 |
|     | `-\w+`                | `-xyz`                                             | `{ "xyz": true }`                     |                                 |
|     | `-\w+.+`              | `-xyzabc`                                          | `{ "xyz": "abc" }`                    |                                 |
|     | `-\w+=.+`             | `-xyz=abc`                                         | `{ "xyz": "abc" }`                    |                                 |
|     | `-\w+ .+`             | `-xyz abc`                                         | `{ "xyz": "abc" }`                    |                                 |
|     | `(-\w+).+ (\1.+)+`    | `-xyza -xyzb -xyzc`                                | `{ "xyz": ["a", "b", "c" }`           |                                 |
|     | `(-\w+)=.+ (\1=.+)+`  | `-xyz=a -xyz=b -xyz=c`                             | `{ "xyz": ["a", "b", "c" }`           |                                 |
|     | `(-\w+) .+ (\1 .+)+`  | `-xyz a -xyz b -xyz c`                             | `{ "xyz": ["a", "b", "c" }`           |                                 |
|     | `-\d`                 | `-1`                                               | `{ "1": true }`                       |                                 |
|     | `(-\d)( \1)+`         | `-1 -1 -1`                                         | `{ "1": 3 }`                          |                                 |
|     | `-(\d)(\1)+`          | `-111`                                             | `{ "1": 3 }`                          |                                 |
|     | `-\d+`                | `-123`                                             | `{ "1": true, "2": true, "3": true }` |                                 |
|     | `-\d.+`               | `-1abc`                                            | `{ "1": "abc" }`                      |                                 |
|     | `-\d=.+`              | `-1=abc`                                           | `{ "1": "abc" }`                      |                                 |
|     | `-\d .+`              | `-1 abc`                                           | `{ "1": "abc" }`                      |                                 |
|     | `(-\d).+ (\1.+)+`     | `-1a -1b -1c`                                      | `{ "1": ["a", "b", "c" }`             |                                 |
|     | `(-\d)=.+ (\1=.+)+`   | `-1=a -1=b -1=c`                                   | `{ "1": ["a", "b", "c" }`             |                                 |
|     | `(-\d) .+ (\1 .+)+`   | `-1 a -1 b -1 c`                                   | `{ "1": ["a", "b", "c" }`             |                                 |
|     | `-\d+`                | `-123`                                             | `{ "123": true }`                     |                                 |
|     | `-\d+.+`              | `-123abc`                                          | `{ "123": "abc" }`                    |                                 |
|     | `-\d+=.+`             | `-123=abc`                                         | `{ "123": "abc" }`                    |                                 |
|     | `-\d+ .+`             | `-123 abc`                                         | `{ "123": "abc" }`                    |                                 |
|     | `(-\d+).+ (\1.+)+`    | `-123a -123b -123c`                                | `{ "123": ["a", "b", "c" }`           |                                 |
|     | `(-\d+)=.+ (\1=.+)+`  | `-123=a -123=b -123=c`                             | `{ "123": ["a", "b", "c" }`           |                                 |
|     | `(-\d+) .+ (\1 .+)+`  | `-123 a -123 b -123 c`                             | `{ "123": ["a", "b", "c" }`           |                                 |
|     | `--`                  |                                                    |                                       | separate options from arguments |
|     | `-`                   |                                                    |                                       | signifies `STDIN`               |
|     | `-h`                  |                                                    | display help                          |                                 |
|     | `-h`                  |                                                    | display short help                    |                                 |
|     | `--help`              |                                                    | display help                          |                                 |
|     | `--help`              |                                                    | display long help                     |                                 |
|     | `.+( .+)*`            | `arg1 arg2 arg3`                                   | { "args": ["arg1", "arg2", "arg3"] }  |                                 |
|     | `..+`                 | `i` or `install`                                   | { "args": ["install"] }               | arguments / subcommands aliases |
|     | value validation      |                                                    |                                       |                                 |
|     | value type validation |                                                    |                                       |                                 |
|     | subcommands           |                                                    |                                       |                                 |
|     | global options        |                                                    |                                       |                                 |
|     | option position 1     | `cli [global options] verb [local options] [args]` |                                       |                                 |
|     | option position 2     | `cli verb [global + local options] [args]`         |                                       |                                 |
|     | suggestions           |                                                    |                                       |                                 |
|     | deprecations          |                                                    |                                       |                                 |

# Sources

- https://pubs.opengroup.org/onlinepubs/9799919799/basedefs/V1_chap12.html
- https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html
