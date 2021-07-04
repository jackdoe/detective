# Chapter 0 - Week 0

```
day0: touch typing on keybr.com
day1: install python and make the program
day2: touch typing using your program
day3: html
day4: touch typing
day5: html
day6: html
```

This week is one of the more difficult ones. Lots of new things, and some of them might not work on your computer because it is too new or too old, for which I apologize. You might need a bit of help. If you get stuck just call a friend or parent to help you out.

## The Computer

All modern computers(laptops, phones, pc master race rgb monsters, etc) have somewhat similar components: Processor, Memory, Video Card, Disk and USB controller, WiFi card etc. Some of them are in one single chip and you cant even see them anymore, but they are there. For example there are chips that have Processor and Video Card together. The term for processor is actually CPU - Central processing unit, but we called it processors when we were kids and it kind of make sense, since it processes stuff.

```
+------------------+
|                  |
|  [           ]   |
|  [ Processor ]   |
|  [           ]   |
|                  |
|  [ Memory ]      |
|  [ Memory ]      |
|                  |
|  [ Video Card ]--+------> monitor
|                  |
|  [ Disk ]        |
|                  |  +---> iphone
|                  |  +---> mouse
|  [ USB ]---------+--+---> keyboard
|                  |
|  [ WiFi ]        |
|  [ ... ]         |
+------------------+

```

Memory (also called RAM), is the place where programs exist to be run by the CPU, it loads its instructions from there, and the instructions tells it what to do, it can write back to memory, or read from a different place or write something to the disk controller or to the video card, etc. Basically the most important stuff happens between the processor and the RAM (which stands for Random Access Memory). It is Random Access because the processor can just go and read or write to specific place, which is pretty cool. Everything in the RAM disappears when you turn off the computer, so usually the programs and all kinds of information is persistently stored on the Disk. Which is called disk because it used to be spinning, but now most computers have SSD disks which are not *disks* at all, google 'SSD versus spinning disk' if you want to find more. So when the computer starts it will load some information from the disk into memory and start the operating system, which is just a program, not very different than the one you wrote for the touch typing, it takes input, does something with it and then has some output.


### Files

When you store something to disk you usually store it as a file, files can be all kinds, only text, or images, or just raw binary data used for different purposes. Text files are literally that, file that contains text, later you will learn that there is no such thing as text to the computer, all is just bits of information, but we (the humans) decided it is quite useless to look at some numbers of information, so our programs display information in different way, for example the file can contain the number 97 in its raw form 01100001, and if seen from a text editor, it will show the letter 'a', but if seem from image editor it might use it to show some blue-ish color, colors are stored in 4 numbers: transparency(called alpha), red,green and blue, but different formats use different way to store the colors, so it will depend on which format does the program think it is going to display.

See it is all in the eye of the beholder, the file still contains only 97 (01100001), but each program will decide what to do with it.

There are also Directories(also called Folders) (which in many file systems are also files.. but we will get to that later), that can contain files or other directories. Example structure looks like:

```

/
├─ home/
│  ├─ jack/
│  |  ├─ type.py
│  |  ├─ example.txt
│  |  ├─ image.jpg

```

The word File is used because people that made those things up were thinking of filing cabinet, with lots of folders that contains pieces of paper.

```
  /
  /
 /____________________
 |________  __________
 /_____  /||   |
|".___."| ||   |
|_______|/ |   |
 || .___."||  /
 ||_______|| /
 |_________|/ Felix Lee
 
 ```

> See how those pictures have sometimes names attached to them. Those are their creators, and those pictures are pretty old, some are from the 80s, so it is amazing we can still find out who made them

So the name 'File' and 'Folder' is used. If people were thinking of libraries I guess we would've been using Book for a file and Shelve for Folder? Anyway, my point is: all things have names, some of the names are strange and old, so don't worry if they make no sense.

<div style="page-break-before:always"></div>

## Touch Typing

Touch typing is just typing without looking, if you are super slow while typing you will get frustrated too soon.

Use https://www.keybr.com/ or ask someone to recommend you something. After using the touch typing app for few days, write your own.

Ask someone to help you to install python, or google 'how to install python', 

After you have installed python3 run IDLE (ask your parent for help) and go to File -> New File and type:

```
import random
letters = 'fjfjfjfjfjfjjfghghghgaa;a;a;abcdefghijklmnopqrstuvwxyz'
difficulty = 2

while True:
    q = ''
    for i in range(difficulty):
        q = q + random.choice(letters)

    a = input(q + ': ')
    if a == q:
        difficulty += 1
    else:
        if difficulty > 2:
            difficulty -= 1
```

Then hit F5 (this will ask you to save the file as well) to run the program. If it is too easy, try adjusting the difficulty or the letters. It will look a bit like:

