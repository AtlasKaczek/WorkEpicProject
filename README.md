# WorkEpicProject
## Project description
Each week, Epic Games Store makes one non free-to-play game available for free. You can see which game here: https://www.epicgames.com/store/en-US/free-games (https://www.epicgames.com/store/pl/free-games for Polish version).

It is a Go application that retrieves information about which game is discounted this way for the current week from this site and sends a Slack message containing the game's title with a link to game's store page. This app uses a webhook found through some tinkering on epic games store free games page (f12->Network->Ctrl+f->Search some game that's currently free->And there is the needed url to get json file).