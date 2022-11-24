# si

[![GoDoc](https://godoc.org/github.com/quartercastle/si?status.svg)](https://pkg.go.dev/github.com/quartercastle/si?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/quartercastle/si)](https://goreportcard.com/report/github.com/quartercastle/si)

> NOTE: this is still a work in progress

si is a unit annotation and conversion library. Its different from other strongly
typed si/unit libraries in the way that all types are aliases to `float64`
instead of concrete types.
The benefit of this approach is that casting can be avoided which results in
more readable code that somewhat resembles the math equation.

To prove my point i have created an example using Einstein's mass-energy equivalence formular $E=mc^2$.

```go
// possible using aliased si types
24.965421631578266*MWh == 1*mg * math.Pow(c, 2)

// problem with conrete si types...
24.965421631578266*MWh == MWh(float64(1*mg) * math.Pow(c, 2))
```

As it can be seen in the example above casting ends up hurting readability and
it really doesn't add anything in terms of type safety because casting will be so
prevalent. It will only be as safe as the aliased approach used in this library.

Another benefit is that derived si types can easily be defined on the fly, this
is not possible when using concrete types unless casting is involved.

```go
const SpeedOfLight = 299792458 * (Meter/Second)
```

It will also work well with the `math` package from standard lib or any other
thirdparty library which is using `float64` as input or output types.
Below is an example of how it could be used with a
[vector](https://github.com/quartercastle/vector) package.

```go
Vector{10 * (Meter/Second)}.Rotate(45*Degree)
```

### Install

```sh
go get github.com/quartercastle/si
```

### Usage

```go
func MassEnergyEquivalence(energy si.Energy, mass si.Mass) bool {
  return energy == mass * math.Pow(si.SpeedOfLight, 2)
}


MassEnergyEquivalence(24.965421631578266*si.MegawattHour, 1*si.Milligram)
```


## Credits
This project is heavily inspired by [github.com/martinlindhe/unit](https://github.com/martinlindhe/unit),
which is a strongly typed unit conversion library that i have used a lot
previously. It would not exist without the knowledge i have gained by using this
library.

## License
This project is licensed under the [MIT License](https://github.com/quartercastle/si/blob/master/LICENSE).
