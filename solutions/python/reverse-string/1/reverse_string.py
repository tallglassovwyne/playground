""" Provides a function that reverse a string.
"""

def reverse(text: str)-> str:
    """ Reverse a string.

    Args:
        text: String of characters.
    Returns:
        str: Reversed string.
    """
    return "".join([char for char in text[::-1]])