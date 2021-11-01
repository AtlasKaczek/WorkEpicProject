# WorkEpicProject
## Project description
Each week, Epic Games Store makes one non free-to-play game available for free. You can see which game here: https://www.epicgames.com/store/en-US/free-games (https://www.epicgames.com/store/pl/free-games for Polish version).

It is a Go application that retrieves information about which game is discounted this way for the current week from this site and sends a Slack message containing the game's title with a link to game's store page. This app uses a webhook found through some tinkering on epic games store free games page:
->f12
->Go to "Network"
->Ctrl+r to reloade the page
->Ctrl+f to search for some game that's currently free and click on the result
->Go to "Headers"
->"Request url" is what we want

This app uses "Incoming WebHooks" app extension to slack channel. Feel free to change url given by this extension to use this on your slack channel.

## Project files
### json_func.go
Contains 2 functions.
"getJson"
It takes an url (string) from epic games store free games page and makes a http.Get request (it doesn't really change so leave it as it is); returns json body ([]byte) and an error (that should be nil).

"parseJson"
It also takes an url (string) from epic games store free games page and gives it to "getJson" function and unmarshals (thanks to structures in json_struct.go file) the result. Returns a "Games" object and an error (that should be nil).

### json_structure.go
Contains only nessecary structures (which are a representation of epic's json file) and useful functions (getters) that return values from "Games" object. We will them need in other functions.

### operations.go
Contains 2 functions.
"CheckFreeGame"
As it's name says, it checks which games is free because of a promotion for current week and returns a list of titles and their indexes in "Games" object.

"PrepMessage"
It prepares the massage that will be sent to slack in form of json syntax. The message contains a short informational text and free game names with appropriate hiperlinks to their store pages.

### slack_func.go
Contains 1 function.
"SendSlackMessege"
It takes "PrepMessage" result and url (string), marshals the text and sends it to slack.

### main.go
Basic main function that sticks together our app.
