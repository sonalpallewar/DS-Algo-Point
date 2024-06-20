// Define the Person class
public class Person {
    // Instance variables
    private String name;
    private int age;

    // Constructor to initialize the instance variables
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // Instance method to get the person's name
    public String getName() {
        return name;
    }

    // Instance method to set the person's name
    public void setName(String name) {
        this.name = name;
    }

    // Instance method to get the person's age
    public int getAge() {
        return age;
    }

    // Instance method to set the person's age
    public void setAge(int age) {
        if (age > 0) {
            this.age = age;
        } else {
            System.out.println("Age must be positive.");
        }
    }

    // Instance method to display the person's details
    public void displayDetails() {
        System.out.println("Name: " + name);
        System.out.println("Age: " + age);
    }

    // Static method to compare the ages of two persons
    public static Person olderPerson(Person p1, Person p2) {
        if (p1.getAge() > p2.getAge()) {
            return p1;
        } else {
            return p2;
        }
    }

    // Main method to demonstrate the usage of the Person class
    public static void main(String[] args) {
        // Create two Person objects using the constructor
        Person person1 = new Person("Alice", 30);
        Person person2 = new Person("Bob", 25);

        // Display the details of person1
        person1.displayDetails();

        // Display the details of person2
        person2.displayDetails();

        // Change the age of person2
        person2.setAge(35);

        // Display the updated details of person2
        person2.displayDetails();

        // Compare the ages of person1 and person2
        Person older = Person.olderPerson(person1, person2);
        System.out.println("The older person is: " + older.getName());
    }
}
