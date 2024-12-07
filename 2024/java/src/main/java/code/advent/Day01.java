package code.advent;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;

public class Day01 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day01();
        solution.test();
        solution.execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        long distance = 0;
        List<Long> left = new ArrayList<>();
        List<Long> right = new ArrayList<>();
        inputStream.forEach(line -> {
            String[] parts = line.split("\\s+");
            left.add(Long.parseLong(parts[0]));
            right.add(Long.parseLong(parts[1]));
        });

        left.sort(null);
        right.sort(null);

        for (int i = 0; i < left.size(); i++) {
            distance += Math.abs(left.get(i) - right.get(i));
        }

        return String.valueOf(distance);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        long score = 0;
        List<Long> left = new ArrayList<>();
        Map<Long, Long> counts = new HashMap<>();
        inputStream.forEach(line -> {
            String[] parts = line.split("\\s+");
            left.add(Long.parseLong(parts[0]));
            counts.compute(Long.parseLong(parts[1]), (k, v) -> v == null ? 1 : v + 1);
        });

        for (long num: left) {
            score += num * counts.getOrDefault(num,0L);
        }

        return String.valueOf(score);
    }
}
