package code.advent;

import java.util.List;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Day14 extends ISolution {

    private static final Pattern NUMBERS = Pattern.compile("-?\\d+");

    public static void main(String[] args) {
        ISolution solution = new Day14();
        solution.test();
        solution.execute();
    }

    private int wrap(int i, int max) {
        if (i < 0) {
            i += max;
        } else if (i >= max) {
            i -= max;
        }
        return i;
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        List<List<MatchResult>> matchLineList = inputStream.map(NUMBERS::matcher)
                .map(matcher -> matcher.results().toList()).toList();

        Robot[] robotList = new Robot[matchLineList.size()];

        int rows = 0;
        int cols = 0;

        for (int i = 0; i < matchLineList.size(); i++) {
            int[] data = new int[4];
            for (int j = 0; j < matchLineList.get(i).size(); j++) {
                data[j] = Integer.parseInt(matchLineList.get(i).get(j).group());
            }

            Position p = new Position(data[1], data[0]);
            Velocity v = new Velocity(data[3], data[2]);
            rows = Math.max(rows, p.i + 1);
            cols = Math.max(cols, p.j + 1);
            robotList[i] = new Robot(p, v);
        }

        int[][] quarterCounts = new int[2][2];

        int time = 100;
        for (Robot r : robotList) {
            int i = (r.position.i + time * r.velocity.i) % rows;
            if (i < 0) i += rows;
            int j = (r.position.j + time * r.velocity.j) % cols;
            if (j < 0) j += cols;
            if (i * 2 == rows - 1 || j * 2 == cols - 1) {
                continue;
            }
            int qi = i * 2 / rows;
            int qj = j * 2 / cols;

            quarterCounts[qi][qj] += 1;
        }

        long factor = 1;
        for (int[] row : quarterCounts) {
            for (int count : row) {
                factor *= count;
            }
        }
        return String.valueOf(factor);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        List<List<MatchResult>> matchLineList = inputStream.map(NUMBERS::matcher)
                .map(matcher -> matcher.results().toList()).toList();

        Robot[] robotList = new Robot[matchLineList.size()];

        int rows = 0;
        int cols = 0;

        for (int i = 0; i < matchLineList.size(); i++) {
            int[] data = new int[4];
            for (int j = 0; j < matchLineList.get(i).size(); j++) {
                data[j] = Integer.parseInt(matchLineList.get(i).get(j).group());
            }

            Position p = new Position(data[1], data[0]);
            Velocity v = new Velocity(data[3], data[2]);
            rows = Math.max(rows, p.i + 1);
            cols = Math.max(cols, p.j + 1);
            robotList[i] = new Robot(p, v);
        }

        long time = 6620;
        int[][] grid = new int[rows][cols];
        for (Robot r : robotList) {
            grid[r.position.i][r.position.j] += 1;
        }

        char[] end = new char[cols];
        for (int i = 0; i < cols; i++) {
            end[i] = '-';
        }

        for (int t = 1; t <= time; t++) {
            for (Robot r : robotList) {
                grid[r.position.i][r.position.j] -= 1;

                r.position.i = wrap(r.position.i + r.velocity.i, rows);
                r.position.j = wrap(r.position.j + r.velocity.j, cols);

                grid[r.position.i][r.position.j] += 1;
            }
        }

        for (int i = 0; i < rows; i++) {
            char[] row = new char[cols];
            for (int j = 0; j < cols; j++) {
                row[j] = grid[i][j] > 0 ? '#' : ' ';
            }
            System.out.println(row);
        }

        return null;
    }

    private static final class Position {
        public int i;
        public int j;

        private Position(int i, int j) {
            this.i = i;
            this.j = j;
        }

        @Override
        public String toString() {
            return "Position[" +
                    "i=" + i + ", " +
                    "j=" + j + ']';
        }
    }

    private record Velocity(int i, int j) {
    }

    private record Robot(Position position, Velocity velocity) {
    }
}
