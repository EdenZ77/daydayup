package class03;

import java.util.HashMap;
import java.util.HashSet;
import java.util.TreeMap;

public class HashMapAndSortedMap {

    public static class Node {
        public int value;

        public Node(int v) {
            value = v;
        }
    }

    public static class Zuo {
        public int value;

        public Zuo(int v) {
            value = v;
        }
    }

    public static void main(String[] args) {

        HashMap<Integer, String> test = new HashMap<>();
        Integer a = 19000000;
        Integer b = 19000000;
        System.out.println(a == b); // false,使用==比较的是地址

        test.put(a, "我是3");
        System.out.println(test.containsKey(b)); // true，使用equals比较的是值

        Zuo z1 = new Zuo(1);
        Zuo z2 = new Zuo(1);
        HashMap<Zuo, String> test2 = new HashMap<>();
        test2.put(z1, "我是z1");
        System.out.println(test2.containsKey(z2)); // false，使用equals比较的是值

        // UnSortedMap 哈希表，增、删、改、查，在使用时，O（1）
        System.out.println("=====================");


        // TreeMap 有序表：接口名
        // 红黑树、avl、sb树、跳表
        // O(logN)
        System.out.println("有序表测试开始");
        TreeMap<Integer, String> treeMap = new TreeMap<>();

        treeMap.put(3, "我是3");
        treeMap.put(4, "我是4");
        treeMap.put(8, "我是8");
        treeMap.put(5, "我是5");
        treeMap.put(7, "我是7");
        treeMap.put(1, "我是1");
        treeMap.put(2, "我是2");

        System.out.println(treeMap.containsKey(1)); // true
        System.out.println(treeMap.containsKey(10)); // false

        System.out.println(treeMap.get(4)); // 我是4
        System.out.println(treeMap.get(10)); // null

        treeMap.put(4, "他是4");
        System.out.println(treeMap.get(4)); // 他是4

        // treeMap.remove(4);
        System.out.println(treeMap.get(4));

        System.out.println("新鲜：=====================");

        System.out.println(treeMap.firstKey()); // 返回最小的key
        System.out.println(treeMap.lastKey());  // 返回最大的key
        // <= 4
        System.out.println(treeMap.floorKey(4)); // 返回小于等于4的最大的key
        // >= 4
        System.out.println(treeMap.ceilingKey(4)); // 返回大于等于4的最小的key
        // O(logN)
    }
}