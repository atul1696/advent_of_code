package code.advent;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

public class Day09 extends ISolution {

    public static void main(String[] args) {
        ISolution solution = new Day09();
        solution.test();
        solution.execute();
    }

    @Override
    protected String part1(Stream<String> inputStream) {
        String input = inputStream.toList().get(0);
        long checksum = 0;
        int length = input.length();

        int[] fileBlocks = new int[(length + 1) / 2];
        int[] freeBlocks = new int[length / 2];

        for (int i = 0; i < length; i++) {
            int num = input.charAt(i) - '0';
            if ((i & 1) == 0) {
                fileBlocks[i / 2] = num;
            } else {
                freeBlocks[i / 2] = num;
            }
        }

        long index = 0;
        int right = fileBlocks.length - 1;
        for (int i = 0; i < fileBlocks.length; i++) {
            for (int j = 0; j < fileBlocks[i]; j++) {
                checksum += index * i;
                index++;
            }

            for (int j = 0; i < right && j < freeBlocks[i]; j++) {
                checksum += index * right;
                index++;
                fileBlocks[right] -= 1;
                if (fileBlocks[right] == 0) {
                    right -= 1;
                }
            }
        }

        return String.valueOf(checksum);
    }

    @Override
    protected String part2(Stream<String> inputStream) {
        String input = inputStream.toList().get(0);
        int length = input.length();

        List<Integer> indexLList = new ArrayList<>();
        List<Integer> fileBlockLList = new ArrayList<>();
        List<Integer> freeBlockLList = new ArrayList<>();

        for (int i = 0; i < length; i++) {
            int num = input.charAt(i) - '0';
            if ((i & 1) == 0) {
                indexLList.add(i / 2);
                fileBlockLList.add(num);
            } else {
                freeBlockLList.add(num);
            }
        }

        if (freeBlockLList.size() < fileBlockLList.size()) {
            freeBlockLList.add(0);
        }

        for (int i = fileBlockLList.size() - 1; i > 0; i--) {
            for (int j = 0; j < i; j++) {
                if (fileBlockLList.get(i) <= freeBlockLList.get(j)) {
                    freeBlockLList.set(i - 1, freeBlockLList.get(i - 1) +
                            freeBlockLList.get(i) + fileBlockLList.get(i));
                    freeBlockLList.set(j, freeBlockLList.get(j) - fileBlockLList.get(i));

                    freeBlockLList.add(j, 0);
                    freeBlockLList.remove(i + 1);

                    fileBlockLList.add(j + 1, fileBlockLList.get(i));
                    fileBlockLList.remove(i + 1);
                    indexLList.add(j + 1, indexLList.get(i));
                    indexLList.remove(i + 1);
                    i++;
                    break;
                }
            }
        }

        long checksum = 0;
        int index = 0;
        for (int i = 0; i < fileBlockLList.size(); i++) {
            for (int j = 0; j < fileBlockLList.get(i); j++) {
                checksum += (long) indexLList.get(i) * index;
                index++;
            }
            for (int j = 0; j < freeBlockLList.get(i); j++) {
                index++;
            }
        }

        return String.valueOf(checksum);
    }
}
