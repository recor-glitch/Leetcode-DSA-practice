import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class helper {
    Scanner sc = new Scanner(System.in);
    Application app = new Application();
    HashMap<String, Node> playlist = new HashMap<>();

    void Task() {
        System.out.println("*****SPOTIFY*****");
        System.out.println();
        System.out.println("Options: ");
        System.out.println();
        System.out.println("1. Create a new Playlist.");
        System.out.println("2. List all the Playlist.");
        System.out.println("3. Select a playlist.");
        System.out.println("5. Exit the Program");
        System.out.println();
        int res = sc.nextInt();
        validate(res);
    }

    void validate(int option) {
        switch (option) {
            case 1:
                addPlaylist();
                break;
            case 2:
                display();
                break;
            case 3:
                String response = getInput(
                        "Enter the name of the playlist, if not remembering than I prefer you to search the list first.");
                selectPlaylist(response);
                break;
            case 5:
                System.exit(0);
                break;
            default:
                break;
        }
    }

    void selectPlaylist(String name) {
        Node node = null;
        boolean inside = true;
        for (Map.Entry<String, Node> entry : playlist.entrySet()) {
            if (entry.getKey().equals(name)) {
                node = entry.getValue();
            }
        }
        if (node != null) {
            while (inside) {
                System.out.println("Available options:");
                System.out.println("1. Play songs.");
                System.out.println("2. Add a song.");
                System.out.println("3. Go to previous menu.");
                switch (sc.nextInt()) {
                    case 1:
                        System.out.println("Playing your playlist " + name + ":");
                        playSongs(node);
                        break;
                    case 2:
                        String res = getInput("Enter the name of the song: ");
                        addSong(res);
                        break;
                    case 3:
                        inside = false;
                        break;
                    default:
                        break;
                }
            }
        } else {
            System.out.println("No matching playlist available.");
        }
    }

    void addSong(String song) {
        Node current = app.head;
        if (current.data == null) {
            current.data = song;
            return;
        }
        while (current != null) {
            if (current.next == app.tail) {
                Node node = new Node(song);
                node.next = app.tail;
                app.tail.prev = node;
                current.next = node;
                node.prev = current;
                break;
            } else {
                current = current.next;
            }
        }
    }

    void playSongs(Node node) {
        boolean toggle = false;
        System.out.println("Want to enable toggle mode (y/n)?");
        String result = sc.next();
        if (result.equals("y")) {
            toggle = true;
        } else {
            toggle = false;
        }
        Node current = app.head;
        Node prev = app.tail;
        if (toggle) {
            while (true) {
                if (current.data == null) {
                    System.out.println("There is no music in your playlist.");
                    break;
                } else {
                    System.out.println("Go next or previous or back (next/prev/back) : ");
                    String res = sc.next();
                    if (res.equals("next")) {
                        System.out.println(current.data);
                        prev = current;
                        current = current.next;
                    }
                    if (res.equals("prev")) {
                        System.out.println(current.data);
                        prev = current.prev;
                    }
                    if (res.equals("back")) {
                        break;
                    }
                }

            }
        } else {
            while (true) {
                if (current.data == null) {
                    System.out.println("There is no music in your playlist.");
                    break;
                } else {
                    System.out.println("Go next or previous or back (next/prev/back) : ");
                    String response = sc.next();
                    if (response.equals("next")) {
                        if (current.equals(app.tail) && current.equals(app.head)) {
                            System.out.println(current.data);
                            return;
                        }
                        if (!current.equals(app.tail)) {
                            System.out.println(current.data);
                            prev = current;
                            current = current.next;
                        } else {
                            System.out.println("There is no more songs in the playlist.");
                            return;
                        }
                    } else if (response.equals("prev")) {
                        if (current.equals(app.tail) && current.equals(app.head)) {
                            System.out.println(current.data);
                            return;
                        }
                        if (!prev.equals(app.head)) {
                            System.out.println(current.data);
                            current = prev;
                            prev = current.prev;
                        } else {
                            System.out.println("No previous song.");
                        }
                    } else if (response.equals("back")) {
                        break;
                    } else {
                        System.out.println("Invalid Input.");
                    }
                }
            }
        }
    }

    void addPlaylist() {
        Node node = null;
        String song = null;
        System.out.println("Enter the name of the playlist: ");
        String name = sc.next();
        System.out.println("Want to add the first song(y/n)?");
        String option = sc.next();
        if (option.equals("y")) {
            System.out.println("Enter the first song in the playlist:");
            song = sc.next();
        }
        if (song != null) {

            node = new Node(song);
        } else {
            node = new Node();
        }
        app.head = node;
        app.tail = node;
        app.head.next = app.tail;
        app.head.prev = app.tail;
        app.tail.prev = app.head;
        app.tail.next = app.head;
        playlist.put(name, node);
    }

    void display() {
        for (Map.Entry<String, Node> entry : playlist.entrySet()) {
            System.out.println(entry.getKey());
        }
    }

    String getInput(String question) {
        System.out.println(question);
        return sc.next();
    }

}
