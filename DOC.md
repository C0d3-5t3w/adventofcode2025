# Advent of Code 2025 - Code Documentation

This document provides detailed explanations of all `main.c` files, their functions, and data structures.

---

## Day 1

### Part 1: `day1/part1/c/main.c`

**Purpose:** Simulates a dial that moves left or right based on commands, counting how many times it lands exactly on zero.

**Functions:**
- `main()` - Entry point that reads direction/distance pairs from `list.txt` and simulates dial movement.

**Logic:**
- Dial starts at position 50 (range 0-99, wrapping around at 100)
- Reads commands in format `<direction><distance>` (e.g., `L25`, `R30`)
- `L` moves the dial left (subtraction with modulo wrap)
- `R` moves the dial right (addition with modulo wrap)
- Counts each time the dial lands exactly on 0 after a move
- Outputs the total zero count

**Variables:**
- `dial` (int): Current position on the dial (0-99)
- `zero_count` (int): Number of times dial landed on 0
- `direction` (char): 'L' or 'R' for left/right movement
- `distance` (int): How far to move the dial

---

### Part 2: `day1/part2/c/main.c`

**Purpose:** Extended version that counts how many times the dial **crosses** zero during movement, not just lands on it.

**Functions:**
- `main()` - Entry point that tracks zero crossings during dial movement.

**Logic:**
- Same dial mechanics as Part 1
- For each move, calculates how many times the dial crosses 0 during the movement
- Uses formula: `crosses = 1 + (distance - first_zero) / 100` when distance >= first_zero
- `first_zero` is the distance to the next zero position from current position

**Variables:**
- Same as Part 1, plus:
- `first_zero` (int): Distance to the next zero crossing
- `crosses` (int): Number of zero crossings in a single move

---

## Day 2

### Part 1: `day2/part1/c/main.c`

**Purpose:** Finds "invalid IDs" within given ranges and sums them. An invalid ID is a number whose decimal representation is exactly two identical halves (e.g., `1212`, `123123`).

**Functions:**
- `is_invalid_id(unsigned long long num)` - Checks if a number is an invalid ID (first half equals second half)
- `next_invalid_id(unsigned long long start)` - Finds the next invalid ID >= start
- `sum_invalid_ids_in_range(unsigned long long start, unsigned long long end)` - Sums all invalid IDs in a range
- `main()` - Parses ranges from `list.txt` and computes total sum

**Logic:**
- Invalid IDs have even-length digit strings where first half == second half
- Uses efficient generation: constructs candidates by repeating half-patterns
- Parses ranges in format `start-end` from input file

**Key Algorithm:**
- For a target length, find the smallest half-pattern that produces a number >= start
- Multiply pattern by `10^half + 1` to create the repeated number

---

### Part 2: `day2/part2/c/main.c`

**Purpose:** Extended invalid ID detection where a number is invalid if it consists of **any** repeating pattern (not just 2 repetitions).

**Functions:**
- `is_invalid_id(unsigned long long num)` - Checks if number has any repeating pattern
- `build_repeated(unsigned long long pattern, int pattern_len, int repeats)` - Constructs a number from repeated pattern
- `add_candidate(unsigned long long val)` - Adds a candidate to the search array
- `next_invalid_id(unsigned long long start)` - Finds next invalid ID using binary search
- `sum_invalid_ids_in_range(...)` - Sums all invalid IDs in range
- `main()` - Entry point

**Data Structures:**
- `candidates[MAX_CANDIDATES]` (global array): Stores potential invalid ID candidates
- `num_candidates` (global int): Count of candidates

**Logic:**
- Checks all possible pattern lengths that evenly divide the number length
- Pattern must repeat at least 2 times
- Uses binary search to efficiently find the smallest valid candidate for each pattern length

---

## Day 3

### Part 1: `day3/part1/c/main.c`

**Purpose:** For each line, finds the maximum 2-digit "joltage" value by selecting any two digits from the line and forming a number.

