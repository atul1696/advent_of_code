package code.advent;

import java.util.*;
import java.util.stream.Stream;

public class Day05 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day05();
        solution.test();
        solution.execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        List<String> input = inputStream.toList();
        Map<String, List<String>> pageAfterMap = new HashMap<>();

        int i = 0;
        for (; i < input.size(); i++) {
            String line = input.get(i);
            if (line.isBlank()) {
                break;
            }
            String[] points = line.split("\\|");
            pageAfterMap.compute(points[0], (s, strings) -> {
                if (Objects.isNull(strings)) {
                    strings = new ArrayList<>();
                }
                strings.add(points[1]);
                return strings;
            });
        }

        int result = 0;
        i += 1;
        for (; i < input.size(); i++) {
            String line = input.get(i);
            String[] pages = line.split(",");
            Set<String> visited = new HashSet<>();
            boolean isCorrectlyOrdered = true;

            for (int p = pages.length - 1; p >= 0; p--) {
                String page = pages[p];
                if (!visited.contains(page)) {
                    visited.addAll(pageAfterMap.getOrDefault(page, Collections.emptyList()));
                } else {
                    isCorrectlyOrdered = false;
                    break;
                }
            }

            if (isCorrectlyOrdered) {
                result += Integer.parseInt(pages[pages.length / 2]);
            }
        }

        return String.valueOf(result);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        List<String> input = inputStream.toList();
        Map<String, Set<String>> pageAfterMap = new HashMap<>();

        int i = 0;
        for (; i < input.size(); i++) {
            String line = input.get(i);
            if (line.isBlank()) {
                break;
            }
            String[] points = line.split("\\|");
            pageAfterMap.compute(points[0], (s, strings) -> {
                if (Objects.isNull(strings)) {
                    strings = new HashSet<>();
                }
                strings.add(points[1]);
                return strings;
            });
        }

        int result = 0;
        i += 1;
        for (; i < input.size(); i++) {
            String line = input.get(i);
            List<String> pages = Arrays.asList(line.split(","));
            Set<String> visited = new HashSet<>();
            boolean isCorrectlyOrdered = true;

            for (int p = pages.size() - 1; p >= 0; p--) {
                String page = pages.get(p);
                if (!visited.contains(page)) {
                    visited.addAll(pageAfterMap.getOrDefault(page, Collections.emptySet()));
                } else {
                    isCorrectlyOrdered = false;
                    break;
                }
            }

            if (!isCorrectlyOrdered) {
                pages.sort((p1, p2) -> {
                    if (p1.equals(p2)) {
                        return 0;
                    }
                    return (pageAfterMap.getOrDefault(p1, Collections.emptySet()).contains(p2)) ? -1 : 1;
                });

                result += Integer.parseInt(pages.get(pages.size() / 2));
            }
        }

        return String.valueOf(result);
    }
}
