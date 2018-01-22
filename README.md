# softwareninjas
## Problem 1 (Unproductive Time Calculator)
You must determine how long I have been away from my desk for non-work related activities. You will be given lines of data with a start time and stop time for each activity. The rub is that I might overlap. For example:

```
1 3 \\ Leaving for walk
2 3 \\ Puzzling
4 5 \\ Bathroom
```


The output for the above list should be 3 hours. Here is the test data, all hours are in 24 hour notation so no AM/PM conversion is necessary. All inputs will be integers:

```
2 4
3 6
1 3
6 8
```
Output: `7`

```
6 8
5 8
8 9
5 7
4 7
```
Output: `5`


Bonus Data:
```
15 18
13 16
9 12
3 4
17 20
9 11
17 18
4 5
5 6
4 5
5 6
13 16
2 3
15 17
13 14
```