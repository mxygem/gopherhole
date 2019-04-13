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
      When a <char_type> is entered into position (0, 0)
      Then that position must contain the expected character
      Examples:
      | board_fill | char_type |
      | spaces     | gopher    |
      | gophers    | earth     |
      | earth      | space     |

   Scenario Outline: Unsuccessful writes
      Given a 3x3 sized board full of <board_fill>
      When a <char_type> is entered into position (0, 0)
      Then a placement error of <error> must be returned
      And that position must contain a <orig_char> character
      Examples:
      | board_fill | char_type | orig_char | error                         |
      | holes      | gopher    | hole      | "holes cannot be overwritten" |
      | spaces     | hole      | space     | "holes cannot be placed"      |
