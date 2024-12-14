package code.advent;

import java.util.*;
import java.util.stream.Stream;

public class Day08 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day08();
        solution.test();
        solution.execute();
    }

    public int gcd(int a, int b) {
        if (b == 0) return a;
        return gcd(b, a % b);
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        Grid grid = new Grid(inputStream.toList());
        Map<Character, List<GridLoc>> antennaMap = grid.getAntennaMap();
        Set<GridLoc> antinodeLocs = new HashSet<>();
        for (Map.Entry<Character, List<GridLoc>> entry : antennaMap.entrySet()) {
            List<GridLoc> locations = entry.getValue();
            int n = locations.size();
            for (int i = 0; i < n; i++) {
                GridLoc a = locations.get(i);
                for (int j = i + 1; j < n; j++) {
                    GridLoc b = locations.get(j);
                    int di = b.i - a.i;
                    int dj = b.j - a.j;

                    if (grid.isInBounds(b.i + di, b.j + dj)) {
                        antinodeLocs.add(new GridLoc(b.i + di, b.j + dj));
                    }
                    if (grid.isInBounds(a.i - di, a.j - dj)) {
                        antinodeLocs.add(new GridLoc(a.i - di, a.j - dj));
                    }
                }
            }
        }

        return String.valueOf(antinodeLocs.size());
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        Grid grid = new Grid(inputStream.toList());

        Map<Character, List<GridLoc>> antennaMap = grid.getAntennaMap();
        Set<GridLoc> antinodeLocs = new HashSet<>();
        for (Map.Entry<Character, List<GridLoc>> entry : antennaMap.entrySet()) {
            List<GridLoc> locations = entry.getValue();
            int n = locations.size();
            for (int i = 0; i < n; i++) {
                GridLoc a = locations.get(i);
                for (int j = i + 1; j < n; j++) {
                    GridLoc b = locations.get(j);
                    int di = b.i - a.i;
                    int dj = b.j - a.j;

                    int d = gcd(di, dj);
                    di /= d;
                    dj /= d;

                    int pi = b.i;
                    int pj = b.j;

                    while (grid.isInBounds(pi, pj)) {
                        antinodeLocs.add(new GridLoc(pi, pj));
                        pi += di;
                        pj += dj;
                    }

                    pi = b.i;
                    pj = b.j;
                    while (grid.isInBounds(pi, pj)) {
                        antinodeLocs.add(new GridLoc(pi, pj));
                        pi -= di;
                        pj -= dj;
                    }
                }
            }
        }

        return String.valueOf(antinodeLocs.size());
    }

    private record GridLoc(int i, int j) {
    }

    private record Grid(List<String> grid) {
        public int rows() {
            return grid.size();
        }

        public int cols() {
            return grid.get(0).length();
        }

        public boolean isInBounds(int i, int j) {
            return i >= 0 && i < this.rows() &&
                    j >= 0 && j < this.cols();
        }

        public Map<Character, List<GridLoc>> getAntennaMap() {
            Map<Character, List<GridLoc>> antennaMap = new HashMap<>();
            for (int j = 0; j < cols(); j++) {
                for (int i = 0; i < rows(); i++) {
                    char c = grid.get(i).charAt(j);
                    if (!Character.isLetterOrDigit(c)) {
                        continue;
                    }
                    if (!antennaMap.containsKey(c)) {
                        antennaMap.put(c, new ArrayList<>());
                    }
                    antennaMap.get(c).add(new GridLoc(i, j));
                }
            }
            return antennaMap;
        }
    }
}
