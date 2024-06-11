using System;

class AdditionProgram
{
    static void Main()
    {
        // Prompt the user to enter the first number
        Console.Write("Enter the first number: ");
        string input1 = Console.ReadLine();

        // Prompt the user to enter the second number
        Console.Write("Enter the second number: ");
        string input2 = Console.ReadLine();

        // Parse the input strings to integers
        if (int.TryParse(input1, out int number1) && int.TryParse(input2, out int number2))
        {
            // Perform the addition
            int sum = number1 + number2;

            // Display the result
            Console.WriteLine($"The sum of {number1} and {number2} is: {sum}");
        }
        else
        {
            // Display an error message if inputs are not valid integers
            Console.WriteLine("Invalid input. Please enter valid integers.");
        }
    }
}
