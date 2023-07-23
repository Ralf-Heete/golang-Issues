# Issue: Using Pointer Receiver Methods in Go Interface
## Problem Description

In Go, we encountered an issue when defining an interface that includes methods with pointer receivers. The error message stated that the type implementing the interface did not satisfy it due to the pointer receiver methods.

Example:

    package main

    import "fmt"

    type Foo interface {
        Fizz()
    }

    type Bar struct{}

    func (b *Bar) Fizz() {
        fmt.Println("fizz")
    }

    func Fizzy(foo Foo) {
        foo.Fizz()
    }

    func main() {
      b := Bar{}
      Fizzy(b)
    }
 
Running the above code results in the error:
cannot use b (type Bar) as type Foo in argument to Fizzy:
Bar does not implement Foo (Fizz method has pointer receiver)
    
## The Solution

The error occurs because a pointer to a type is considered a different type than the actual type. In the example, *Bar is a different type than Bar. The Fizz method is defined with a pointer receiver on *Bar, not on Bar, which is why only *Bar satisfies the Foo interface.

To fix this issue, we need to use a pointer to the actual type when implementing the interface.

## Example:
  
    package main

    import "fmt"
    
    type Foo interface {
        Fizz()
    }
    
    type Bar struct{}

    func (b *Bar) Fizz() {
       fmt.Println("fizz")
    }

    func Fizzy(foo Foo) {
       foo.Fizz()
    }

    func main() {
       b := &Bar{} // Use a pointer to Bar
       Fizzy(b)
    }
This time, the code will compile and run without errors, printing "fizz" as expected.

## Workaround

To work around this limitation, we can create wrapper functions for methods with pointer receivers and call those functions in the interface implementation.

## Example:
    package main

    import "fmt"

    type value interface{}
    type dict map[string]value
    type list []value

    func (recv *dict) SetItem(k value, v value) bool {
        key, ok := k.(string)
        if ok {
            (*recv)[key] = v
            return true
        }
        return false
    }
    
    func (recv *list) SetItem(k value, v value) bool {
        i, ok := k.(int)
        if !ok || len(*recv) <= i {
           return false
        }

        (*recv)[i] = v
        return true
    }
       
    // Wrapper function that calls SetItem and returns the result.
    func SetItemWrapper(d value, k value, v value) bool {
        switch t := d.(type) {
        case *dict:
            return t.SetItem(k, v)
        case *list:
            return t.SetItem(k, v)
        default:
            return false
        }
    }
        
    func main() {
        d := dict{}
        l := list{"1", "2"}
        v := value(&d)

        fmt.Printf("dict:%v\n", SetItemWrapper(v, "foo", "bar"))
        fmt.Printf("d:'%v'\n", d)
        
        v = value(&l)
        fmt.Printf("list:%v\n", SetItemWrapper(v, 1, "bar"))
        fmt.Printf("l:'%v'\n", l)
    }
    
In this workaround, we use a value interface to handle both pointer and non-pointer types. The SetItemWrapper function acts as a wrapper for SetItem, allowing us to call the method on both the pointer and non-pointer types in the interface.

Feel free to adjust the template to include more details or explanations specific to your use case. I hope this helps you create the Readme.md for your issue!
