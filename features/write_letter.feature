Feature: Writing letters

   In order to play the game, a player must be
   able to write letter to particular places on
   the board

   Holes cannot be overwritten

   Available options:
    * "g" - gopher
    * "e" - earth (player entered empty)
    * " " - space (unknown/default)

   Scenario Outline: Successful writes
    Given a board full of <board_fill>
    When <test_type> is entered to the first position
    Examples:
    | board_fill | test_type |
    | space      | gopher    |
    | gopher     | earth     |
    | earth      | space     |

   Scenario: Unsuccessful write
    Given a board full of holes
    When gopher is entered to the first position
