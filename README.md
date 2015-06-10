# config

    go get "github.com/codemodus/config"

Package config provides an interface and initialization function for handling 
configuration values stored as JSON.  The JSON structure is defined by a user 
configurable struct which implements Configurator.  Nested Configurator 
instances by calling Init within the parent's InitPost.

## Usage

```go
func Init(c Configurator, file string) (err error)
type Config
    func (c *Config) InitPost() error
type Configurator
```

### Setup

```go
import (
    "fmt"

    "github.com/codemodus/config"
)

func main() {
    myConf := &struct {
        *config.Config

        SampleText string
        TestText   string
    }{}

    if err := config.Init(myConf, "test_dir/config.json"); err != nil {
        fmt.Println(err)
    }

    fmt.Println(myConf.SampleText) // "sampled"
    fmt.Println(myConf.TestText)   // "tested"
}
```

### JSON Example

```json
{
  "SampleText": "sampled",
  "TestText": "tested"
}
```

### Nested Configurator Objects

```go
type testConf struct {
    // Not needed because InitPost() is defined.
    // *config.Config
    
    SampleText string
    TestText   string

    *embeddedConf
}

func (tc *testConf) InitPost() error {
    emConf := &embeddedConf{}
    if err := config.Init(emConf, "test_dir/config2.json"); err != nil {
        return err
    }
    tc.embeddedConf = emConf
    return nil
}

type embeddedConf struct {
    *config.Config

    SampleText2 string
    TestText2   string
}
```

## More Info

N/A

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/config)

## Benchmarks

N/A
