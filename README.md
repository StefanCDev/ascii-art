ASCII Art Generator
This Go project converts a string input into a graphical representation using ASCII characters. The program takes any string as input and displays it in an ASCII art form.

Project Overview
The ASCII art generator translates strings into a visual representation using ASCII characters. It is capable of handling numbers, letters, spaces, special characters, and newline characters.

Example Input and Output
Input:
"Hello"

Output:

 _    _          _    _          _
| |  | |        | |  | |        | |
| |__| |   ___  | |  | |   ___  | |
|  __  |  / _ \ | |  | |  / _ \ | |
| |  | | |  __/ | |  | | | (_) || |
|_|  |_|  \___| |_|  |_|  \___/ |_|
Input:
"Hello There"

Output:

 _    _          _   _               _______   _
| |  | |        | | | |             |__   __| | |
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___|
Features
Converts any given string into ASCII art representation.
Handles various characters, including alphabets, digits, special characters, and newline (\n).
Designed to run from the command line.
Installation
Clone the repository:

git clone https://learn.01founders.co/git/lotfi256/ascii-art
Navigate to the project directory:

cd ascii-art
Run the program:

go run .
Usage
The program accepts a string as input and prints its ASCII art representation. For example:

To print "Hello", run:


go run . "Hello"
For multiline input (e.g., "Hello\nThere"), run:


go run . "Hello\nThere"
Allowed Packages
Only the standard Go packages are allowed for this project.

Key Learning Points
Go file system (fs) API
Data manipulation and string handling
Command-line tool development in Go
Troubleshooting
If you encounter any issues, feel free to submit a bug report or propose changes by opening an issue in the repository.
