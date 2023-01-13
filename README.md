# 99 Bottles exercise

This exercise is adapted from the appendix of "99 Bottles of OOP" by Sandi Metz and Katrina Owen.

Download the code for this exercise with
```sh
git@github.com:zeus-health/shuang-scratch.git -b 99bottles
cd 99bottles/bottles
go mod download
```

## Doing the exercise

To run the test suite, run the following from the `bottles` folder.
```sh
go test ./...
```

The test suite contains one failing test, and many commented-out tests. Your goal is to write code that passes all of the tests. Follow this protocol:

- run the tests and examine the failure
- write only enough code to pass the failing test
- uncomment the next test (this simulates writing it yourself)

Repeat the above until no test is skipped, and you’ve written code to pass each one.

Work on this task for 30 minutes. The vast majority of folks do not finish in 30 minutes, but it’s useful, for later comparison purposes, to record how far you got.  Even if you can’t force yourself to stop at that point, take a break at 30 minutes and save your code.
