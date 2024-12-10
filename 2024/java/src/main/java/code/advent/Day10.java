package code.advent;

import java.util.List;
import java.util.Stack;
import java.util.stream.Stream;

public class Day10 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day10();
        solution.test();
        solution.execute();
    }

    private static final class Grid {
        public final int[][] grid;
        public final int rows;
        public final int cols;

        private record Loc(int i, int j) {
        }

        private Grid(int[][] grid) {
            this.grid = grid;
            this.rows = grid.length;
            this.cols = grid[0].length;
        }

        public static Grid parse(List<String> input) {
            int[][] grid = new int[input.size()][];
            for (int i = 0; i < input.size(); i++) {
                grid[i] = input.get(i).chars().map(k -> k - '0').toArray();
            }
            return new Grid(grid);
        }

        public boolean isInBounds(int i, int j) {
            return i >= 0 && i < rows && j >= 0 && j < cols;
        }

        public int getPaths(int x, int y, boolean unique) {
            int score = 0;
            Stack<Loc> depthFirst = new Stack<>();
            depthFirst.add(new Loc(x, y));

            boolean[][] visited = new boolean[rows][cols];

            while (!depthFirst.isEmpty()) {
                Loc loc = depthFirst.pop();
                int i = loc.i;
                int j = loc.j;

                if (unique && visited[i][j]) {
                    continue;
                }

                visited[i][j] = true;

                if (grid[i][j] == 9) {
                    score += 1;
                    continue;
                }

                if (isInBounds(i - 1, j) && grid[i - 1][j] == grid[i][j] + 1) {
                    depthFirst.add(new Loc(i - 1, j));
                }
                if (isInBounds(i, j + 1) && grid[i][j + 1] == grid[i][j] + 1) {
                    depthFirst.add(new Loc(i, j + 1));
                }
                if (isInBounds(i + 1, j) && grid[i + 1][j] == grid[i][j] + 1) {
                    depthFirst.add(new Loc(i + 1, j));
                }
                if (isInBounds(i, j - 1) && grid[i][j - 1] == grid[i][j] + 1) {
                    depthFirst.add(new Loc(i, j - 1));
                }
            }
            return score;
        }

        public int getScore(int i, int j) {
            return getPaths(i, j, true);
        }

        public int getRating(int i, int j) {
            return getPaths(i, j, false);
        }
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        Grid trail = Grid.parse(inputStream.toList());
        int score = 0;

        for (int i = 0; i < trail.rows; i++) {
            for (int j = 0; j < trail.cols; j++) {
                if (trail.grid[i][j] == 0) {
                    score += trail.getScore(i, j);
                }
            }
        }

        return String.valueOf(score);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        Grid trail = Grid.parse(inputStream.toList());
        int rating = 0;

        for (int i = 0; i < trail.rows; i++) {
            for (int j = 0; j < trail.cols; j++) {
                if (trail.grid[i][j] == 0) {
                    rating += trail.getRating(i, j);
                }
            }
        }

        return String.valueOf(rating);
    }
}
