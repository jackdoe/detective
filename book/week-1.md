# Chapter 1 - Week 1

```
day0: make new touch typing program
day1: touch typing on keybr.com
day2: touch typing using your program
day3: HTML
day4: HTML
day5: touch typing
day6: HTML
```

## [DAY-0] New Touch Typing Program

Make new file on your desktop, touchtyping.html and type in it:

```
<center>
  <pre style="font-size: 30px; letter-spacing: 4px;" id="question"></pre>
</center>

<script>
var makeNewQuestion = function() {
  var letters = 'fjfjfjfjfjfjjfghghghgaa;a;a;abcdefghijklmnopqrstuvwxyz'
  var difficulty = 10
  var q = []
  for (var i = 0; i < difficulty; i++) {
    q.push(letters.charAt(Math.floor(Math.random() * letters.length)));
  }
  return q
}

var currentQuestion = makeNewQuestion()
var questionElement = document.getElementById("question")
questionElement.innerHTML = currentQuestion.join("")

var position = 0
document.body.addEventListener('keydown', (e) => {
  if (e.key == currentQuestion[position]) {
    position++
    if (position == currentQuestion.length) {
      currentQuestion = makeNewQuestion()
      position = 0
    } else {
      currentQuestion[position-1] = '<b>' + currentQuestion[position-1] + '</b>'
    }
    questionElement.innerHTML = currentQuestion.join("")
  }
})
</script>
```

Don't worry about what it all means. Just type it character by character. Now double click on the file and try out your program, just start typing as the web page is open, and characters will become bold as you type them.

