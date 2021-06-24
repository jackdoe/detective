def line(width, symbol='x'):
    for x in range(width):
        print(symbol,end='')

def box(width,height,symbol='x'):
    for y in range(height):
        line(width,symbol)
        print('')

box(16,20)
box(10,20,'#')
box(15,13,'*')