package code.advent;

import java.util.List;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Day13 extends ISolution {

    private static final Pattern NUMBERS = Pattern.compile("\\d+");
    private static final int BUTTON_A_COST = 3;
    private static final int BUTTON_B_COST = 1;

    public static void main(String[] args) {
        ISolution solution = new Day13();
        solution.test();
        solution.execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        long tokens = 0;
        List<String> input = inputStream.toList();
        for (int i = 0; i < input.size(); i += 4) {
            List<Long> l1 = NUMBERS.matcher(input.get(i)).results()
                    .map(match -> Long.valueOf(match.group())).toList();
            List<Long> l2 = NUMBERS.matcher(input.get(i + 1)).results()
                    .map(match -> Long.valueOf(match.group())).toList();
            List<Long> constants = NUMBERS.matcher(input.get(i + 2)).results()
                    .map(match -> Long.valueOf(match.group())).toList();

            Equation eq1 = new Equation(l1.get(0), l2.get(0), constants.get(0));
            Equation eq2 = new Equation(l1.get(1), l2.get(1), constants.get(1));

            LinearEquation2d equation2d = LinearEquation2d.parse(eq1, eq2);
            if (equation2d.isSolvable()) {
                Solution2d s = equation2d.solve();
                if (s != null && s.a <= 100 && s.b <= 100) {
                    tokens += s.a * BUTTON_A_COST + s.b * BUTTON_B_COST;
                }
            }
        }

        return String.valueOf(tokens);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        long measurementError = 10000000000000L;
        long tokens = 0;
        List<String> input = inputStream.toList();
        for (int i = 0; i < input.size(); i += 4) {
            List<Long> l1 = NUMBERS.matcher(input.get(i)).results()
                    .map(match -> Long.valueOf(match.group())).toList();
            List<Long> l2 = NUMBERS.matcher(input.get(i + 1)).results()
                    .map(match -> Long.valueOf(match.group())).toList();
            List<Long> constants = NUMBERS.matcher(input.get(i + 2))
                    .results().map(match -> Long.valueOf(match.group())).toList();

            Equation eq1 = new Equation(l1.get(0), l2.get(0),
                    constants.get(0) + measurementError);
            Equation eq2 = new Equation(l1.get(1), l2.get(1),
                    constants.get(1) + measurementError);

            LinearEquation2d equation2d = LinearEquation2d.parse(eq1, eq2);
            if (equation2d.isSolvable()) {
                Solution2d s = equation2d.solve();
                if (s != null) {
                    tokens += s.a * BUTTON_A_COST + s.b * BUTTON_B_COST;
                }
            }
        }

        return String.valueOf(tokens);
    }

    private record Equation(long coefficient1, long coefficient2, long constant) {
    }

    private record Solution2d(long a, long b) {
    }

    private record LinearEquation2d(long[][] data, long[] constants) {

        public static LinearEquation2d parse(Equation eq1, Equation eq2) {
            long[][] data = new long[2][2];
            data[0][0] = eq1.coefficient1;
            data[0][1] = eq1.coefficient2;

            data[1][0] = eq2.coefficient1;
            data[1][1] = eq2.coefficient2;

            long[] constants = new long[2];
            constants[0] = eq1.constant;
            constants[1] = eq2.constant;

            return new LinearEquation2d(data, constants);
        }

        public long determinant() {
            return data[0][0] * data[1][1] - data[0][1] * data[1][0];
        }

        public boolean isSolvable() {
            return determinant() != 0;
        }

        public Solution2d solve() {
            long det = determinant();
            long s1 = data[1][1] * constants[0] - data[0][1] * constants[1];
            long s2 = data[0][0] * constants[1] - data[1][0] * constants[0];
            if (s1 % det == 0 && s2 % det == 0) {
                return new Solution2d(s1 / det, s2 / det);
            }
            return null;
        }
    }
}
