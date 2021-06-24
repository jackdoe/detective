
def init():
    game = [
      [' ','a','b','c','d'],
      ['1','-','-','-','-'],
      ['2','-','-','-','-'],
      ['3','-','-','-','-'],
      ['4','-','-','-','-']
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
while True:
    show(g)
    p = input('place ' + x_or_zero + ': ')

    if p == 'a1':
        g[1][1]=x_or_zero
    if p == 'a2':
        g[2][1]=x_or_zero
    if p == 'a3':
        g[3][1]=x_or_zero
    if p == 'a4':
        g[4][1]=x_or_zero

    if p == 'b1':
        g[1][2]=x_or_zero
    if p == 'b2':
        g[2][2]=x_or_zero
    if p == 'b3':
        g[3][2]=x_or_zero
    if p == 'b4':
        g[4][2]=x_or_zero

    if p == 'c1':
        g[1][3]=x_or_zero
    if p == 'c2':
        g[2][3]=x_or_zero
    if p == 'c3':
        g[3][3]=x_or_zero
    if p == 'c4':               
        g[4][3]=x_or_zero

    if p == 'd1':
        g[1][4]=x_or_zero    
    if p == 'd2':
        g[2][4]=x_or_zero
    if p == 'd3':
        g[3][4]=x_or_zero
    if p == 'd4':
        g[4][4]=x_or_zero


    if p == 'clear':
        g = init()

        
    if x_or_zero =='x':
        x_or_zero = '0'
    else:
        x_or_zero='x'