**Functions:**
- `max_joltage(const char *line, int len)` - Finds maximum 2-digit value from any two positions
- `main()` - Reads lines from `list.txt` and sums max joltages

**Logic:**
- For each pair of positions (i, j) where i < j, form a 2-digit number: `line[i]*10 + line[j]`
- Track the maximum across all pairs
- Sum all line maximums

**Constants:**
- `MAX_LINE_LENGTH` (1024): Maximum characters per line

---

### Part 2: `day3/part2/c/main.c`

**Purpose:** Extended version finding maximum 12-digit joltage by greedily selecting the largest available digits.

**Functions:**
- `max_joltage(const char *line, int len)` - Greedy algorithm to build maximum 12-digit number
- `main()` - Reads lines and sums joltages

**Logic:**
- Greedy selection: for each of 12 digit positions, choose the largest digit that still leaves enough characters for remaining positions
- Uses constraint: `max_pos = len - (NUM_DIGITS - i)` to ensure enough digits remain
- Returns 0 for lines shorter than 12 characters

**Constants:**
- `NUM_DIGITS` (12): Target number of digits to extract

---

## Day 4

### Part 1: `day4/part1/c/main.c`

**Purpose:** Counts "accessible" paper rolls (`@`) in a grid. A roll is accessible if it has fewer than 4 adjacent rolls.

**Functions:**
- `main()` - Reads grid from `list.txt` and counts accessible rolls

**Logic:**
- Parses a 2D character grid
- For each `@` character, counts adjacent `@` characters in 8 directions
- If count < 4, the roll is "accessible"

**Data Structures:**
- `grid[MAX_SIZE][MAX_SIZE]` (char array): 2D grid storage
- `dx[]`, `dy[]` (int arrays): Direction vectors for 8 neighbors

**Constants:**
- `MAX_SIZE` (256): Maximum grid dimension

---

### Part 2: `day4/part2/c/main.c`

**Purpose:** Simulates iterative removal of accessible paper rolls until no more can be removed.

**Functions:**
- `count_adjacent_rolls(char grid[][MAX_SIZE], int rows, int cols, int r, int c)` - Counts neighboring `@` cells
- `main()` - Simulates removal rounds

**Logic:**
- Repeatedly finds all accessible rolls (< 4 neighbors) in each round
- Marks them for removal simultaneously (to avoid order-dependent results)
- Replaces removed rolls with `.`
- Continues until no rolls are removed in a round
- Counts total removed rolls

**Data Structures:**
- `to_remove[MAX_SIZE][MAX_SIZE]` (bool array): Marks cells for removal each round

---

## Day 5

### Part 1: `day5/part1/c/main.c`

**Purpose:** Counts how many ingredient IDs fall within any of the "fresh" ranges.

**Functions:**
- `main()` - Parses ranges and IDs, counts fresh ingredients

**Data Structures:**
- `Range` (struct): Contains `start` and `end` (long long) for an ID range
- `ranges[MAX_RANGES]` (Range array): Stores all fresh ranges

**Logic:**
- First section of input: ranges in format `start-end`
- Second section (after blank line): individual IDs to check
- For each ID, check if it falls within any range
- Count matching IDs

**Constants:**
- `MAX_RANGES` (1000): Maximum number of ranges
- `MAX_LINE` (100): Maximum line length

---

### Part 2: `day5/part2/c/main.c`

**Purpose:** Counts the **total number** of fresh ingredient IDs by merging overlapping ranges.

**Functions:**
- `compare_ranges(const void *a, const void *b)` - Comparison function for qsort
- `main()` - Parses, sorts, merges ranges, and counts total IDs

**Data Structures:**
- `Range` (struct): Same as Part 1
- Dynamically allocated `ranges` array

**Logic:**
- Parses all ranges from input
- Sorts ranges by start value using qsort
- Merges overlapping/adjacent ranges (if `ranges[i].start <= current_end + 1`)
- Sums the size of each merged range: `(end - start + 1)`

**Constants:**
- `MAX_RANGES` (10000): Increased capacity for more ranges

---

## Day 6

### Part 1: `day6/part1/c/main.c`