```
aj: aj
akg: akg
jjgh: jjgh
```

If you type it correctly, it will ask for a bigger sequence of characters, and if you make mistake it will ask for smaller. 

Don't rush it.

Place your fingers properly on the keyboard, open images.google.com and search for 'touch typing' you will see many images of how to place your fingers.

<div style="page-break-before:always"></div>

## Magic

```
 ______________     _             _,-----------._        ___
|              |   (_,.-      _,-'_,-----------._`-._    _)_)
| THE _  _  _  |      |     ,'_,-'  ___________  `-._`.
| |  / \|_)| \ |     `'   ,','  _,-'___________`-._  `.`.
| |__\_/| \|_/ |        ,','  ,'_,-'     .     `-._`.  `.`.
|              |       /,'  ,','        >|<        `.`.  `.\
| OF THE  _ _  |      //  ,','      ><  ,^.  ><      `.`.  \\
| |_)||\|/_(_  |     //  /,'      ><   / | \   ><      `.\  \\
| | \|| |\_|_) |    //  //      ><    \/\^/\/    ><      \\  \\
|______________|   ;;  ;;              `---'              ::  ::
                   ||  ||              (____              ||  ||
 DOORS OF DURIN   _||__||_            ,'----.            _||__||_
                 (o.____.o)____        `---'        ____(o.____.o)
                   |    | /,--.)                   (,--.\ |    |
                   |    |((  -`___               ___`   ))|    |
                   |    | \\,'',  `.           .'  .``.// |    |
                   |    |  // (___,'.         .'.___) \\  |    |
                  /|    | ;;))  ____) .     . (____  ((\\ |    |\
                  \|.__ | ||/ .'.--.\/       `/,--.`. \;: | __,|;
                   |`-,`;.| :/ /,'  `)-'   `-('  `.\ \: |.;',-'|
                   |   `..  ' / \__.'         `.__/ \ `  ,.'   |
                   |    |,\  /,                     ,\  /,|    |
                   |    ||: : )          .          ( : :||    |
                  /|    |:; |/  .      ./|\,      ,  \| :;|    |\
                  \|.__ |/  :  ,/-    <--:-->    ,\.  ;  \| __,|;
                   |`-.``:   `'/-.     '\|/`     ,-\`;   ;'',-'|
                   |   `..   ,' `'       '       `  `.   ,.'   |
                   |    ||  :                         :  ||    |
                   |    ||  |                         |  ||    |
                   |    ||  |                         |  ||    |
                   |    |'  |            _            |  `|    |
                   |    |   |          '|))           |   |    |
                   ;____:   `._        `'           _,'   ;____:
                  {______}     \___________________/     {______}
              SSt |______|_______________________________|______|
```

Have you seen The Lord of the Rings? Do you remember the Door of Durin? Where Gandalf couldn't remember how to enter and the fellowship almost got eaten?

The experience through this book will feel like that. Sometimes you will be stuck, and I mean really really stuck, and at some point, as if the cosmos itself colludes, something will click, and you will level up. Sadly I can not know when this will happen to **you**, as every person is different.


### HTML

We will start with HTML.


Making a page can be a bit overwhelming, so just start small. for example:

Open notepad and make a file on your desktop named "first.html", then write in it:

```
<html>
 <body>
   <h1>welcome!</h1>
   <p>this is my page</p>
 </body>
</html>
```




### ASCII

Computers do not understand text, so when you see text it is usually some number that is desplayed as a character, for example 65 is A, 66 is B, and so on. About 60 years ago some people agreed on a common way to map number to character, called THE ASCII STANDARD, super fancy name, for such a simple thing.

```
32  SPC   64  @    96  `
33  !     65  A    97  a
34  "     66  B    98  b
35  #     67  C    99  c
36  $     68  D   100  d
37  %     69  E   101  e
38  &     70  F   102  f
39  '     71  G   103  g
40  (     72  H   104  h
41  )     73  I   105  i
42  *     74  J   106  j
43  +     75  K   107  k
44  ,     76  L   108  l
45  -     77  M   109  m
46  .     78  N   110  n
47  /     79  O   111  o
48  0     80  P   112  p
49  1     81  Q   113  q
50  2     82  R   114  r
51  3     83  S   115  s
52  4     84  T   116  t
53  5     85  U   117  u
54  6     86  V   118  v
55  7     87  W   119  w
56  8     88  X   120  x
57  9     89  Y   121  y
58  :     90  Z   122  z
59  ;     91  [   123  {
60  <     92  \   124  |
61  =     93  ]   125  }
62  >     94  ^   126  ~
63  ?     95  _

```

Try to decode: 110 105 99 101 32 119 111 114 107 33
