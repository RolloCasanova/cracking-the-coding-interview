# Cracking the coding interview

A collection of solutions for the book "Cracking the coding interview" by Gayle Laakmann McDowell (6th edition) made by me in Go just for fun.

## How to run

```bash
go run ./interview_questions/<chapter>/<exercise>/main.go [optional-params]
```

e.g.

```bash
go run interview_questions/1_arrays_and_strings/1.1_is_unique/main.go abcde
```

## How to test

### A single exercise

```bash
go test -v ./interview_questions/<chapter>/<exercise>/...
```

e.g.

```bash
go test -v ./interview_questions/1_arrays_and_strings/1.1_is_unique/...
```

### All exercises of a chapter

``` bash
go test -v ./interview_questions/<chapter>/...
```

e.g.

``` bash
go test -v ./interview_questions/1_arrays_and_strings/...
```

### All exercises

```bash
go test -v ./interview_questions/...
```

### Progress

- Chapter 1: Arrays and Strings (2/9)
  - [✅] [1.1 Is Unique](interview_questions/1_arrays_and_strings/1.1_is_unique)
  - [✅] [1.2 Check Permutation](interview_questions/1_arrays_and_strings/1.2_check_permutation)
