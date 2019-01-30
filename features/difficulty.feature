Feature: Game difficulty

  To provide challenges to players of differing
  skill levels, the game supports two
  difficulties:
  * Medium/Standard
  * Hard

  Boards are filled according to a difficulty
  level and that dictates how much of the board
  will be occupied by gophers and holes
  * Medium: ~30%
  * Hard: ~50%

  Scenario Outline: Medium difficulty
    When a <dimension> board is filled at medium difficulty
    Then approximately <count> spaces will be filled
      Examples:
      | dimension | count |
      | 4x4       | 5     |
      | 5x5       | 8     |
      | 50x50     | 750   |
      | 100x75    | 2250  |

  Scenario Outline: Hard difficulty
    When a <dimension> board is filled at hard difficulty
    Then approximately <count> spaces will be filled
      Examples:
      | dimension | count |
      | 4x4       | 8     |
      | 5x5       | 13    |
      | 50x50     | 1250  |
      | 100x75    | 3750  |
