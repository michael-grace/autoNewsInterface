# AutoNews Selection Interface

 This vaguely goes with thonkyBoi to allow URY presenters to disable autonews for their shows.
It uses the same config.json file as thonkyBoi autoSelector.

## Working with This

### (best of luck, this is made with little love)

* Copy `config.json.example` to `config.json`
* By default is on port 3000, can be changed at top of `main.go` (this was definitely not going in a log file, or the API key, the one in config **doesn't do stuff** in this project :) ). Make sure the other consts are suitable
* Build with `go build`
* That's all I have to say

###### Michael Grace 2020