import java.io.*;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

class Home {
    public static void main(String[] args) throws IOException {
        int n = 0;
        int i = 0;
        int entry_iteration = 0;
        int receiver_amt = -1;
        String donner = "";
        int donation_no = -1;
        HashMap<String, Double> persons = new HashMap<>();

        File file = new File("./gift.in.txt");

        Scanner sc = new Scanner(file);

        while (sc.hasNextLine()) {
            if (i == 0) {
                n = Integer.parseInt(sc.nextLine());
            } else if (i <= n) {
                String name = sc.nextLine();
                persons.put(name, 0.0);
            } else if (i > n) {
                if (donner == "") {
                    donner = sc.nextLine();
                } else if (donner != "" && receiver_amt == -1 && donation_no == -1) {
                    if (receiver_amt == 0) {
                        continue;
                    }
                    String[] amt = sc.nextLine().split(" ");
                    receiver_amt = Integer.parseInt(amt[0]);
                    donation_no = Integer.parseInt(amt[1]);
                    double d_amt = persons.get(donner);
                    persons.put(donner, d_amt - receiver_amt);
                } else if (donner != "" && receiver_amt != -1 && donation_no != -1) {
                    if (entry_iteration < donation_no && donation_no != 0) {
                        if (receiver_amt % donation_no != 0 && entry_iteration == 0) {
                            double prev_amt = persons.get(donner);
                            persons.put(donner, prev_amt + (receiver_amt % donation_no));
                        }
                        Double acc_amt = (double) (receiver_amt / donation_no);
                        String person = sc.nextLine();
                        double amt = persons.get(person);
                        amt = amt + acc_amt;
                        persons.put(person, amt);
                        entry_iteration++;
                    } else {
                        donation_no = -1;
                        receiver_amt = -1;
                        donner = "";
                        entry_iteration = 0;
                    }
                }
            }
            i++;
        }
        sc.close();
        System.out.println(persons);
        BufferedWriter bw = new BufferedWriter(new FileWriter("./gift.out.txt"));
        for (Map.Entry<String, Double> entry : persons.entrySet()) {
            bw.write(entry.getKey() + " " + entry.getValue());
            bw.newLine();
        }
        bw.close();
    }
}