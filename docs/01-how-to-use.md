---
title: How To Use
expires_at: never
tags: [goshims, volume-services]
---

# How to use

In your struct for your class add a varible referencing the interface:

```
package abroker
import (
	"code.cloudfoundry.org/goshims/ioutilshim"
	"code.cloudfoundry.org/goshims/osshim"
)

type broker struct {
	os          osshim.Os
	ioutil      ioutilshim.Ioutil
}

func New(
	os osshim.Os, ioutil ioutilshim.Ioutil,
) *broker {
	theBroker := broker{
		os:          os,
		ioutil:      ioutil,
	}
	return &theBroker
}

func (b *broker) Serialize(state interface{}) error {
	stateFile := "/tmp/abrokerstate.json"

	stateData, err := json.Marshal(state)
	if err != nil {
		return err
	}

	err = b.ioutil.WriteFile(stateFile, stateData, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

```

In the factory method to construct that class, dependency inject the right version of the implemenation.
For example, your test code would use the fakes:

```
package abroker_test
import(
	"github.com/something/abroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"code.cloudfoundry.org/goshims/osshim/os_fake"
	"code.cloudfoundry.org/goshims/ioutilshim/ioutil_fake"
)

var (
	fakeOs             *os_fake.FakeOs
	fakeIoutil         *ioutil_fake.FakeIoutil
... 
BeforeEach(func() {
	fakeOs = &os_fake.FakeOs{}
	fakeIoutil = &ioutil_fake.FakeIoutil{}
...

It("Should error if write state fails", func(){
	fakeIoutil.WriteFileReturns(errors.New("Error writing file."))
	broker = abroker.New(fakeOs, fakeIoutil)
	err := broker.Serialize(someData)
	Expect(err).To(HaveOccurred())
})
```

In your production code you would use the real implementation:

```
pacakge main
import(
	"github.com/something/abroker"
	"code.cloudfoundry.org/goshims/ioutilshim"
	"code.cloudfoundry.org/goshims/osshim"
)
...

func main() {
	broker := abroker.New(&osshim.OsShim{}, &ioutilshim.IoutilShim{})

...

```


