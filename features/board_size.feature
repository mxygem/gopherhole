Feature: Board Size

  Users can specify the size of board that they
  would like to play, though if nothing is set, a
  4x4 board will be returned.

  Scenario: No board size is set returns a 4x4 board
    When a new game is requested with no board size set
    Then a 4x4 board must be returned

  Scenario Outline: Requesting a 5x5 board returns a 5x5 board
    When a new game is requested with a <dimension> board size
    Then a <dimension> board must be returned
      Examples:
      | dimension |
      | 5x5       |
      | 50x50     |
      | 100x75    |
