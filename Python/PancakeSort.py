def fibonacci_iterative(n):
    # Check for the first two Fibonacci numbers
    if n <= 0:
        return []
    elif n == 1:
        return [0]
    elif n == 2:
        return [0, 1]

    # Initialize the first two Fibonacci numbers
    fib_sequence = [0, 1]

    # Generate the Fibonacci sequence iteratively
    for i in range(2, n):
        next_number = fib_sequence[-1] + fib_sequence[-2]
        fib_sequence.append(next_number)

    return fib_sequence

# Main function to demonstrate the usage of the fibonacci_iterative function
def main():
    n = 10  # Example: Generate first 10 Fibonacci numbers
    result = fibonacci_iterative(n)
    print(f"The first {n} Fibonacci numbers (iterative approach) are: {result}")

if __name__ == "__main__":
    main()
