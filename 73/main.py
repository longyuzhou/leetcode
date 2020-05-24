def setZeroes(matrix):
    stack = []
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            if matrix[i][j] == 0:
                stack.append((i, j))

    while len(stack) > 0:
        y, x = stack.pop()
        for i in range(len(matrix[y])):
            matrix[y][i] = 0
        for i in range(len(matrix)):
            matrix[i][x] = 0


def printMatrix(matrix):
    for i in range(len(matrix)):
        print(matrix[i])


matrix = [
    [1, 1, 1],
    [1, 0, 1],
    [1, 1, 1]
]
setZeroes(matrix)
printMatrix(matrix)

matrix = [
    [0, 1, 2, 0],
    [3, 4, 5, 2],
    [1, 3, 1, 5]
]
setZeroes(matrix)
printMatrix(matrix)