**Purpose:** Processes vertical columns of numbers with operators, computing a grand total.

**Functions:**
- `main()` - Parses columnar input and evaluates expressions

**Logic:**
- Input format: multiple rows of numbers arranged in columns, with an operator row at bottom
- Identifies "groups" of columns (separated by all-space columns)
- For each group:
  - Extracts numbers from each row horizontally within the group
  - Reads the operator (`+` or `*`) from the bottom row
  - Applies the operator to all numbers in the group
- Sums all group results

**Data Structures:**
- `lines[MAX_ROWS][MAX_LINE]` (char array): Stores all input lines
- `numbers[MAX_ROWS]` (long long array): Numbers extracted from a column group

**Constants:**
- `MAX_LINE` (16384): Maximum line length
- `MAX_ROWS` (10): Maximum number of rows

---

### Part 2: `day6/part2/c/main.c`

**Purpose:** Similar to Part 1, but reads numbers **vertically** within each column (digits stacked on top of each other).

**Functions:**
- `main()` - Parses vertical numbers and evaluates expressions

**Logic:**
- Same column-group detection as Part 1
- Within each group, iterates column by column (right to left)
- For each column position, reads digits vertically from all rows to form a number
- Applies the operator to all extracted numbers
- Numbers are processed right-to-left within each group

**Data Structures:**
- `numbers[256]` (long long array): Increased capacity for vertical numbers

**Key Difference from Part 1:**
- Part 1 reads numbers horizontally across a row
- Part 2 reads numbers vertically down a column

---

## Day 7

### Part 1: `day7/part1/c/main.c`

**Purpose:** Simulates a tachyon beam splitting through a manifold grid, counting the total number of beam splits.

**Functions:**
- `main()` - Entry point that reads the manifold grid from `list.txt` and simulates beam propagation.

**Logic:**
- A tachyon beam enters at position 'S' on the first row and moves downward
- Beams pass freely through empty space (`.`)
- When a beam encounters a splitter (`^`), it splits into two new beams: one going to the immediate left and one to the immediate right
- **Key insight:** When multiple beams converge on the same column position, they merge into a single beam (not tracked as multiple separate beams)
- Each splitter activation counts as one split
- Simulation continues until all beams exit the grid or hit splitters

**Data Structures:**
- `grid[MAX_ROWS][MAX_COLS]` (char array): 2D grid storing the manifold diagram
- `beams[cols]` (bool array): Tracks which columns have active beams (true/false, not count)
- `next_beams[cols]` (bool array): Buffer for next row's beam positions

**Algorithm:**
1. Find starting position 'S' in the first row
2. Initialize a single beam at the starting column
3. For each row (top to bottom):
   - For each column with an active beam:
     - If cell is `^`: increment split count, mark left and right positions for next row
     - Otherwise: beam continues straight down
   - Swap beam arrays for next iteration
4. Output total split count

**Constants:**
- `MAX_ROWS` (256): Maximum grid rows
- `MAX_COLS` (256): Maximum grid columns

**Variables:**
- `start_col` (int): Column position of 'S'
- `total_splits` (int): Running count of splitter activations
- `rows`, `cols` (int): Grid dimensions

---

### Part 2: `day7/part2/c/main.c`

**Purpose:** Simulates a quantum tachyon particle traversing a manifold, counting the total number of distinct timelines created by the many-worlds interpretation of beam splitting.

**Functions:**
- `main()` - Entry point that reads the manifold grid from `list.txt` and simulates quantum particle propagation.

**Logic:**
- A single tachyon particle enters at position 'S' and moves downward
- The particle passes freely through empty space (`.`)
- When the particle encounters a splitter (`^`), time itself splits: in one timeline the particle goes left, in another it goes right
- **Key insight:** Unlike Part 1, timelines do NOT merge when they reach the same position - each timeline is tracked separately
- The answer is the total number of timelines after the particle completes all possible journeys

