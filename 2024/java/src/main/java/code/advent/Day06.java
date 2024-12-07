package code.advent;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Stream;

public class Day06 extends ISolution {

    private static final char OBSTACLE = '#';
    private static final char START = '^';

    public static void main(String[] args) {
        ISolution solution = new Day06();
        solution.test();
        solution.execute();
    }

    private enum GridDir {
        UP(-1, 0), DOWN(1, 0), LEFT(0, -1), RIGHT(0, 1);

        public final int i, j;

        GridDir(int i, int j) {
            this.i = i;
            this.j = j;
        }

        public GridDir rightTurn() {
            switch (this) {
                case UP -> {
                    return RIGHT;
                }
                case RIGHT -> {
                    return DOWN;
                }
                case DOWN -> {
                    return LEFT;
                }
                case LEFT -> {
                    return UP;
                }
                default -> throw new IllegalArgumentException(this.name());
            }
        }
    }

    private record GridLoc(int i, int j) {
        public GridLoc move(GridDir dir) {
            return new GridLoc(this.i + dir.i, this.j + dir.j);
        }
    }

    private record Grid(List<String> grid) {
        public int rows() {
            return grid.size();
        }

        public int cols() {
            return grid.get(0).length();
        }

        public boolean isNotInBounds(GridLoc loc) {
            return loc.i < 0 || loc.i >= this.rows() ||
                    loc.j < 0 || loc.j >= this.cols();
        }

        public GridLoc getStartLoc() {
            for (int i = 0; i < rows(); i++) {
                for (int j = 0; j < cols(); j++) {
                    if (grid.get(i).charAt(j) == START) {
                        return new GridLoc(i, j);
                    }
                }
            }
            throw new IllegalStateException("No start location found");
        }

        public char atPos(GridLoc loc) {
            return grid.get(loc.i).charAt(loc.j);
        }

        public Set<GridLoc> traverse() {
            GridLoc pos = getStartLoc();
            GridDir dir = GridDir.UP;
            Set<GridLoc> visited = new HashSet<>();

            while (true) {
                visited.add(pos);
                GridLoc nextPos = pos.move(dir);

                if (this.isNotInBounds(nextPos)) {
                    break;
                }

                if (this.atPos(nextPos) == OBSTACLE) {
                    dir = dir.rightTurn();
                    nextPos = pos.move(dir);
                }
                pos = nextPos;
            }
            return visited;
        }

        record Visit(GridDir dir, GridLoc loc) {}

        public boolean hasCycle(GridLoc loc) {
            String original = grid.get(loc.i);

            StringBuilder builder = new StringBuilder(original);
            builder.setCharAt(loc.j, OBSTACLE);
            String updated = builder.toString();
            grid.set(loc.i, updated);

            boolean result = true;

            GridLoc pos = getStartLoc();
            GridDir dir = GridDir.UP;
            Set<Visit> visited = new HashSet<>();

            while (true) {
                Visit visit = new Visit(dir, pos);
                if (visited.contains(visit)) {
                    break;
                }
                visited.add(visit);
                GridLoc nextPos = pos.move(dir);

                if (this.isNotInBounds(nextPos)) {
                    result = false;
                    break;
                }

                if (this.atPos(nextPos) == OBSTACLE) {
                    dir = dir.rightTurn();
                } else {
                    pos = nextPos;
                }
            }
            grid.set(loc.i, original);
            return result;
        }
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        Grid grid = new Grid(inputStream.toList());
        Set<GridLoc> visited = grid.traverse();
        return String.valueOf(visited.size());
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        Grid grid = new Grid(new ArrayList<>(inputStream.toList()));
        Set<GridLoc> visited = grid.traverse();

        int result = 0;
        for (GridLoc loc: visited) {
            if (loc.equals(grid.getStartLoc())) {
                continue;
            }
            if (grid.hasCycle(loc)) {
                result += 1;
            }
        }
        return String.valueOf(result);
    }
}
