# Experimental Go gist

The purpose of this was to quickly try out some theory for a larger project and see how proficient I am on the command line.

See the repo by [clicking here](https://github.com/MizouziE/go-input-gist/).

## Hypothesis

1. Create a slice of options
1. Shown the list of options, read user input
1. Initialise a string with two `%v` placeholders
1. Create a second list of options
1. Shown the second list of options, read user input
1. Take two user-chosen options and use them in a `Printf()` output that incorporates string with placeholders initialised **before** reading second input

The `Printf()` function should be able to act as normal

## Conclusion

It did act as normal, so the value placeholders do not become invalid through parsing/storing in memory.

The problems came when trying to select options from the slices 😅

## Try it for yourslef

### Assumptions

- You have Go installed
- You're ok with a super simple CLI

### Instructions

```sh
go run main.go
```

Run that from within the directory and it will start the application and prompt you for an input.

> **Note**
> This has absolutely no safeguarding so will panic if the input isn't as expected i.e not a number from the lists of options

---

This was just for fun, and I wanted to see if I could throw this all together super quick from the terminal, which I have done. Check it:

![screenshot of terminal commands](/go-gist.png)

