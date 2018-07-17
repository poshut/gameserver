# gameserver
## General
gameserver is a program which you can use to remotely execute predefined programs. You can connect to it via ```telnet```.

## Why?
We created some games in our AP Computer Programming class. 
Because they were scattered among many different machines,
I wanted to have something so you could connect to a server from a shell, choose the game you want to play, play a match and then disconnect.
I heard about the cool concurrency features that Go offers and created this as a personal project.

## Configuration
The default separator is , (comma), but you can use your own with the ```-s``` option.
A config file is made of lines which each represent a game option when you connect onto the server.
Each lines must have at least two parameters, separated by the separator.
The first parameter is the display name of the option. You see only this when you log in.
The second parameter is the program to execute when this option is selected. A program in the ```$PATH``` only needs its name to be written, for any other program, an absolute path has to be given.
Any further parameters are arguments to the program.

## Warning
**DO NOT GIVE USERS ACCESS TO PROGRAMS THAT ALLOW ARBITRARY CODE EXECUTION, SUCH AS SHELLS, INTERACTIVE INTERPRETERS OR OTHER PROGRAMS.
I DO NOT GIVE ANY WARRANTY THAT THIS PROGRAM WORKS AS INTENDED AND
I DO NOT TAKE ANY LIABILITY FOR DAMAGES THAT OCCUR TO YOUR SYSTEM BECAUSE OF ANY REASON WHATSOEVER.**

## TODO
* Add long versions for log file and separator option
* Apply Go naming conventions
* Use better example config file
