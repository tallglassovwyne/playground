""" Implements functions that determine the type of a triangle based on its sides.

Each function takes the sides of a triangle and returns a boolean indicating whether the triangle is of that type.

A triangle must have all its sides greater than zero and the sum of any two sides must be greater than or equal to the third side to be valid.
"""

def is_valid_triangle(sides: list[float]) -> bool:
    """ Validate if the given sides can form a triangle.

    A triangle is valid if all sides are greater than zero 
    and the sum of any two sides is greater than or equal the third side.

    Args:
        sides (list): sides of a triangle
    Returns:
        bool: True if the triangle is valid, False otherwise.
    """
    is_sides_greater_than_zero = all(side > 0 for side in sides)
    return is_sides_greater_than_zero and (sides[0] + sides[1] >= sides[2] and sides[0] + sides[2] >= sides[1] and sides[1] + sides[2] >= sides[0])

def equilateral(sides: list[float]) -> bool:
    """ Determine if a triangle is equilateral.

    An equilateral triangle has its three sides the same length.

    Args:
        sides (list[float]): sides of a triangle
    Returns:
        bool: True if the triangle is equilateral, False otherwise.
    """
    return is_valid_triangle(sides) and all(side == sides[0] for side in sides)

def isosceles(sides: list[float]) -> bool:
    """ Determine if a triangle is isosceles.
    
    An isosceles triangle has at least two sides the same length.
    
    Args:
        sides (list[float]): sides of a triangle
    Returns:
        bool: True if the triangle is isosceles, False otherwise.
    """
    return is_valid_triangle(sides) and (sides[0] == sides[1] or sides[1] == sides[2] or sides[0] == sides[2])

def scalene(sides: list[float]) -> bool:
    """ Determine if a triangle is scalene.

    A scalene triangle has all three sides of different lengths.

    Args:
        sides (list[float]): sides of a triangle
    Returns:
        bool: True if the triangle is scalene, False otherwise.
    """
    return is_valid_triangle(sides) and sides[0] != sides[1] and sides[0]!= sides[2] and sides[1] != sides[2]