Feature: Starting Board

   Before a new board can be returned to a player
   it needs to have the gophers removed.

  Scenario: Starting boards contain no gophers
    Given a medium 5x5 board
    When a new game is started
    Then no gophers should be returned to the player