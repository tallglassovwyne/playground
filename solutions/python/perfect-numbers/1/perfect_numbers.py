""" Provides a function that checks if a number is Perfect, Abundant, or Deficient based on Nicomachus' classification scheme for positive integers.
"""

def get_aliquot(number: int) -> int:
    """Get the aliquot sum of a number.

    The aliquot sum is defined as the sum of the factors of a number not including the number itself.
    Raises an error if the number is less than 1.

    Args:
        number: Positive integer.
    Returns:
        int: sum of factors of the number 
    """
    if number < 1:
        raise ValueError("Classification is only possible for positive integers.")
    total = 0
    for num in range(1, number):
        if number % num == 0:
            total += num
    return total

def classify(number: int) -> str:
    """ Check if a number is Perfect, Abundant, or Deficient.

    A number is perfect when it equals its aliquot sum.
    A number is abundant when it is less than its aliquot sum.
    A number is deficient when it is greater than its aliquot sum

    The aliquot sum is defined as the sum of the factors of a number not including the number itself.

    Args:
        number: A positive integer.
    Returns: 
        str: The classification of the input integer.
    """
    if number == get_aliquot(number): return "perfect"
    if number < get_aliquot(number): return "abundant"
    return "deficient"  
