package code.advent;

import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Day03 extends ISolution {

    public static void main(String[] args) {
        new Day03().execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        long result = 0;
        Pattern pattern = Pattern.compile("mul\\((\\d{1,3}),(\\d{1,3})\\)");
        for (String line : inputStream.toList()) {
            result += pattern.matcher(line).results().mapToLong(m -> Long.parseLong(m.group(1)) * Long.parseLong(m.group(2))).sum();
        }
        return String.valueOf(result);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        long result = 0;
        Pattern pattern = Pattern.compile("mul\\((\\d{1,3}),(\\d{1,3})\\)|don't\\(\\)|do\\(\\)");
        boolean enabled = true;
        for (String line : inputStream.toList()) {
            for (MatchResult m : pattern.matcher(line).results().toList()) {
                if (m.group().equals("do()")) {
                    enabled = true;
                } else if ("don't()".equals(m.group())) {
                    enabled = false;
                } else if (enabled) {
                    result += Long.parseLong(m.group(1)) * Long.parseLong(m.group(2));
                }
            }
        }
        return String.valueOf(result);
    }
}
