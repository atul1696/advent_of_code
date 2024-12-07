package code.advent;

import java.io.BufferedReader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Objects;
import java.util.function.Function;
import java.util.stream.Stream;

public abstract class ISolution {

    private static final String INPUT_DIR_PATH = "input";
    private static final String SAMPLE_INPUT_DIR_PATH = "sample";

    public void execute() {
        System.out.println("INPUT:");
        run(String.format("%s/%s.txt", INPUT_DIR_PATH, this.getClass().getSimpleName()).toLowerCase());
    }

    public void test() {
        System.out.println("SAMPLE:");
        run(String.format("%s/%s.txt", SAMPLE_INPUT_DIR_PATH, this.getClass().getSimpleName()).toLowerCase());
    }

    private void run(String filename) {
        System.out.println("Part 1: " + solve(this::part1, filename));
        System.out.println("Part 2: " + solve(this::part2, filename));
    }

    protected abstract String part1(Stream<String> inputStream);

    protected abstract String part2(Stream<String> inputStream);

    private String solve(Function<Stream<String>, String> fn, String fileName) {
        try {
            Path path = Paths.get(Objects.requireNonNull(this.getClass().getClassLoader().getResource(fileName)).getPath());
            try (BufferedReader reader = Files.newBufferedReader(path)) {
                return fn.apply(reader.lines());
            } catch (Exception e) {
                System.err.println("ERROR: " + e.getMessage());
                e.printStackTrace();
            }
        } catch (NullPointerException e) {
            System.err.println("ERROR: File not found " + fileName);
        }
        System.exit(1);
        return null;
    }
}
