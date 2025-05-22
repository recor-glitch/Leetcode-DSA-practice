public class SortList {
    class Node {

        int data;
        int count;
        Node next;

        public Node(int data) {
            this.data = data;
            this.count = 1;
            this.next = null;
        }

    }

    public Node head = null;
    public Node tail = null;

    void display() {
        Node current = head;

        if (head == null) {
            System.out.println("List is empty");
            return;
        }
        while (current != null) {
            System.out.println("Data: " + current.data + " " + "count: " + current.count);
            current = current.next;
        }

        System.out.println();
    }

    void sortIncreasingOrder() {
        Node current = head, index = null;
        int temp;

        if (head == null) {
            System.out.println("Nothing to sort.");
            return;
        } else {
            while (current != null) {
                index = current.next;

                while (index != null) {
                    if (current.data > index.data) {
                        temp = current.data;
                        current.data = index.data;
                        index.data = temp;
                    }
                    index = index.next;
                }
                current = current.next;
            }
            System.out.println("Successfully Sorted.");
        }
    }

    void sortDecreasingOrder() {
        Node current = head, index = null;
        int temp;

        if (head == null) {
            System.out.println("Nothing to sort.");
            return;
        } else {
            while (current != null) {
                index = current.next;

                while (index != null) {
                    if (current.data < index.data) {
                        temp = current.data;
                        current.data = index.data;
                        index.data = temp;
                    }
                    index = index.next;
                }
                current = current.next;
            }
            System.out.println("Successfully Sorted.");
        }
    }

    void addNode(int data) {
        Node current = head;
        int i = 0;

        if (current == null) {
            System.out.println("Inside null current!");
            Node newNode = new Node(data);
            head = newNode;
            tail = newNode;
        } else {
            System.out.println("Inside not null current!");
            while (current != null) {
                if (current.data == data) {
                    current.count = current.count + 1;
                    i++;
                }
                current = current.next;
            }
            if (i == 0) {
                Node newNode = new Node(data);
                tail.next = newNode;
                tail = newNode;
            }
        }
    }

    void removeNode(int data) {
        Node temp = head, prev = null;
        if (temp != null && temp.data == data) {
            if (temp.count > 1) {
                temp.count = temp.count - 1;
                return;
            } else {
                head = temp.next;
                return;
            }
        }
        while (temp != null && temp.data != data) {
            prev = temp;
            temp = temp.next;
        }

        if (temp == null)
            return;

        prev.next = temp.next;
    }
}