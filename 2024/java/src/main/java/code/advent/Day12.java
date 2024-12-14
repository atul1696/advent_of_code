package code.advent;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;
import java.util.stream.Stream;

public class Day12 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day12();
        solution.test();
        solution.execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        Garden garden = Garden.parse(inputStream.toList());
        long price = 0;
        for (Plot p : garden.getPlots()) {
            price += p.getFencePriceWithPerimeter();
        }
        return String.valueOf(price);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        Garden garden = Garden.parse(inputStream.toList());
        long price = 0;
        for (Plot p : garden.getPlots()) {
            price += p.getFencePriceWithEdges();
        }
        return String.valueOf(price);
    }

    private static final class Plot {
        long area;
        long perimeter;
        long corners;

        Plot() {
            this.area = 0;
            this.perimeter = 0;
            this.corners = 0;
        }

        long getFencePriceWithPerimeter() {
            return area * perimeter;
        }

        long getFencePriceWithEdges() {
            return area * corners;
        }

        @Override
        public String toString() {
            return "Plot{" +
                    "area=" + area +
                    ", perimeter=" + perimeter +
                    ", corners=" + corners +
                    '}';
        }
    }

    private static class Garden {
        private final char[][] garden;
        private final int rows;
        private final int cols;

        private Garden(char[][] garden) {
            this.garden = garden;
            this.rows = garden.length;
            this.cols = garden[0].length;
        }

        public static Garden parse(List<String> input) {
            int rows = input.size();
            int cols = input.get(0).length();
            char[][] garden = new char[rows][cols];
            for (int i = 0; i < rows; i++) {
                for (int j = 0; j < cols; j++) {
                    garden[i][j] = input.get(i).charAt(j);
                }
            }

            return new Garden(garden);
        }

        private int getCornersAt(int i, int j) {
            int count = 0;
            boolean[][] surround = new boolean[3][3];
            for (int k = 0; k < 3; k++) {
                for (int l = 0; l < 3; l++) {
                    int i2 = i + k - 1;
                    int j2 = j + l - 1;
                    surround[k][l] = i2 >= 0 && i2 < rows && j2 >= 0 && j2 < cols &&
                            garden[i][j] == garden[i2][j2];
                }
            }

            // external
            if (!surround[0][1] && !surround[1][2]) {
                count += 1;
            }
            if (!surround[1][2] && !surround[2][1]) {
                count += 1;
            }
            if (!surround[2][1] && !surround[1][0]) {
                count += 1;
            }
            if (!surround[1][0] && !surround[0][1]) {
                count += 1;
            }

            // internal
            if (!surround[0][0] && surround[0][1] && surround[1][0]) {
                count += 1;
            }
            if (!surround[0][2] && surround[0][1] && surround[1][2]) {
                count += 1;
            }
            if (!surround[2][2] && surround[1][2] && surround[2][1]) {
                count += 1;
            }
            if (!surround[2][0] && surround[1][0] && surround[2][1]) {
                count += 1;
            }
            return count;
        }

        private Plot getPlot(int x, int y, boolean[][] visited) {
            Stack<Loc> stack = new Stack<>();
            stack.add(new Loc(x, y));
            visited[x][y] = true;
            Plot p = new Plot();
            while (!stack.isEmpty()) {
                Loc loc = stack.pop();
                int i = loc.i;
                int j = loc.j;

                p.area += 1;

                if (i - 1 < 0 || garden[i - 1][j] != garden[i][j]) {
                    p.perimeter += 1;
                } else if (!visited[i - 1][j]) {
                    visited[i - 1][j] = true;
                    stack.add(new Loc(i - 1, j));
                }

                if (i + 1 >= rows || garden[i + 1][j] != garden[i][j]) {
                    p.perimeter += 1;
                } else if (!visited[i + 1][j]) {
                    visited[i + 1][j] = true;
                    stack.add(new Loc(i + 1, j));
                }

                if (j - 1 < 0 || garden[i][j - 1] != garden[i][j]) {
                    p.perimeter += 1;
                } else if (!visited[i][j - 1]) {
                    visited[i][j - 1] = true;
                    stack.add(new Loc(i, j - 1));
                }

                if (j + 1 >= cols || garden[i][j + 1] != garden[i][j]) {
                    p.perimeter += 1;
                } else if (!visited[i][j + 1]) {
                    visited[i][j + 1] = true;
                    stack.add(new Loc(i, j + 1));
                }

                p.corners += getCornersAt(i, j);

            }
            return p;
        }

        public List<Plot> getPlots() {
            List<Plot> plotList = new ArrayList<>();
            boolean[][] visited = new boolean[rows][cols];

            for (char c = 'A'; c <= 'Z'; c++) {
                for (int i = 0; i < rows; i++) {
                    for (int j = 0; j < cols; j++) {
                        if (garden[i][j] == c && !visited[i][j]) {
                            plotList.add(getPlot(i, j, visited));
                        }
                    }
                }
            }

            return plotList;
        }

        private record Loc(int i, int j) {
        }
    }
}
