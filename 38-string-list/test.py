import sys


def base_convert(val, n, chars):
    s = ''
    total_chars = len(chars)

    while val > 0:
        s = chars[val % total_chars] + s
        val /= total_chars

    if len(s) < n:
        s = chars[0] * (n - len(s)) + s

    return s


def get_possibilities(n, chars):
    chars = list(set(chars))

    possibilities = []

    for i in xrange(len(chars) ** n):
        possibilities.append(base_convert(i, n, chars))

    return possibilities


# for i in xrange(2 ** 3):
#     print base(i, 3, ['p', 'o'])

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    line = test.strip()

    total_chars, chars = line.split(',')

    total_chars = int(total_chars)

    print get_possibilities(total_chars, chars)