**Data Structures:**
- `grid[MAX_ROWS][MAX_COLS]` (char array): 2D grid storing the manifold diagram
- `timelines[cols]` (long long array): Tracks number of timelines at each column position
- `next_timelines[cols]` (long long array): Buffer for next row's timeline counts

**Algorithm:**
1. Find starting position 'S' in the first row
2. Initialize with one timeline at the starting column
3. For each row (top to bottom):
   - For each column with active timelines:
     - If cell is `^`: send timeline count to both left and right positions (each timeline splits)
     - Otherwise: timelines continue straight down
   - Swap timeline arrays for next iteration
4. Sum all timeline counts across all columns at the end
5. Output total timeline count

**Key Difference from Part 1:**
- Part 1 tracks boolean presence of beams (beams merge at same position)
- Part 2 tracks timeline counts (timelines remain separate, each split doubles the count for that path)

**Constants:**
- `MAX_ROWS` (256): Maximum grid rows
- `MAX_COLS` (256): Maximum grid columns

**Variables:**
- `start_col` (int): Column position of 'S'
- `total_timelines` (long long): Final count of all timelines
- `rows`, `cols` (int): Grid dimensions

---

## Day 8

### Part 1: `day8/part1/c/main.c`

**Purpose:** Connects junction boxes in 3D space by their shortest distances, then calculates the product of the three largest circuit sizes after 1000 connection attempts.

**Functions:**
- `init_union_find(int n)` - Initializes Union-Find data structure for n elements
- `find(int x)` - Finds root of element x with path compression
- `union_sets(int x, int y)` - Merges sets containing x and y, returns false if already same set
- `distance(Point a, Point b)` - Calculates Euclidean distance between two 3D points
- `compare_edges(const void *a, const void *b)` - Comparison function for sorting edges by distance
- `compare_int_desc(const void *a, const void *b)` - Comparison function for descending integer sort
- `main()` - Entry point that reads junction boxes, connects them, and computes result

**Data Structures:**
- `Point` (struct): Contains `x`, `y`, `z` (int) for 3D coordinates
- `Edge` (struct): Contains `box1`, `box2` (int) indices and `distance` (double)
- `parent[MAX_BOXES]` (int array): Union-Find parent pointers
- `rank_arr[MAX_BOXES]` (int array): Union-Find rank for balancing

**Logic:**
- Reads 1000 junction box coordinates from input (X,Y,Z format)
- Calculates all pairwise distances (499,500 edges)
- Sorts edges by distance (shortest first)
- Processes the 1000 shortest pairs using Union-Find:
  - If boxes are in different circuits, connect them
  - If already in same circuit, skip (but still counts toward 1000)
- Counts size of each resulting circuit
- Multiplies the three largest circuit sizes

**Algorithm:** Kruskal's-like approach with Union-Find (Disjoint Set Union)

**Constants:**
- `MAX_BOXES` (1000): Maximum number of junction boxes
- `CONNECTIONS` (1000): Number of shortest pairs to process

---

### Part 2: `day8/part2/c/main.c`

**Purpose:** Continues connecting junction boxes until all form a single circuit, then returns the product of X coordinates of the last two connected boxes.

**Functions:**
- `init_union_find(int n)` - Initializes Union-Find data structure for n elements
- `find(int x)` - Finds root of element x with path compression
- `union_sets(int x, int y)` - Merges sets containing x and y, returns false if already same set
- `distance(Point a, Point b)` - Calculates Euclidean distance between two 3D points
- `compare_edges(const void *a, const void *b)` - Comparison function for sorting edges by distance
- `main()` - Entry point that connects all boxes and finds the final connection

**Data Structures:**
- `Point` (struct): Contains `x`, `y`, `z` (int) for 3D coordinates
- `Edge` (struct): Contains `box1`, `box2` (int) indices and `distance` (double)
- `parent[MAX_BOXES]` (int array): Union-Find parent pointers
- `rank_arr[MAX_BOXES]` (int array): Union-Find rank for balancing

**Logic:**
- Same setup as Part 1: read coordinates, calculate all distances, sort edges
- Connects boxes in order of shortest distance until all are in one circuit
- Requires exactly n-1 connections to connect n boxes
- Tracks the last successful connection made
- Multiplies X coordinates of the two boxes in the final connection

