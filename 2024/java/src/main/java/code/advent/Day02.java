package code.advent;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Stream;

public class Day02 extends ISolution {

    public static void main(String[] args) {
        new Day02().execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        long safeLevels = inputStream.filter(line -> {
            List<Long> report = Arrays.stream(line.split("\\s")).map(Long::parseLong).toList();
            boolean increasing = report.get(1) - report.get(0) > 0;
            for (int i = 1; i < report.size(); i++) {
                long diff = report.get(i) - report.get(i-1);
                if (((diff <= 0 || !increasing) && (diff >= 0 || increasing)) || Math.abs(diff) > 3) {
                    return false;
                }
            }
            return true;
        }).count();

        return String.valueOf(safeLevels);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        long safeLevels = inputStream.filter(line -> {
            List<Long> report = Arrays.stream(line.split("\\s")).map(Long::parseLong).toList();

            for (int index = 0; index <= report.size(); index++) {
                List<Long> r = new LinkedList<>(report);
                if (index < report.size()) {
                    r.remove(index);
                }
                boolean increasing = r.get(1) - r.get(0) > 0;
                boolean isSafe = true;
                for (int i = 1; i < r.size(); i++) {
                    long diff = r.get(i) - r.get(i - 1);
                    if (((diff <= 0 || !increasing) && (diff >= 0 || increasing)) || Math.abs(diff) > 3) {
                        isSafe = false;
                        break;
                    }
                }

                if (isSafe) {
                    return true;
                }
            }

            return false;
        }).count();

        return String.valueOf(safeLevels);
    }
}
