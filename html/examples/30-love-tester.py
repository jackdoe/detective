def love_test(a,b):
    sum = 0
    for c in a+b:
        print(c + ': ' + str(ord(c)))
        sum += ord(c)

    print('sum is:' + str(sum))
    return sum % 100


while True:
    nameA = input("first name: ")
    nameB = input("second name: ")

    print("love match: " + str(love_test(nameA,nameB)))