**Algorithm:** Minimum Spanning Tree construction using Kruskal's algorithm with Union-Find

**Key Difference from Part 1:**
- Part 1 processes exactly 1000 pairs (some may be redundant)
- Part 2 continues until all boxes connected (exactly 999 successful connections for 1000 boxes)

**Constants:**
- `MAX_BOXES` (1000): Maximum number of junction boxes

**Variables:**
- `last_box1`, `last_box2` (int): Indices of boxes in the final connection
- `connections_needed` (int): n-1 connections required for n boxes

---

## Day 9

### Part 1: `day9/part1/c/main.c`

**Purpose:** Finds the largest rectangle that can be formed using any two red tiles as opposite corners in a movie theater tile grid.

**Functions:**
- `main()` - Entry point that reads red tile coordinates from `list.txt` and finds the maximum rectangle area.

**Data Structures:**
- `Point` (struct): Contains `x`, `y` (int) for 2D coordinates of red tiles

**Logic:**
- Reads all red tile coordinates in format `x,y` from input file
- Tries all pairs of points as potential opposite corners of a rectangle
- Calculates inclusive area: `(|x2 - x1| + 1) * (|y2 - y1| + 1)`
- The `+1` is needed because both corner tiles are included in the rectangle dimensions
- Tracks and outputs the maximum area found

**Algorithm:**
1. Parse all red tile coordinates into an array
2. For each pair of points (i, j) where i < j:
   - Calculate width = `|x_j - x_i| + 1` (inclusive)
   - Calculate height = `|y_j - y_i| + 1` (inclusive)
   - Calculate area = width × height
   - Update maximum if this area is larger
3. Output the largest rectangle area

**Key Insight:**
- The rectangle area is inclusive of both corner tiles, so we add 1 to both dimensions
- Example: corners at (2,5) and (11,1) give width = 10, height = 5, area = 50

**Constants:**
- `MAX_POINTS` (1000): Maximum number of red tile coordinates

**Variables:**
- `points[MAX_POINTS]` (Point array): Stores all red tile coordinates
- `count` (int): Number of red tiles read
- `max_area` (long long): Largest rectangle area found

---

### Part 2: `day9/part2/c/main.c`

**Purpose:** Finds the largest rectangle using red tiles as opposite corners where ALL tiles within the rectangle are red or green (inside or on the boundary of the polygon formed by red tiles).

**Functions:**
- `cmp_hseg(const void *a, const void *b)` - Comparison function for sorting horizontal segments by y-coordinate, then x_min
- `cmp_vseg(const void *a, const void *b)` - Comparison function for sorting vertical segments by x-coordinate, then y_min
- `is_inside_or_on_boundary(int x, int y)` - Determines if a point is inside or on the boundary of the polygon using ray casting
- `is_rectangle_valid(int x1, int y1, int x2, int y2)` - Checks if all tiles along the rectangle's edges are valid (red or green)
- `main()` - Entry point that builds the polygon and finds the maximum valid rectangle

**Data Structures:**
- `Point` (struct): Contains `x`, `y` (int) for 2D coordinates of red tiles
- `HSegment` (struct): Horizontal segment with `y`, `x_min`, `x_max` for boundary edges
- `VSegment` (struct): Vertical segment with `x`, `y_min`, `y_max` for boundary edges
- `h_segments[MAX_POINTS]` (global array): Stores horizontal boundary segments
- `v_segments[MAX_POINTS]` (global array): Stores vertical boundary segments

**Logic:**
- Red tiles form a closed polygon connected by green tiles (axis-aligned edges)
- Green tiles are: (1) on polygon edges between consecutive red tiles, (2) inside the polygon
- A valid rectangle must have all its tiles be red or green
- Uses ray casting algorithm to determine if points are inside the polygon
- Validates rectangle by checking all points along all four edges

