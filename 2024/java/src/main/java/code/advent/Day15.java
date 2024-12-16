package code.advent;

import java.util.List;
import java.util.Stack;
import java.util.stream.Stream;

public class Day15 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day15();
        solution.test();
        solution.execute();
    }

    private static class Loc {
        public int i, j;

        public Loc(int i, int j) {
            this.i = i;
            this.j = j;
        }

        @Override
        public String toString() {
            return "Loc{" +
                    "i=" + i +
                    ", j=" + j +
                    '}';
        }
    }

    private record Range(int index, int left, int right) {
    }

    private record Warehouse(int rows, int cols, char[][] map, Loc currentLoc) {

        static Warehouse parse(List<String> input) {
            int rows = input.size();
            int cols = input.get(0).length();
            char[][] map = new char[rows][];
            Loc currentLoc = null;
            for (int i = 0; i < rows; i++) {
                map[i] = input.get(i).toCharArray();
                for (int j = 0; j < cols; j++) {
                    if (map[i][j] == '@') {
                        currentLoc = new Loc(i, j);
                    }
                }
            }
            if (currentLoc == null) {
                throw new IllegalStateException("No start location");
            }
            return new Warehouse(rows, cols, map, currentLoc);
        }

        void moveRobot(char dir) {
            switch (dir) {
                case '>':
                    moveRight();
                    break;
                case '<':
                    moveLeft();
                    break;
                case 'v':
                    moveDown();
                    break;
                case '^':
                    moveUp();
                    break;
                default:
                    throw new IllegalArgumentException(String.valueOf(dir));
            }
        }

        void moveRobot2(char dir) {
            switch (dir) {
                case '>':
                    moveRight();
                    break;
                case '<':
                    moveLeft();
                    break;
                case 'v':
                    moveDown2();
                    break;
                case '^':
                    moveUp2();
                    break;
                default:
                    throw new IllegalArgumentException(String.valueOf(dir));
            }
        }

        private void moveRight() {
            int j = currentLoc.j;
            while (map[currentLoc.i][j] != '.') {
                if (map[currentLoc.i][j] == '#') {
                    return;
                }
                j++;
            }
            for (int k = j; k > currentLoc.j; k--) {
                map[currentLoc.i][k] = map[currentLoc.i][k - 1];
            }
            map[currentLoc.i][currentLoc.j] = '.';
            currentLoc.j += 1;
        }

        private void moveLeft() {
            int j = currentLoc.j;
            while (map[currentLoc.i][j] != '.') {
                if (map[currentLoc.i][j] == '#') {
                    return;
                }
                j--;
            }
            for (int k = j; k < currentLoc.j; k++) {
                map[currentLoc.i][k] = map[currentLoc.i][k + 1];
            }
            map[currentLoc.i][currentLoc.j] = '.';
            currentLoc.j -= 1;
        }

        private void moveDown() {
            int i = currentLoc.i;
            while (map[i][currentLoc.j] != '.') {
                if (map[i][currentLoc.j] == '#') {
                    return;
                }
                i++;
            }
            for (int k = i; k > currentLoc.i; k--) {
                map[k][currentLoc.j] = map[k - 1][currentLoc.j];
            }
            map[currentLoc.i][currentLoc.j] = '.';
            currentLoc.i += 1;
        }

        private void moveUp() {
            int i = currentLoc.i;
            while (map[i][currentLoc.j] != '.') {
                if (map[i][currentLoc.j] == '#') {
                    return;
                }
                i--;
            }
            for (int k = i; k < currentLoc.i; k++) {
                map[k][currentLoc.j] = map[k + 1][currentLoc.j];
            }
            map[currentLoc.i][currentLoc.j] = '.';
            currentLoc.i -= 1;

        }

        private void moveDown2() {
            Stack<Range> stack = new Stack<>();

            int i = currentLoc.i;
            int left = currentLoc.j, right = currentLoc.j;

            while (true) {
                stack.add(new Range(i + 1, left, right));
                if (map[i + 1][left] == ']') {
                    left -= 1;
                }
                if (map[i + 1][right] == '[') {
                    right += 1;
                }

                i += 1;

                boolean isEmpty = true, hasFreeSpace = false;
                for (int k = left; k <= right; k++) {
                    if (map[i][k] == '#') {
                        return;
                    }
                    if (map[i][k] != '.') {
                        isEmpty = false;
                    } else {
                        hasFreeSpace = true;
                    }
                }

                if (!isEmpty) {
                    if (!hasFreeSpace) {
                        // row is full of boxes
                        continue;
                    }
                    while (left < right && map[i][left] == '.') {
                        left++;
                    }
                    while (right > left && map[i][right] == '.') {
                        right--;
                    }
                } else {
                    // found required space, push
                    while (!stack.isEmpty()) {
                        Range r = stack.pop();
                        for (int j = r.left; j <= r.right; j++) {
                            map[r.index][j] = map[r.index - 1][j];
                            map[r.index - 1][j] = '.';
                        }
                    }
                    currentLoc.i += 1;
                    return;
                }
            }
        }

        private void moveUp2() {
            Stack<Range> stack = new Stack<>();

            int i = currentLoc.i;
            int left = currentLoc.j, right = currentLoc.j;

            while (true) {
                stack.add(new Range(i - 1, left, right));
                if (map[i - 1][left] == ']') {
                    left -= 1;
                }
                if (map[i - 1][right] == '[') {
                    right += 1;
                }

                i -= 1;

                boolean isEmpty = true, hasFreeSpace = false;
                for (int k = left; k <= right; k++) {
                    if (map[i][k] == '#') {
                        return;
                    }
                    if (map[i][k] != '.') {
                        isEmpty = false;
                    } else {
                        hasFreeSpace = true;
                    }
                }

                if (!isEmpty) {
                    if (!hasFreeSpace) {
                        // row is full of boxes
                        continue;
                    }
                    while (left < right && map[i][left] == '.') {
                        left++;
                    }
                    while (right > left && map[i][right] == '.') {
                        right--;
                    }
                } else {
                    // found required space, push
                    while (!stack.isEmpty()) {
                        Range r = stack.pop();
                        for (int j = r.left; j <= r.right; j++) {
                            map[r.index][j] = map[r.index + 1][j];
                            map[r.index + 1][j] = '.';
                        }
                    }
                    currentLoc.i -= 1;
                    return;
                }
            }
        }

        long coordinateSum(char box) {
            long sum = 0;
            for (int i = 0; i < rows; i++) {
                for (int j = 0; j < cols; j++) {
                    if (map[i][j] == box) {
                        sum += 100L * i + j;
                    }
                }
            }
            return sum;
        }
    }

    private String transform(String input) {
        StringBuilder sb = new StringBuilder(input.length() * 2);
        for (char c : input.toCharArray()) {
            switch (c) {
                case '#':
                    sb.append("##");
                    break;
                case '.':
                    sb.append("..");
                    break;
                case 'O':
                    sb.append("[]");
                    break;
                case '@':
                    sb.append("@.");
                    break;
                default:
                    throw new IllegalArgumentException(String.valueOf(c));
            }
        }
        return sb.toString();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        List<String> input = inputStream.toList();
        int splitIndex = 0;
        while (splitIndex < input.size() && !input.get(splitIndex).isBlank()) {
            splitIndex++;
        }
        Warehouse warehouse = Warehouse.parse(input.subList(0, splitIndex));

        for (int i = splitIndex + 1; i < input.size(); i++) {
            String moves = input.get(i);
            for (Character direction : moves.toCharArray()) {
                warehouse.moveRobot(direction);
            }
        }

        return String.valueOf(warehouse.coordinateSum('O'));
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        List<String> input = inputStream.toList();
        int splitIndex = 0;
        while (splitIndex < input.size() && !input.get(splitIndex).isBlank()) {
            splitIndex++;
        }
        Warehouse warehouse = Warehouse.parse(input.subList(0, splitIndex)
                .stream().map(this::transform).toList());

        for (int i = splitIndex + 1; i < input.size(); i++) {
            String moves = input.get(i);
            for (Character direction : moves.toCharArray()) {
                warehouse.moveRobot2(direction);
            }
        }

        return String.valueOf(warehouse.coordinateSum('['));
    }
}
