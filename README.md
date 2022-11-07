# Project
Push-swap

# Authors
* Julia
* Karljohannes
* Hgwtra

# Technologies

Project is built with Golang version 1.18

# Description

Push-Swap is a project that uses a Non-Comparative Sorting Algorithm to create 2 programs:
* push-swap: which calculates and displays on the standard output the smallest program using push-swap instruction language that sorts integer arguments received.
* checker: which takes integer arguments and reads instructions on the standard output. Once read, checker executes them and displays OK if integers are sorted. Otherwise, it will display KO.

# Set Up
1. Clone the repo in VSCode or your text editor of choice.
2. Build the program
```console
sh build.sh
```
3. Then use (or follow the audit questions)
```console
./push-swap "2 1 3 6 5 8"
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
```

# List of commands

* pa push the top first element of stack b to stack a
* pb push the top first element of stack a to stack b
* sa swap first 2 elements of stack a
* sb swap first 2 elements of stack b
* ss execute sa and sb
* ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
* rb rotate stack b
* rr execute ra and rb
* rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
* rrb reverse rotate b
* rrr execute rra and rrb

