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

#### Chapter 1: Arrays and Strings (5/9)

|  #  | Problem                        | Tests         | Solved |
|-----|--------------------------------|:-------------:|:------:|
| 1.1 | [Is Unique][p1.1]              | [tests][t1.1] |   ✅   |
| 1.2 | [Check Permutation][p1.2]      | [tests][t1.2] |   ✅   |
| 1.3 | [URLify][p1.3]                 | [tests][t1.3] |   ✅   |
| 1.4 | [Palindrome Permutation][p1.4] | [tesst][t1.4] |   ✅   |
| 1.5 | [One Away][p1.5]               | [tests][t1.5] |   ✅   |

<!-- References -->
[p1.1]: interview_questions/1_arrays_and_strings/1.1_is_unique/main.go
[t1.1]: interview_questions/1_arrays_and_strings/1.1_is_unique/main_test.go
[p1.2]: interview_questions/1_arrays_and_strings/1.2_check_permutation/main.go
[t1.2]: interview_questions/1_arrays_and_strings/1.2_check_permutation/main_test.go
[p1.3]: interview_questions/1_arrays_and_strings/1.3_urlify/main.go
[t1.3]: interview_questions/1_arrays_and_strings/1.3_urlify/main_test.go
[p1.4]: interview_questions/1_arrays_and_strings/1.4_palindrome_permutation/main.go
[t1.4]: interview_questions/1_arrays_and_strings/1.4_palindrome_permutation/main_test.go
[p1.5]: interview_questions/1_arrays_and_strings/1.5_one_away/main.go
[t1.5]: interview_questions/1_arrays_and_strings/1.5_one_away/main_test.go
