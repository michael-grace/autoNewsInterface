# AutoNews Selection Interface

 This vaguely goes with thonkyBoi to allow URY presenters to disable autonews for their shows.
It uses the same config.json file as thonkyBoi autoSelector.

## Working with This

### (best of luck, this is made with little love)

* Copy `config.json.example` to `config.json`. This is the switcher config
* Copy `interface_config.json.example` to `interface_config.json`. This is the interface's config. Changes to this should be reflected in the struct towards the top of `main.go`
* By default is on port 3000, can be changed in the `interface_config.json`. The file path of this config is towards the top of `fun main(){` and contains various useful things.
* Build with `go build`
* That's all I have to say

###### Michael Grace 2020