# ethan-go README
### overview
this is ethan written in go.

### explanation of ethan
ethan is basically really terrible dreidel. 
Everybody starts with n chips, Ethan (the "pot" so to speak) has 0, and takes turns rolling 2 dice. If the player rolls a 4, they receive all of Ethan's chips. If Ethan Eyes rule is active and teh player rolls a 2, player loses all chips to Ethan. All other rolls, the player gives a chip to Ethan. Once the player runs out of chips, no more rolls. If all players run out of chips, Ethan wins, but if a player ends up with all of the chips, they are victorious. 
Yes I played this in real life. 

### How to build
there are no external libraries required for this so just go build -o <somedir> ./... after checking it out. 

### How to play
run the binary and press enter to do turns. You may turn on autoplay option if you just want to run through a game as fast as possible. You may set the number of players, number of starting chips per player, and whether or not Ethan Eyes is active. 
