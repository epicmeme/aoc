# Advent of Code 🎄

This is where I store my solutions for Advent of Code. It is a mess and was made in like half an hour or so, but it kinda works.

## Codegen
Use the generators to create the boilerplate for a new day.

```
go run main.go <year> <day>
```

It will scaffold a new directory with the following files:

- `main.go` - The main entry point for the day's solution.
- `main_test.go` - A benchmarking file for the day's solution.
- `input.txt` - The input for the day.

## Notes

- The input is fetched from the AoC website, so you need to have a valid session cookie in the `config.yml` file.
- This project is very poorly constructed, but it works for me. I am still a pleb when it comes to Go.


## TODO

- Clean up the mess that is this repo. I am not a good Go developer, so the structure and naming is probably weird and not very idiomatic.
- Add a script to push the solutions to the AoC website.
- Resolve the input fetching, it is confusing since we fetch in some places, and read the locally stored input in other places.
