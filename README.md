# GO-STATUS

## OUTLINE

- [GO-STATUS](#go-status)
  - [OUTLINE](#outline)
  - [The Issue](#the-issue)
  - [Installation](#installation)

## The Issue
While writing a little state machine in Go, it bugged me that I could not find a clean way to implement flag operations on enumerations, without repeating code. It may have had to do with a lack of experience, having started to use Go only a few months ago, but I did find a way that I could consider "clean", using the generics feature, available since Go version 1.18. 

## Installation

```bash
go get github.com/discomco/go-status
```
