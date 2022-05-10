### Documentation:

## Setup:
Instructions on how to build/run your application
(Preferred OS- MacOS) 

-  The repo exists within a standard golang route. (go/src/github.com/)
Therefore clone this repo within a go/src/github folder.

- Under the file go/src/github.com/tictacgo/ execute the following MacOS/Linux command-
    source env/development.env
    go run cmd/tictacgo/main.go 
    or 
    under cmd/tictacgo/ Run
    go run .


  Test - 
  To run unit test for each operation run

  go test test/playgroundService/playground_test.go

    Versions used -
      go version go1.15.3 darwin/amd64
    The backend will be up on localhost:8080.
    
 
 ## Game Play:
    The user plays the game of TicTacToe with an AI.
    The user can start the game by pressing s or quit using q.
    Press s to start game, q to quit
    
    The user then enter the locations of the character X in the format r c where 0<=r<=2 and 0<=c<=2 based on where they want to put in and the board gets     displayed after that depicting the position of user input. The game AI plays their move as a result of that. The board is then displayed to the user       and they can spot the location of Os.
    eg-
    1 1
    [     ]
    [  X  ]
    [     ]

    [     ]
    [O X  ]
    [     ]
    
    After this the user is again asked for their input and the game continues.
    
    2 2
    [     ]
    [O X  ]
    [    X]

    [    O]
    [O X  ]
    [    X]
    
    0 0
    [X   O]
    [O X  ]
    [    X]

    [X   O]
    [O X  ]
    [  O X]
    
    
    The game either ends with X or O being the winner or ends in a draw.
    
    Game is over
    X is the winner

    
    Once the game is finished the user is again prompted to start a new game or quit.
    Press s to start game, q to quit
    
    
## TODO
    Implement a leaderboard to determine the top players.
    Implement a system which allows the user to save their game and load it later.
