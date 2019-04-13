import java.util.*;
import java.io.*;

class Node {
    Node left, right;
    int data;

    Node(int data) {
        this.data = data;
        left = right = null;
    }
}

class Solution {

    public static int getHeight(Node root) {
        // Write your code here
        if (root == null) {
            return -1;
        }
        ArrayList<Node> blacklisted = new ArrayList();
        ArrayList<Integer> depths = new ArrayList();
        Node node = root;
        int currDepth = 0;
        depths.add(currDepth);
        while (true) {
            if (node == root && !canTraverseLeft(node, blacklisted) && !canTraverseRight(node, blacklisted)) {
                return max(depths);
            }
            if (canTraverseLeft(node, blacklisted)) {
                currDepth++;
                node = node.left;
            } else if (canTraverseRight(node, blacklisted)) {
                currDepth++;
                node = node.right;
            } else {
                blacklisted.add(node);
                depths.add(currDepth);
                currDepth = 0;
                node = root;
            }
        }
    }

    static boolean canTraverseLeft(Node node, ArrayList<Node> blacklisted) {
        return node.left != null && !blacklisted.contains(node.left);
    }

    static boolean canTraverseRight(Node node, ArrayList<Node> blacklisted) {
        return node.right != null && !blacklisted.contains(node.right);
    }

    static int max(ArrayList<Integer> list) {
        int maximum = list.get(0);
        for (int i = 1; i < list.size(); i++) {
            if (list.get(i) > maximum) {
                maximum = list.get(i);
            }
        }
        return maximum;
    }

    public static Node insert(Node root, int data) {
        if (root == null) {
            return new Node(data);
        } else {
            Node cur;
            if (data <= root.data) {
                cur = insert(root.left, data);
                root.left = cur;
            } else {
                cur = insert(root.right, data);
                root.right = cur;
            }
            return root;
        }
    }

    public static void main(String args[]) {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        Node root = null;
        while (T-- > 0) {
            int data = sc.nextInt();
            root = insert(root, data);
        }
        int height = getHeight(root);
        System.out.println(height);
    }
}