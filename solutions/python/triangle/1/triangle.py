def validate_sides(sides: list[int]) -> bool:
    '''Validate the sides of a triangle

    Args:
        sides (list): sides of a triangle 
    
    Returns:
        bool: if the sides satisfy the conditions of a triangle 

    Notes:
        A triangle needs to have the length of each side greater than 0
        A triangle needs to have the sum of the lengths of two sides greater than 
        or equal to the length of one side
    '''
    if sides[0] + sides[1] < sides[2] or sides[1] + sides[2] < sides[0] or sides[0] + sides[2] < sides[1]:
        return False
    for side in sides:
        if side <= 0 :
            return False
    return True


def equilateral(sides: list[int]) -> bool:
    '''Determine if a triangle is equilateral

    Args:
        sides (list): sides of a triangle

    Returns:
        bool: if a triangle satisfies the condition of an equilateral triangle
    Notes:
        An equilateral triangle has the length of its sides equal
    '''
    if not validate_sides(sides=sides):
        return False
    if sides[0] == sides[1]:
        if sides[1] == sides[2]:
            return True
    return False


def scalene(sides: list[int]) -> bool:
    '''Determine if a triangle is scalene

    Args:
        sides (list): sides of a triangle

    Returns:
        bool: if a triangle satisfies the condition of a scalene triangle

    Notes:
        A scalene triangle has all sides of different lengths
    '''
    if not validate_sides(sides=sides):
        return False
    if sides[0] == sides[1] or sides[0] == sides[2] or sides[1] == sides[2]:
            return False
    return True

def isosceles(sides: list[int]) -> bool:
    '''Determine if a triangle is isosceles

    Args:
        sides (list): sides of a triangle

    Returns:
        bool: if a triangle satisfies the condition of an isosceles triangle

    Notes:
        An isosceles triangle has at least two of its sides with the same length
    '''
    if not validate_sides(sides=sides):
        return False
    if sides[0] == sides[1] or sides[0] == sides[2] or sides[1] == sides[2]:
        return True
    return False