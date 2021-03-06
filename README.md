# Driving Evolutionary Algorithm

Started with a Go port of a JS box2D top-down car demo (http://domasx2.github.io/gamejs-box2d-car-example/), using the Pixel game library and ByteArena's Go box2D port. Ended up with a car evolving to drive around obstacles. [Video of evolved solutions can be found here!](https://youtu.be/F2pwEOJhNMs)

# Controllable Car Demo

```
cd cmd/controlled
go build && ./controlled
```

- W to accelerate
- S to break
- A/D to steer
- Q to toggle gear between forwards/backwards
- R to reset car

# Evolutionary Car

The car evolves to drive to and gracefully brake on a green target point. It uses roulette wheel selection combined with random cross overs, swaps and mutations to produce the following generation. The fitness function is a combination of the distance from the target and the final velocity at the end of the sequence execution.  

```
cd cmd/evolutionary
go build && ./evolutionary -h
```

Example population after 64 evolutionary iterations:
```
./evolutionary -seed=0

=================== 64 ===================
[0 1 3 2 0 3 3 2 1 0 3 0 4 4 2 3 4 4 4 4] -> 46.30 (d=22.72, v=23.59)
[0 1 1 3 0 3 3 2 1 0 1 0 4 4 3 3 4 4 4 4] -> 46.40 (d=23.13, v=23.26)
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 0 4] -> 64.37 (d=17.35, v=47.02)
[0 1 3 3 4 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 99.60 (d=83.94, v=15.66)
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 51.93 (d=23.52, v=28.41)
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 50.28 (d=23.07, v=27.21)
[0 1 3 3 0 3 3 2 1 0 1 0 0 4 2 3 4 4 4 4] -> 91.33 (d=48.41, v=42.92)
[3 1 3 3 4 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 100.00 (d=84.33, v=15.68)
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 44.90 (d=18.33, v=26.58)
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 48.70 (d=19.39, v=29.31)
```

You can also replay a sequence:
```
Strong example for level 1:
./evolutionary -seed=0 -seq="0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4"
[0 1 3 3 0 3 3 2 1 0 1 0 4 4 2 3 4 4 4 4] -> 44.73 (d=18.25, v=26.48)
```
```
Strong examples for level 2:
./evolutionary -seed=0 -level=2 -seq="0 2 0 3 3 0 4 1 1 0 4 0 3 3 3 3 2 3 0 1 1 2 0 2 0 4 3 4 0 4"
[0 2 0 3 3 0 4 1 1 0 4 0 3 3 3 3 2 3 0 1 1 2 0 2 0 4 3 4 0 4] -> 119.33 (d=16.84, v=102.49)
// conventional solution
[4 3 3 1 4 2 0 3 3 0 0 3 3 3 2 0 4 2 3 0 0 1 4 2 4 2 4 4 4 4] -> 52.00 (d=21.11, v=30.89)
// over the top
[3 4 2 3 3 1 2 1 1 1 3 1 1 3 3 3 1 2 1 3 3 1 1 3 1 4 4 4 4 4] -> 71.55 (d=4.15, v=67.40)
```
