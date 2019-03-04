Feature: Writing characters

   In order to play the game, a player must be
   able to write characters to particular places on
   the board

   Holes cannot be written or overwritten

   Available options:
    * "g" - gopher
    * "e" - earth (player entered empty)
    * " " - space (unknown/default)

   Scenario Outline: Successful writes
    Given a 3x3 sized board full of <board_fill>
    When a <test_type> is entered into position (0, 0)
    Then that position must contain the expected character
    Examples:
    | board_fill | test_type |
    | space      | gopher    |
    | gopher     | earth     |
    | earth      | space     |

   Scenario: Unsuccessful write
    Given a board full of holes
    When a gopher is entered to the first position
