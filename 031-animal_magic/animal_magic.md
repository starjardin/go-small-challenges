## Task 1: Role a die

Implement a RollADie function.

This will be the traditional twenty-sided die with numbers 1 to 20.

```javascript
d := RollADie() // d will be assigned a random int, 1 <= d <= 20
```

## Task 2: Generate a wand energy

Implement a GenerateWandEnergy function. The wand energy should be a random floating point number between 0.0 and 12.0.

```javascript
f := GenerateWandEnergy()  // f will be assigned a random float64, 0.0 <= f < 12.0
```

## Task 3: Shuffle a slice

The game features eight different animals:

- ant
- beaver
- cat
- dog
- elephant
- fox
- giraffe
- hedgehog

Write a function ShuffleAnimals that returns a slice with the eight animals in random order.