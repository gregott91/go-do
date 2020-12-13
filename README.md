# Go-Do

## Installation
This package is distributed with a custom scoop bucket.

To install, run the following commands in Powershell:

    iwr -useb get.scoop.sh | iex   # installs scoop
    scoop bucket add gregott91 https://github.com/gregott91/go-do.git   # adds a custom scoop bucket
    scoop install go-do

## Running
To run, simply run `go-do` from powershell.

Optionally, if you have AutoHotKey installed, double-click either of the .ahk files. Following that, Windows key + Z will run go-do.
* run-godo-cmd will run go-do in a command line window.
* run-godo-wt will run go-do in a windows terminal window.

## Usage

### Basic
To switch between text input and the list, type Ctrl+S
After typing a note, hit Enter to add it to the list
When focused on the list, type Ctrl+D to delete an entry
No matter the focus, hit Ctrl+C or the Escape key to exit

### Configuration
To change keyboard shortcuts, open the go-do.config.json file