# Alien Invasion Simulator

A deterministic CLI application to simulate alien invasion as described in the [task](./resources/task.md)

# Assumptions

1. City names don't contain space and =
2. Path between cities is bidirectional
3. Each city can be connected at most to 4 other cities 
4. Only one city can be present in one direction
5. While initializing ts possible to land multiple aliens in one city. These aliens fight in special "Iteration 0" killing each other and destroying the city
6. If more than one aliens are present in the same city then all of them die while fighting and destroying the city
7. Aliens don't fight in the path between cities and can move to cross each other. i.e Alien 0 moves from city "A" to city "B" and at the same time Alien 1 moves from city "B" to city "A"
8. Program also terminates when all the aliens are trapped
9. If an alien is trapped in a city then at the end of the simulation city still exists in the World

# Implementation

## Steps 

1. Read arguments from the CLI or use default values
2. Initialize the world using a provided input file. The world is modeled as a bidirectional graph with the city as nodes and the path between them as edges.
3. Place aliens randomly in the cities of the world using a pseudo-random generator provided by golang's math/rand package
4. Aliens fight after they are placed, killing each other and destroying the city. Modeled this as "Iteration 0"
5. Simulation termination conditions: If any of the condition is met simulation is terminated
    1. All Aliens are dead
    2. All Aliens are trapped
    3. Max iterations are reached
6. Simulation steps:
    1. Move: Aliens move randomly in the neighbouring city if a path exists
    2. Fight: If more than two aliens are present in the same city they fight and kill each other destroying city
7. Print all trapped aliens(if any)
8. Print post-invasion World  

# Usage

## Requirements
- [Latest Golang version](https://go.dev/doc/install)

## Test
```go test -v -cover ./...```

## Build
```go build -o simulator```

## Run

### Help
```./simulator -h```

```
Usage of ./simulator:
  -inputFile string
        World map input file
  -maxIterations int
        Maximum number of iterations (default 10000)
  -numberOfAliens int
        Number of aliens invading (default 5)
  -seedValue int
        Seed value for deterministic state (default 42)
```

### Example

Sample world map with default arguments

```./simulator -inputFile ./test/testData/sample.txt```

```
Alien 0 placed in city Bar
Alien 1 placed in city Bee
Alien 2 placed in city Foo
Alien 3 placed in city Bar
Alien 4 placed in city Foo
Iteration 0
Bar has been destroyed by aliens 0, 3!
Foo has been destroyed by aliens 2, 4!
Iteration 1
Alien 1 trapped in city: Bee

Baz
Bee
Qu-ux
```

Sample world map with 10 Aliens

```/simulator -inputFile ./test/testData/sample.txt -numberOfAliens 10```

```
Alien 0 placed in city Bar
Alien 1 placed in city Bee
Alien 2 placed in city Foo
Alien 3 placed in city Bar
Alien 4 placed in city Foo
Alien 5 placed in city Bar
Alien 6 placed in city Bee
Alien 7 placed in city Baz
Alien 8 placed in city Foo
Alien 9 placed in city Foo
Iteration 0
Bar has been destroyed by aliens 0, 3, 5!
Bee has been destroyed by aliens 1, 6!
Foo has been destroyed by aliens 2, 4, 8, 9!
Iteration 1
Alien 7 trapped in city: Baz

Baz
Qu-ux
```

# Improvements
- [ ] Add mocks for unit tests
- [ ] Separate unit tests and integrations tests
- [ ] Add github action CI/CD workflows