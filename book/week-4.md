# Chapter 3 - Week 3

```
day0: Basics of Basics
day1: Dungeon Game
day2: Dungeon Game
day3: Dungeon Game
day4: Dungeon Game
day5: Dungeon Game
day6: Dungeon Game
```

## [DAY-0] Basics of Basics

Make a new file form IDLE, and write in it:
```
print("Hi!")
```

Hit F5 and enjoy!

Your first(or second) python program.. Quite useless, but nevertheless. A program is a program!

`print()` is called a function, functions are kind of mini useful programs, this one will print whatever you tell it. In this case `print("Hi!")` will output `Hi!`. Pretty amazing, I hope one day to explain to you what goes into showing a character on your screen. Remind me to show you Ben Eater's 6502 videos when you grow up.


```
while True:
  print("Hi!")
```

`while True` that means forever. so you will see:

```
Hi!
Hi!
Hi!
Hi!
Hi!
Hi!
Hi!
Hi!
...

```

Hi forever.


```
while 1 == 1:
  print("Hi!)
```

`1 == 1` is also True, so this is the same as `while True`


```
while True:
  name = input("what is your name: ")
  print(name)
```

`input` asks you to type something, and it returns whatever you typed.

`name = input(....)` takes whatever input returns, and puts it into memory that later we can use.

`print(name)` prints whatever is in the memory pointed by `name`.

```

Lets make a more complicated one, this one will ask you what is your name, and if you answer pikachu will print pikaaaaaaachuuuuu and stop, any other answer it will print "Hello, " + answer, so if you type "Jane" it will show "Hello, Jane" and it will ask again.

```
while True:
  name = input("what is your name: ")
  if name == "pikachu":
    print("pikaaaaaaachuuuuu")
    break
  else:
    print("Hello, " + name)
```



WIP WIP WIPW IPWIPWIPWIW

`while,if,else,break` are keywords, kind of like `<html>` and `<body>`, `<h1>` etc, those are coming from python itself.

`if name == "pikachu":` checks if name is equal to `pickachu`, you see how  


checks whatever is in `...` to be `True`, but what does `True` mean? It take a lifetime to find out what is truth, but luckily python is easier than life. for example `1 == 1` is True, `1 == 0` is False (not true), `1 != 0` is True, `!` means not, so `1` is in fact not equal to `0`. The `...` is called `condition` 


For example `1 == 1` is `True`, `1 == 0` is `False`. Same for `if ...`


`while True` means "forever" because `while ...` . , `True` is also `True`. What is `5 == 1`?


If you master those few things you are 75% there, the rest is just searching through sorted and unsorted deck of cards.. anyway, I will show you later.




## [DAY-1] Dungeon Game

```
print("Welcome to the forest!")
print("This is a world of magic and wonders. You are on a cross road, you can go north east south or west.")
print("South leads to the swamps, where the aligators live")
print("West leads to the mountains, where the yeti lives")
print("North leads to the jungle, where the tigres live")
print("East leads to the desert, where the meerkats live")
print("")

what = input("what would you do next: ")

if what == "east":
  print("Welcome to the desert")
  print("It is very hot, and you need remember to put sun screen.")
  print("Oh, you remember to put your hat as well.")
  print("In the distance you see a meerkat approaching.")
  print("What would you do? Run or Fight the meerkat?")
  print("")

  what = input("what would you do next: ")
  if what == "run":
    print("You start running, and the meerkat is super fast, it catches you in no time.")
  elif what == "fight":
    print("You try to fight it, but turns out it was friendly and wants to become your friend.")
    print("What would you do?")
    print("")
    what = input("Do you want to be its friend: ")

    if what == "yes":
      print("Great! Now the meerkat wants to introduce you to his family.")
    elif what == "no":
      print("The meerkat starts crying!")

  else:
    print("I dont understand: " + what)

```

That is a lot of typing.

In python there are few very important things. First even though it doesn't look like it, it is actually a tree, like HTML tree. The parent of `print("The meerkat starts crying!")` is `elif what == "no":`

```
                    |
                    |
              if what == "east"
                   |
                 / |  \
                /  |   \
               /   |    \
              /    |     \
             /     |      \
          if      else    elif
    what == "run"  |     what == "fight
          |        |        |
        print   print      / \
                          /   \
                         /     \
                        /       \
                      if        elif
                what == "yes" what == "no"
                     |            |
                  [ ... ]      [ ... ]
```

Lets discuss another example:

```
a = "he"
b = "ll"
c = "o"
if a == "he":
  if b == "ll":
    if c == "o":
      print("hello")
```

Can you tell who is the parent of who?

BTW, you can also use `and`, `or` and `not` when you want to see if something is `True`

```
a = "he"
b = "ll"
c = "o"
if a == "he" and b == "ll" and c == "o":
  print("hello")
```


This is going waay too fast, and is super confusing, but its ok, we will be doing this for few weeks, so don't worry if it makes no sense now.

## [DAY-2] Dungeon Game

Our game is quite limited, and a quick step to improve it is to make you ask where you want to go until you pick one of the options.

```

def ask(possible_answers):
  answer = ''
  print("---") # print empty line
  while True:
    answer = input("> What would you do next: ")
    if answer not in possible_answers:
      print("> try again, it must be one of:", possible_answers)
    else:
      return answer

print("Welcome to the forest!")
#...

what = ask(["east","west","north","south"])

if what == "east":
  print("Welcome to the desert")
  #...

  what = ask(["run","fight"])
  if what == "run":
    print("You start running, and the meerkat is super fast, it catches you in no time.")
  elif what == "fight":
    print("You try to fight it, but turns out it was friendly and wants to become your friend.")
    # ...
    what = ask(["yes","no"])
    if what == "yes":
      print("Great! Now the meerkat wants to introduce you to his family.")
    elif what == "no":
      print("The meerkat starts crying!")
elif what == "north":
  print("Welcome to the north pole, it is super cold here.")

```

We made a small program called `ask` that takes a list of possible answers, and asks you to enter one of them until what you wrote is one of the possible answers, and when the answer matches it returns it to wherever the program was called. In our case we do `what = ask(["yes","no"])` that means `ask` will keep asking, using the `input` program.

In fact, try making `ask` into a standalone program, make a new file:

```
possible_answers = ["yes","no"]
answer = ''
while True:
  answer = input("> What would you do next: ")
  if answer not in possible_answers:
    print("> try again, it must be one of:", possible_answers)
  else:
    break
```

I cheated a bit changing `return answer` to `break`, but I need to stop the infinite loop, previously it was stopped by `return`, and you cant `return` from the main program (in python).

I user the word `program` as just a piece of code that is run by the processor, there is very little difference between the `main` program, and the `print`, `ask`, `input`, or anything you call with `()`, another name for those is **functions**, so `print` is a function, and `ask`, and `input`. 

Functions usually take some parameters and return some values, for example:

```
def sum(a,b):
  return a + b

r = sum(1737,1231231)
print(r)

```

`def` is a python keyword that tells it, look whatever is after that, before `(`, is the name of the function, and inside `(...)` is information coming into the function, like in `sum`'s case, it is the two numbers we want to sum, and since we dont know the numbers before someone calls `sum(1,2)`, we use those placeholders `a` and `b`, what we know is 
## [DAY-2] Dungeon Game
## [DAY-3] Dungeon Game
## [DAY-4] Dungeon Game
## [DAY-5] Dungeon Game
