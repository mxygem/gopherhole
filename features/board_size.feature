Feature: Board Size

  Users can specify the size of board that they would like to play, though if nothing is set, a 5x5 board will be returned.

  Scenario: No board size is set returns a 4x4 board
    When a new game is requested with no board size set
    Then a 4x4 board must be returned

  Scenario: Requesting a 5x5 board returns a 5x5 board
    When a new game is requested with a 5x5 board size
    Then a 5x5 board must be returned
