package code.advent;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.stream.Stream;

public class Day11 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day11();
        solution.test();
        solution.execute();
    }

    private record State(long num, int depth) {}

    private final HashMap<State, Long> stateMap = new HashMap<>();

    private int getDigitCount(long num) {
        int count = 0;
        while (num > 0) {
            count++;
            num /= 10;
        }
        return count;
    }

    private long getStones(long num, int depth) {
        State s = new State(num, depth);
        if (stateMap.containsKey(s)) {
            return stateMap.get(s);
        }
        if (depth == 0) {
            return 1;
        }
        if (num == 0) {
            long c = getStones(1, depth - 1);
            stateMap.put(new State(1, depth-1), c);
            return c;
        }
        int digitCount = getDigitCount(num);
        if ((digitCount & 1) == 0) {
            long mask = 1;
            for (int i = 0; i < digitCount / 2; i++) {
                mask *= 10;
            }
            long r = getStones(num / mask, depth - 1);
            stateMap.put(new State(num / mask, depth - 1), r);
            long l = getStones(num % mask, depth - 1);
            stateMap.put(new State(num % mask, depth - 1), l);

            return r + l;
        } else {
            long c = getStones(num * 2024L, depth - 1);
            stateMap.put(new State(num * 2024L, depth - 1), c);
            return c;
        }
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        List<Long> nums = Arrays.stream(inputStream.toList().get(0).split(" "))
                .map(Long::parseLong).toList();
        long count = 0;
        for (long num: nums) {
            count += this.getStones(num, 25);
        }
        return String.valueOf(count);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        List<Long> nums = Arrays.stream(inputStream.toList().get(0).split(" "))
                .map(Long::parseLong).toList();
        long count = 0;
        for (long num: nums) {
            count += this.getStones(num, 75);
        }
        return String.valueOf(count);
    }
}
