# CIC Website Project

The repo for the Calling In Computing website project!

# Initial Installations

## Golang

You will need to install the go programming language to build and run this project.
Here's the [website](https://go.dev/) for info on how to do that for your operating system.

## The Project

Clone this repo to a place you like on your computer.

## Air

Now that you have go installed, you need to install a tool called `air`. It allows for
hot-reloading of the project for a better workflow. In your terminal, run:
```sh
go install github.com/cosmtrek/air@latest
```
This should install `air` globally to your computer.

## Project Dependencies

In your terminal, you'll need to run this command to grab the packages that the
project will use to run the server:
```sh
go mod tidy
```

Now you should be all setup to run the project!

# How do I run this?

In VSCode (or whatever editor you are using) open a terminal. Make sure you
are in the root folder of the project, and type the command:
```sh
air
```

This will begin to "hot-reload", which means that as you work on the project and save
your progress, it will automatically reload the server so you can see your changes
fast!

If you want to stop `air`, just go to your terminal, and use the keybind 'Control-C'
(both on Windows and Mac) to cancel out of `air`.
