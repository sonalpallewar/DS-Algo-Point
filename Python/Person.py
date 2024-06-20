class Person:
    # Class variable
    population = 0

    def __init__(self, name, age):
        # Instance variables
        self.name = name
        self.age = age
        Person.population += 1

    # Instance method to get the person's name
    def get_name(self):
        return self.name

    # Instance method to set the person's name
    def set_name(self, name):
        self.name = name

    # Instance method to get the person's age
    def get_age(self):
        return self.age

    # Instance method to set the person's age
    def set_age(self, age):
        if age > 0:
            self.age = age
        else:
            print("Age must be positive.")

    # Instance method to display the person's details
    def display_details(self):
        print(f"Name: {self.name}")
        print(f"Age: {self.age}")

    # Class method to get the current population
    @classmethod
    def get_population(cls):
        return cls.population


# Main function to demonstrate the usage of the Person class
def main():
    # Create two Person objects
    person1 = Person("Alice", 30)
    person2 = Person("Bob", 25)

    # Display the details of person1
    person1.display_details()

    # Display the details of person2
    person2.display_details()

    # Change the age of person2
    person2.set_age(35)

    # Display the updated details of person2
    person2.display_details()

    # Display the current population
    print(f"Current population: {Person.get_population()}")

if __name__ == "__main__":
    main()
