def convert(number: int) -> str:
    """Convert a number to its corresponding raindrop sound.

    Args:
        number: integer to convert.
    Returns:
        str: raindrop sound.
    """
    result = ""
    if number % 3 == 0:
        result += "Pling" 
    if number % 5 == 0:
        result += "Plang"
    if number % 7 == 0:
        result += "Plong"

    if len(result) <= 0:
        return str(number)
    return result