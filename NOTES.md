# notes

## Initial game loop

* Display game board
* Take in user input
* Check for completion

## Later Steps/Things

* Return message if placement invalid?
* Score based on time and hints
* Timing info
* Reset button
* Undo button
* Game
    * User
    * User stats
    * hints
* Create and store seed for levels when made so that they can be replayed in the future.

## Errors to check for
    * Gophers too close
    * Rows/Columns having too many gophers
    * Rows/Columns being filled but not having enough gophers

## Thoughts/ideas

* Possibly change to check for empty hole and make sure there is at least one other empty hole in the surrounding 8 areas? Currently looking for open spot, then randomly checking in a direction to see if it's in bounds and empty
