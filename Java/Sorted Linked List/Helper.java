import java.util.Scanner;

class Helper {
    Scanner sc = new Scanner(System.in);
    SortList sl = new SortList();

    void Task() {
        System.out.println("Hello This is a Link list problem!");
        System.out.println("What you want to do?");
        System.out.println("Options:");
        System.out.println("1. Sort in increasing order.");
        System.err.println("-1. Sort in decreasing order.");
        System.out.println("2. Display the list.");
        System.out.println("3. Remove item.");
        System.out.println("4. Insert Item.");
        System.out.println("5. Exit the program.");
        String res = sc.next();
        validateInput(res);
    }

    void validateInput(String input) {
        switch (input) {
            case "1":
                sl.sortIncreasingOrder();
                break;
            case "-1":
                sl.sortDecreasingOrder();
                break;
            case "2":
                sl.display();
                break;
            case "3":
            int remove_res = getInput("Enter the element to remove: ");
            sl.removeNode(remove_res);
                break;
            case "4":
                int res = getInput("Enter the element: ");
                sl.addNode(res);
                break;
                case "5":
                dispose();
                System.exit(0);
                break;
            default:
                System.out.println("Invalid Input");
                break;
        }
    }

    int getInput(String question) {
        System.out.print(question);
        return sc.nextInt();
    }

    void dispose() {
        sc.close();
    }
}