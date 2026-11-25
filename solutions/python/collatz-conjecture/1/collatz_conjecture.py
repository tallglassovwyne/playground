""" Provides a function that determines the number of steps it takes for a number to find its way back to 1.
"""

def steps(number: int) -> int:
    """ Uses the Collatz Conjecture to determine the number of steps it takes for a number to find its way back to 1.

    Rules for Collatz Conjecture:
    1. If the number is even, divide it by 2.
    2. If the number is odd, multiply it by 3 and add 1.

    Args:
        number: A positive integer.
    Returns:
        int: Number of steps taken.
    """
    if number < 1:
        raise ValueError("Only positive integers are allowed")
    total_number_of_steps = 0
    while number != 1:
        if number % 2 == 0:
            number //= 2
        else:
            number = (number * 3) + 1
        total_number_of_steps += 1
    return total_number_of_steps