import sys

digit_mapping = {
    0: [0],
    1: [1],
    2: [2, 4, 8, 6],
    3: [3, 9, 7, 1],
    4: [4, 6],
    5: [5],
    6: [6],
    7: [7, 9, 3, 1],
    8: [8, 4, 2, 6],
    9: [9, 1]
}


def build_endings(base, power):
    last_digit = int(base[-1])
    power = int(power)

    total_cycles = power / len(digit_mapping[last_digit])
    total_remainder = power % len(digit_mapping[last_digit])

    result = {}
    for n, digit in enumerate(digit_mapping[last_digit]):
        result[digit] = total_cycles
        if n < total_remainder:
            result[digit] += 1

    string_result = []
    for i in range(10):
        string_result.append('%s: %s' % (i, result.get(i, 0), ))

    return ', '.join(string_result)


test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    line = test.strip()

    base, power = line.split(' ')

    print build_endings(base, power)
