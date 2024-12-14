package code.advent;

import java.util.List;
import java.util.stream.Stream;

public class Day04 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day04();
        solution.test();
        solution.execute();
    }

    private int wordSearch(Puzzle puzzle, String word, int i, int j, int index, int di, int dj) {
        if (index == word.length() - 1) {
            return 1;
        }
        int result = 0;
        int x = i + di;
        int y = j + dj;
        if (0 <= x && x < puzzle.rows && 0 <= y && y < puzzle.cols) {
            if (puzzle.input.get(x).charAt(y) == word.charAt(index + 1)) {
                result = result + wordSearch(puzzle, word, x, y, index + 1, di, dj);
            }
        }
        return result;
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        Puzzle puzzle = new Puzzle(inputStream.toList());
        int count = 0;
        for (int i = 0; i < puzzle.rows; i++) {
            for (int j = 0; j < puzzle.cols; j++) {
                if (puzzle.input.get(i).charAt(j) == 'X') {
                    for (int di = -1; di <= 1; di++) {
                        for (int dj = -1; dj <= 1; dj++) {
                            if (di == 0 && dj == 0) {
                                continue;
                            }
                            count += wordSearch(puzzle, "XMAS", i, j, 0, di, dj);
                        }
                    }
                }
            }
        }
        return String.valueOf(count);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        Puzzle puzzle = new Puzzle(inputStream.toList());
        int count = 0;

        for (int i = 0; i < puzzle.rows - 2; i++) {
            for (int j = 0; j < puzzle.cols - 2; j++) {
                boolean left = puzzle.input.get(i + 1).charAt(j + 1) == 'A' && (
                        (puzzle.input.get(i).charAt(j) == 'M' && puzzle.input.get(i + 2).charAt(j + 2) == 'S') ||
                                (puzzle.input.get(i).charAt(j) == 'S' && puzzle.input.get(i + 2).charAt(j + 2) == 'M'));
                boolean right = puzzle.input.get(i + 1).charAt(j + 1) == 'A' && (
                        (puzzle.input.get(i + 2).charAt(j) == 'M' && puzzle.input.get(i).charAt(j + 2) == 'S') ||
                                (puzzle.input.get(i + 2).charAt(j) == 'S' && puzzle.input.get(i).charAt(j + 2) == 'M'));
                count += left && right ? 1 : 0;
            }
        }

        return String.valueOf(count);
    }

    private static class Puzzle {
        public List<String> input;
        public int rows, cols;

        public Puzzle(List<String> input) {
            this.input = input;
            this.rows = input.size();
            this.cols = input.get(0).length();
        }
    }
}
