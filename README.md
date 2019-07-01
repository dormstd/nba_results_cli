# nba_results_cli
This is a really simple program developed to learn a bit of go while doing something interesting.

It parses the information from https://stats.nba.com to show in the command line the results of the matches for a given date.
## Compilation
To compile it simply run:

    go build

## Usage
There is only one parameter that can be passed to the program:
    --date MM/DD/YYYY

To run it simply run the binary:

    ./nba_results_cli


Or if you want the results for a given date:

    ./nba_results_cli --date 03/06/2015
