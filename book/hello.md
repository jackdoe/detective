# Hello World

::: warning
This chapter is for parents, kids skip to the next one.
:::

This book is for parents who know how to code and for kids who don't, but especially for parents and kids who can spend 30 minutes per day, _every day_. I am writing this book as I am teaching my daughter (10), and you know how in some cooking shows, they skip the part where the food is cooking? I wont do that. The book will be longer than it should.

Again, if your children are older, or younger, this book might not work, you could of course still find pieces that work for you.

This is more of a log of my experience so far.

What you need:

* Computer
* Patience
* Internet
* Patience

If you don't have a you can buy raspberry pi 400 for 70$ or so, or something similar that you attach to your TV. If you don't have patience, buy some chamomile tea. You don't need internet subscription, but you would need a bit of internet to download python and do few google searches.

> NB: In the book I refer to 'notepad' or 'double click' assuming most people run windows, if you don't, you will have to translate what the instructions mean.

The schedule is roughly as follows:

* Week 0
  * Touch Typing
  * HTML
* Week 1
  * HTML: h1, marquee, pink text
  * Touch Typing
* Week 2
  * HTML: tables
  * HTML: lists
  * Touch Typing
* Week 3
  * HTML: tables lists
  * HTML: images, licenses
  * JavaScript: few small programs
* Week 4
  * python: code together a text dungeon game
* Week 5
  * python: super basic python (print and input and while True)
* Week 6
  * python: make hangman game
* Week 7
  * python: turtle
* Week 8
  * HTML: love match
  * python: love match
  * python: text trivia game
* Week 9
  * python: super basic def and for
* Week 10
  * python: turtle
* Week 11
  * python: make text tic tac toe game
* Week 12
  * python: pygame
  * python: turtle
* Week 13
  * python: pygame
  * python: turtle
* Week 14
  * python: basics
  * python: turtle
* Week 15
  * HTML: make a table
  * python: make tic tac toe again
  * python: basics
  * python: pygame
* Week 16
  * python: love match
  * python: trivia game
  * python: turtle
* Week 17
  * python: dungeon game
  * python: turtle grid
  * python: turtle keyboard input
* Week 18
  * python: turtle tic tac toe


In most of the weeks you also go back, waaay back, every day you re-iterate variables and for loops, print the numbers from one to 10 forever, ask how many times to be printed, etc.

The reason for so much emphasis on HTML is because it helps with understanding hierarchy, `tr` is child of `table`, `table` is child of `body` body is child of `html`, `td` is a sister to `td` and both are children of `tr` etc, it helped a lot with my daughter understanding how the `else` is sibling to the `if` and how both are children to `while`.

Also it is very easy to debug, and inspect, and get immediate feedback. There are many 'hackers' now on tictok or youtube that shows you how to get a lot robux(money) in roblox by inspecting the page and modifying the HTML, so I think HTML is very important to be understood, not only because it teaches hierarchy, but also it is the canvas of the web.


## Other materials

Play other games as well, https://tomorrowcorporation.com/ has some brilliant games, Human Resource Machine is great way to learn loops and conditional jumps, and 7 Billion Humans is amazing for high level concepts, including recursion and sorting.

The Robot Turtles game is amazing as well, you can find it here: https://www.thinkfun.com/products/robot-turtles/

Scratch works for some kids, mine didn't enjoy it much.

Buy few Arduino NANOs (cheap clones from amazon work as well, but you need to install ch340 driver), and some servo motors and write few super basic programs that turn the servo slowly in one direction or another. Connect `black/brown wire to gnd, red wire to 5v and orange wire to D9`, and run:

```
#include <Servo.h>
#define SERVO_PIN 9
Servo s;
void setup()  {
  s.attach(SERVO_PIN, 1000, 2000);
}

void loop()  {
  s.write(0);
  delay(500);

  s.write(60);
  delay(500);

  s.write(120);
  delay(500);
}
```


## Motivation

Sometimes there is very little motivation. Kids are super tired from school and playdates and extracurricular activities, makes it hard to spend 30 minutes per day on something. Netflix and YouTube and etc are so much more interesting than what I have to offer, it is a rigged game.. Sad that I have to compete with billion parameter models, but here we are.

Anything worth doing, is worth doing poorly. If you cant spend 30 minutes, spend 15, spend 5 minutes if you have to, or even 1. It is all worth it.

Try to spend time in the morning, at least on the weekends, and first thing after school on school days.

We tried to setup time before or after dinner, and it is very very difficult.

<div style="page-break-before:always"></div>
