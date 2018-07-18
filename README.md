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

## Instructions
Do ```go build src/gameserver``` to obtain the executable. You can run it just like it is, or use a docker image. 

The docker image copies the executable (assumed to be named ```gameserver```) from the current directory into ```/```. It uses the directory ```./config``` as a bind mount to ```/data``` to load the config, custom programs and store logfiles.
It exposes port ```8080``` to the outside world.
The docker image uses Alpine, you can add custom interpreters or runtimes with ```apk add```, for Java do: ```apk add openjdk8```. To run the image, execute ```docker-compose up```.

## Warning
**DO NOT GIVE USERS ACCESS TO PROGRAMS THAT ALLOW ARBITRARY CODE EXECUTION, SUCH AS SHELLS, INTERACTIVE INTERPRETERS OR OTHER PROGRAMS.
I DO NOT GIVE ANY WARRANTY THAT THIS PROGRAM WORKS AS INTENDED AND
I DO NOT TAKE ANY LIABILITY FOR DAMAGES THAT OCCUR TO YOUR SYSTEM BECAUSE OF ANY REASON WHATSOEVER.**

## TODO
* Apply Go naming conventions
* Add option to set current working directory in config file