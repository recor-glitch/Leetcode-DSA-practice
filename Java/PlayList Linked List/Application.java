class Application {

    public Node head = null;
    public Node tail = null;
}

class Node {
    String data;
    Node next;
    Node prev;

    Node() {
        this.data = null;
        next = null;
        prev = null;
    }

    Node(String data) {
        this.data = data;
        this.next = null;
        this.prev = null;
    }
}