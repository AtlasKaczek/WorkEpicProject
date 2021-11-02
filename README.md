# WorkEpicProject

## Project description

Each week, Epic Games Store makes one non free-to-play game available for free. You can see which game here: https://www.epicgames.com/store/en-US/free-games (https://www.epicgames.com/store/pl/free-games for Polish version).

It is a Go application that retrieves information about which game is discounted this way for the current week from this site and sends a Slack message containing the game's title with a link to game's store page. This app uses a request found through some tinkering on epic games store free games page:
> 1. f12
> 1. Go to "Network"
> 1. Ctrl+r to reloade the page
> 1. Ctrl+f to search for some game that's currently free and click on the result
> 1. Go to "Headers"
> 1. "Request url" is what we want

This app uses "Incoming WebHooks" app extension to slack channel. Feel free to change url given by this extension to use this on your slack channel.

## Project files

### json_func.go

Contains 2 functions.
`getJson`
It takes an url (string) from epic games store free games page and makes a http.Get request (it doesn't really change so leave it as it is); returns json body ([]byte) and an error (that should be nil).

`ParseJSON`
It also takes an url (string) from epic games store free games page and gives it to `getJSON` function and unmarshals (thanks to structures in json_struct.go file) the result. Returns a `Games` object and an error (that should be nil).

### json_structure.go

Contains only nessecary structures (which are a representation of epic's json file) and useful functions (getters) that return values from `Games` object. We will them need in other functions.

### operations.go

Contains 3 functions.
`CheckFreeGame`
As it's name says, it checks which games is free because of a promotion for current week and returns a list of titles and their indexes in `Games` object.

`PrepMessage`
It prepares the massage that will be sent to slack in form of json syntax. The message contains a short informational text and free game names with appropriate hiperlinks to their store pages.

`EpicWeekly`
It checks every 60 minutes if there is new free game and uses `SendSlackMessage`.

### slack_func.go

Contains 1 function.
`SendSlackMessege`
It takes `PrepMessage` result and url (string), marshals the text and sends it to slack.

### main.go

Basic main function that sticks together our app.

### Makefile

Contains few parameters, so it's easier to use our app.

- ```make run```
Compiles and runs our app.

- ```make build```
Compiles our app.

- ```make compile```
Compiles our app for other OS and platforms (linux, linux64 and freebsd)

- ```make tidy```
Checks binary for our app

### Dockerfile

Help us create docker image so we can make a container running application.

### .gitignore

Tells github what files to ignore.

## How to run the app

At first pull the code:

    git clone https://github.com/AtlasKaczek/WorkEpicProject.git

### Locally

Run command and set SLACKURL to your webhook

    export SLACKURL="yourSlackURL"
	make run

### Locally via Docker

First of you need to create an image by running a command:

    docker build -t workepicproject:v1
The second step is to run a container of our image:

    docker run --rm -e SLACKURL= --name epiccontainer workepicproject:v1

### Kubernetes

    kubectl apply -f kubernetes_files/namespace.yaml

Change `slack_url` in `kubernetes_files/secret.yaml` to your webhook

    kubectl -n epic-free-game apply -f kubernetes_files/secret.yaml
    kubectl -n epic-free-game apply -f kubernetes_files/deployment.yaml