
def init():
    game = [
      [' ','a','b','c','d','e'],
      ['1','-','-','-','-','-'],
      ['2','-','-','-','-','-'],
      ['3','-','-','-','-','-'],
      ['4','-','-','-','-','-'],
      ['5','-','-','-','-','-']
    ]
    return game

def show(game):
    for row in game:
        #[' ','a','b','c','d']
        #...
        for col in row:
            #' '
            #'a'
            #'b'
            print (col, end=' ')
        print('')


g = init()
x_or_zero='x'

letter_to_col = {
    'a': 1,
    'b': 2,
    'c': 3,
    'd': 4,
    'e': 5,
}

while True:
    show(g)

    p = input('place ' + x_or_zero + ': ')

    # a1
    # a2
    # ...
    if len(p) == 2:
        # if you type 'a3', then
        # p[0] is 'a'
        # p[1] is '3'
        col = letter_to_col[p[0]]
        row = int(p[1])
        g[row][col] = x_or_zero

    if p == 'clear':
        g = init()

    if x_or_zero =='x':
        x_or_zero = '0'
    else:
        x_or_zero='x'
