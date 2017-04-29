package config_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/codemodus/config"
)

type testConf struct {
	// Not needed because InitPost() is defined.
	// *config.Config

	SampleText string
	TestText   string

	*embeddedConf
}

func (tc *testConf) InitPost() error {
	emConf := &embeddedConf{}
	if err := config.Init(emConf, "testdata/config2.json"); err != nil {
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

type badConf struct {
	SampleText string
	TestText   string

	*embeddedConf
}

func (bc *badConf) InitPost() error {
	return errors.New("fake an error")
}

func Example() {
	myConf := &struct {
		*config.Config

		SampleText string
		TestText   string
	}{}

	if err := config.Init(myConf, "testdata/config.json"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(myConf.SampleText)
	fmt.Println(myConf.TestText)

	// Output:
	// sampled
	// tested
}

func TestConfig(t *testing.T) {
	myConf := &struct {
		*config.Config

		SampleText string
		TestText   string
	}{}

	if err := config.Init(myConf, "testdata/config.json"); err != nil {
		t.Fatal(err)
	}

	if myConf.SampleText != "sampled" {
		t.Error("Sample text is wrong. (config)")
	}
	if myConf.TestText != "tested" {
		t.Error("Test text is wrong. (config)")
	}
}

func TestBadData(t *testing.T) {
	conf := &struct {
		*config.Config
	}{}

	if err := config.Init(conf, ""); err == nil {
		t.Error("Should have returned error - no config file exists")
	}

	if err := config.Init(conf, "testdata/bad.json"); err == nil {
		t.Error("Should have returned error - no data in config file")
	}

	badConf := &badConf{}
	if err := config.Init(badConf, "testdata/config.json"); err == nil {
		t.Error("Should have returned error - InitPost returns an error")
	}
}

func TestEmbedded(t *testing.T) {
	conf := &testConf{}
	if err := config.Init(conf, "testdata/config.json"); err != nil {
		t.Fatal(err)
	}

	if conf.SampleText != "sampled" {
		t.Error("Sample text is wrong. (embedded)")
	}
	if conf.TestText != "tested" {
		t.Error("Test text is wrong. (embedded)")
	}
	if conf.embeddedConf.SampleText2 != "sampled2" {
		t.Error("Sample text 2 is wrong. (embedded)")
	}
	if conf.embeddedConf.TestText2 != "tested2" {
		t.Error("Test text 2 is wrong. (embedded)")
	}
}