**Algorithm:**
1. Parse red tile coordinates from input
2. Build horizontal and vertical segments from consecutive point pairs
3. Sort segments for efficient lookup
4. For each pair of red tiles (potential rectangle corners):
   - Only consider if potential area exceeds current maximum
   - Validate by checking all points on rectangle edges using `is_inside_or_on_boundary`
   - Update maximum if valid and larger
5. Output the largest valid rectangle area

**Ray Casting Algorithm (`is_inside_or_on_boundary`):**
- Cast a ray from point (x, y) to the right
- Count vertical segment crossings where ray intersects segment
- If point is on any segment, return true immediately
- Odd number of crossings = inside polygon

**Constants:**
- `MAX_POINTS` (1000): Maximum number of red tile coordinates

**Global Variables:**
- `h_count` (int): Number of horizontal segments
- `v_count` (int): Number of vertical segments

**Key Difference from Part 1:**
- Part 1 allows any rectangle between two red tiles
- Part 2 restricts rectangles to only include red/green tiles (inside or on polygon boundary)

---

## Day 10

### Part 1: `day10/part1/c/main.c`

**Purpose:** Determines the minimum number of button presses needed to configure indicator lights on factory machines to match their target patterns.

**Functions:**
- `parse_lights(const char *start, int *num_lights)` - Parses the indicator light diagram `[.##.]` and returns the target state as a bitmask
- `parse_button(const char *start, const char **end)` - Parses a button schematic `(0,3,4)` and returns the toggle mask
- `min_presses(int target, int *buttons, int num_buttons, int num_lights)` - Uses BFS to find minimum button presses to reach target state
- `main()` - Entry point that reads machine configurations from `list.txt` and sums minimum presses

**Data Structures:**
- `buttons[MAX_BUTTONS]` (int array): Bitmasks representing which lights each button toggles
- `dist[max_state]` (int array): BFS distance array for state space search
- `queue[max_state]` (int array): BFS queue for state exploration

**Logic:**
- Each machine has indicator lights (initially all off) and buttons that toggle specific lights
- Lights are represented as bits in an integer (bit i = light i)
- Pressing a button XORs the current state with the button's toggle mask
- Since pressing a button twice cancels out, each button is pressed 0 or 1 times
- Uses BFS over the state space (2^num_lights possible states) to find minimum presses

**Algorithm:**
1. For each line in input:
   - Parse target state from `[.##.]` notation (`.` = off/0, `#` = on/1)
   - Parse all button schematics `(0,3,4)` into toggle masks
   - Run BFS from state 0 (all off) to target state
   - Each BFS transition tries pressing each button once (XOR with mask)
   - BFS guarantees finding the minimum number of button presses
2. Sum all minimum press counts across all machines

**Input Format:**
- Each line describes one machine
- `[.##.]` - Target light pattern (4 lights: off, on, on, off)
- `(1,3)` - Button that toggles lights 1 and 3 (0-indexed)
- `{3,5,4,7}` - Joltage requirements (ignored)

**Example:**
- `[.##.] (0,2) (0,1)` - Target: lights 1,2 on (bitmask 0110 = 6)
- Pressing `(0,2)` toggles lights 0,2 → state 0101 = 5
- Pressing `(0,1)` toggles lights 0,1 → state 0110 = 6 ✓
- Minimum: 2 presses

**Constants:**
- `MAX_LINE` (4096): Maximum line length
- `MAX_LIGHTS` (16): Maximum number of indicator lights
- `MAX_BUTTONS` (20): Maximum number of buttons per machine

**Variables:**
- `target` (int): Bitmask of target light configuration
- `num_lights` (int): Number of indicator lights on current machine
- `num_buttons` (int): Number of buttons on current machine
- `total` (long long): Sum of minimum presses for all machines

---

### Part 2: `day10/part2/c/main.c`

**Purpose:** Determines the minimum number of button presses needed to configure joltage level counters on factory machines to match their target values. Unlike Part 1 (which uses XOR toggling), each button press increments affected counters by 1.

