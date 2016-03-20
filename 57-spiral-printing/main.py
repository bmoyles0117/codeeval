import sys


def get_inner_box(box):
    inner_box = []

    height = len(box)
    width = len(box[0])

    for y in xrange(height - 2):
        inner_box.append(box[y + 1][1:width - 1])

    return inner_box


def get_spiral_values(box):
    height = len(box)
    if height <= 0:
        return []

    width = len(box[0])
    if width <= 0:
        return []

    values = box[0][:]
    for i in xrange(1, height):
        values.append(box[i][width - 1])

    if height - 1 > 0:
        values.extend(box[height - 1][:-1][::-1])

    if width > 1:
        for i in xrange(height - 2, 0, -1):
            values.append(box[i][0])

    values.extend(get_spiral_values(get_inner_box(box)))

    return values


test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    line = test.strip()
    height, width, values = line.split(';')

    height = int(height)
    width = int(width)
    values = values.split(' ')

    box = []
    for y in xrange(height):
        row = []
        for x in xrange(width):
            row.append(values[y * width + x])
        box.append(row)

    print ' '.join(get_spiral_values(box))
