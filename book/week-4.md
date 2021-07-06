# Chapter 3 - Week 3

```
day0: Dungeon Game
day1: Dungeon Game
day2: Dungeon Game
day3: Dungeon Game
day4: Dungeon Game
day5: Dungeon Game
day6: Dungeon Game
```

## [DAY-0] Dungeon Game

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

## [DAY-1] Dungeon Game

Our game is quite limited, and also there there is no way 

## [DAY-2] Dungeon Game
## [DAY-3] Dungeon Game
## [DAY-4] Dungeon Game
## [DAY-5] Dungeon Game
