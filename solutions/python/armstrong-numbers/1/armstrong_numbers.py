""" Provides a function that determines if a number is an Armstrong number.

An Armstrong number is a number that is the sum of its own digits,
with each digit raised to the power of the number of digits. 
"""

def is_armstrong_number(number: int) -> bool:
    """ Implements the Armstrong number rule on a number.

    Args:
        number(int): integer to check.
    Returns:
        bool: True if it conforms to the rules of an Armstrong number, False otherwise.
    """
    sum = 0
    for num in list(str(number)):
        sum += int(num) ** len(str(number))
    return sum == number