**Functions:**
- `parse_button(const char *start, const char **end, int *indices)` - Parses a button schematic `(1,3)` and returns which counters it affects
- `parse_targets(const char *start, int *targets)` - Parses joltage requirements `{3,5,4,7}` into an array of target values
- `check_solution(int presses[])` - Verifies if a given button press combination achieves all target values
- `solve_recursive(int presses[], int button_idx, int current_sum, long long best_so_far)` - Recursive solver with pruning for small cases
- `solve_gauss_with_optimization(void)` - Main solver using Gaussian elimination with null space exploration
- `solve_min_presses(...)` - Entry point for solving each machine configuration
- `main()` - Reads machine configurations from `list.txt` and sums minimum presses

**Data Structures:**
- `g_coeff[MAX_COUNTERS][MAX_BUTTONS]` (global int array): Coefficient matrix where `g_coeff[c][b] = 1` if button b affects counter c
- `g_targets[MAX_COUNTERS]` (global int array): Target values for each counter
- `g_buttons[MAX_BUTTONS][MAX_COUNTERS]` (global int array): Which counters each button affects
- `g_button_counts[MAX_BUTTONS]` (global int array): Number of counters affected by each button

**Logic:**
- Each machine has numeric counters (initially 0) and buttons that increment specific counters by 1
- Pressing button `(1,3)` increments counters 1 and 3 each by 1
- Unlike Part 1, buttons can be pressed multiple times (not just 0 or 1)
- This is an Integer Linear Programming (ILP) problem: find non-negative integers x_i minimizing Σx_i such that Ax = targets
- Uses Gaussian elimination to find the solution space, then searches for minimum

**Algorithm:**
1. For each line in input:
   - Parse target values from `{3,5,4,7}` notation
   - Parse all button schematics into counter-affect lists
   - Build coefficient matrix A where A[counter][button] = 1 if button affects counter
2. Solve using Gaussian elimination with optimization:
   - Perform Gaussian elimination to get reduced row echelon form
   - Identify pivot columns (basic variables) and free columns (free variables)
   - If no free variables: unique solution, verify non-negative integers
   - If free variables exist: enumerate all combinations of free variable values
   - For each combination: compute basic variables, check validity (non-negative, integer), track minimum
3. Fallback to recursive search with pruning for small cases
4. Sum all minimum press counts across all machines

**Gaussian Elimination Details:**
- Builds augmented matrix [A | targets]
- Uses partial pivoting for numerical stability
- Identifies free variables (columns without pivots)
- Expresses pivot variables in terms of free variables
- Searches over free variable space (0 to max_target) to find minimum total presses

**Input Format:**
- Each line describes one machine
- `[.##.]` - Indicator light pattern (ignored in Part 2)
- `(1,3)` - Button that increments counters 1 and 3 (0-indexed)
- `{3,5,4,7}` - Target joltage levels (counter 0 → 3, counter 1 → 5, etc.)

**Example:**
- `{3,5,4,7}` with buttons `(3)`, `(1,3)`, `(2)`, `(2,3)`, `(0,2)`, `(0,1)`
- Need: counter0=3, counter1=5, counter2=4, counter3=7
- One solution: press `(3)` once, `(1,3)` 3×, `(2,3)` 3×, `(0,2)` once, `(0,1)` 2× = 10 presses

**Constants:**
- `MAX_LINE` (4096): Maximum line length
- `MAX_COUNTERS` (16): Maximum number of joltage counters per machine
- `MAX_BUTTONS` (20): Maximum number of buttons per machine

**Variables:**
- `g_num_counters` (global int): Number of counters on current machine
- `g_num_buttons` (global int): Number of buttons on current machine
- `total` (long long): Sum of minimum presses for all machines

**Key Difference from Part 1:**
- Part 1 uses XOR toggling: pressing a button flips light states, pressing twice cancels out (0 or 1 presses per button)
- Part 2 uses additive increments: each press adds 1 to counters, buttons can be pressed any number of times
- Part 1 uses BFS over finite state space (2^n states)
- Part 2 solves a system of linear equations with non-negative integer constraints (ILP)

---
