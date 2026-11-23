"""Provides a function that determines if a number is an Armstrong number.

An Armstrong number is a number that is the sum of its own digits,
with each digit raised to the power of the number of digits. 
"""

def is_armstrong_number(number: int) -> bool:
    """Checks whether a number is an Armstrong number.

    Args:
        number: integer to check.
    Returns:
        bool: True if it conforms to the rules of an Armstrong number, False otherwise.
    """
    digits = str(number)
    power = len(str(number))
    return number == sum(int(digit) ** power for digit in digits)