package code.advent;

import java.util.Collections;
import java.util.List;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Day07 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day07();
        solution.test();
        solution.execute();
    }

    private enum Operator {
        ADD, MULTIPLY, CONCAT;

        public long apply(long num1, long num2) {
            switch (this) {
                case ADD -> {
                    return num1 + num2;
                }
                case MULTIPLY -> {
                    return num1 * num2;
                }
                case CONCAT -> {
                    long result = num1;
                    long num = num2;
                    while (num > 0) {
                        num /= 10;
                        result *= 10;
                    }
                    return result + num2;
                }
                default -> throw new IllegalArgumentException(this.name());
            }
        }
    }

    private record Calibration(List<Long> nums, long result) {

        private static final Pattern numberPattern = Pattern.compile("\\d+");

        public static Calibration parse(String stringInput) {
            List<Long> nums = numberPattern.matcher(stringInput).results()
                    .map(matchResult -> Long.parseLong(matchResult.group())).toList();
            return new Calibration(Collections.unmodifiableList(nums.subList(1, nums.size())), nums.get(0));
        }

        public boolean valid(Operator... operators) {
            boolean isValid = false;

            for (Operator op : operators) {
                isValid |= validCheckHelper(1, nums.get(0), op);
            }

            return isValid;
        }

        public boolean validCheckHelper(int index, long val, Operator operator) {
            if (index == nums.size()) {
                return val == result;
            } else if (val > result) {
                return false;
            } else {
                boolean result = false;
                val = operator.apply(val, nums.get(index));
                for (Operator op : Operator.values()) {
                    result |= validCheckHelper(index + 1, val, op);
                }
                return result;
            }
        }
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        long result = 0;

        for (String line : inputStream.toList()) {
            Calibration calibration = Calibration.parse(line);
            if (calibration.valid(Operator.ADD, Operator.MULTIPLY)) {
                result += calibration.result;
            }
        }

        return String.valueOf(result);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        long result = 0;

        for (String line : inputStream.toList()) {
            Calibration calibration = Calibration.parse(line);
            if (calibration.valid(Operator.ADD, Operator.MULTIPLY, Operator.CONCAT)) {
                result += calibration.result;
            }
        }

        return String.valueOf(result);
    }
}
