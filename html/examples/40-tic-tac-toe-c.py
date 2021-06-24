def showgame(game):
    for row in range(len(game)):
        for col in range(len(game[row])):
            print(game[row][col],end=' ')
        print('')

def ask(what):
    return input("place " + what + ": ")

def initial():
    #return [' ','a','b','c','1','-','-','-','2','-','-','-','3','-','-','-',]
    return [
        [' ','a','b','c'],
        ['1','-','-','-'],
        ['2','-','-','-'],
        ['3','-','-','-'],
    ]

def place(game, row, col, symbol):
    if game[row][col]=='-':
        game[row][col]=symbol
        return True
    return False


board = initial()
symbol='x'
while True:
    showgame(board)
    position = ask(symbol)
    success = False
    if position == 'a1':
        success = place(board,1,1,symbol)
    if position == 'b1':
        success = place(board,1,2,symbol)
    if position == 'c1':
        success = place(board,1,3,symbol)
    if position == 'a2':
        success = place(board,2,1,symbol)
    if position == 'b2':
        success = place(board,2,2,symbol)
    if position == 'c2':
        success = place(board,2,3,symbol)
    if position == 'a3':
        success = place(board,3,1,symbol)
    if position == 'b3':
        success = place(board,3,2,symbol)
    if position == 'c3':
        success = place(board,3,3,symbol)

    if position == 'clear':
        board=initial()
        
    if success:    
        if symbol == 'x':
            symbol='0'
        else:
            symbol='x'
