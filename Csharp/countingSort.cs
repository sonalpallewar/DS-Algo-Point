using System;

public class CountingSort
{
    public static void countingSort_letters(char [] letters)
    {
        int length = characters.Length;
        char[] sorted = new char[length];
        int[] count = new int[256];

        // Count frequencies of characters
        foreach (char c in characters)
            count[c]++;

        // Calculate positions of characters in sorted array
        for (int i = 1; i < 256; i++)
            count[i] += count[i - 1];

        // Build the output character array
        for (int i = length - 1; i >= 0; i--)
        {
            sorted[count[characters[i]] - 1] = characters[i];
            count[characters[i]]--;
        }

        // Display sorted characters
        Console.Write("Sorted Characters: ");
        foreach (char c in sorted)
            Console.Write(c + " , ");
        Console.WriteLine();
    }

    // Method to perform counting sort for integers
    public static void SortIntegers(int[] integers)
    {
        int length = integers.Length;
        int[] sorted = new int[length];

        // Find the minimum and maximum integers
        int min = integers[0];
        int max = integers[0];
        foreach (int num in integers)
        {
            if (num < min) min = num;
            if (num > max) max = num;
        }

        // Initialize count array
        int[] count = new int[max - min + 1];

        // Count frequencies of integers
        foreach (int num in integers)
            count[num - min]++;

        // Recalculate counts to get positions
        for (int i = 1; i < count.Length; i++)
            count[i] += count[i - 1];

        // Build the output integer array
        for (int i = length - 1; i >= 0; i--)
        {
            sorted[count[integers[i] - min] - 1] = integers[i];
            count[integers[i] - min]--;
        }

        // Display sorted integers
        Console.WriteLine("Sorted Integers:");
        foreach (int num in sorted)
            Console.Write(num + " , ");
        Console.WriteLine();
    }
   public static void PrintCharacters(char[] characters)
    {
        foreach (char c in characters)
            Console.Write(c + " , ");
    }
    // Main method for testing
    public static void Main()
    {
        Console.WriteLine("Welcome to Counting Sort Algorithm\n");
        Console.WriteLine("What would you like to sort today?");
        Console.WriteLine("1. Characters");
        Console.WriteLine("2. Integers\n");

        Console.Write("Enter your choice: ");
        int choice = Convert.ToInt32(Console.ReadLine());

        if (choice == 1)
        {
            Console.Write("Enter characters: ");
            char[] characters = Console.ReadLine().ToCharArray();
            countingSort_letters(characters);
        }
        else if (choice == 2)
        {
            Console.Write("How many integers? ");
            int count = Convert.ToInt32(Console.ReadLine());

            int[] integers = new int[count];
            Console.WriteLine("Enter integers:");
            for (int i = 0; i < count; i++)
                integers[i] = Convert.ToInt32(Console.ReadLine());

            SortIntegers(integers);
        }
        else
        {
            Console.WriteLine("Invalid choice!");
        }
    }
}
