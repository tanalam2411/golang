
###### Difference between new() and regular allocation (&)

[ref](https://stackoverflow.com/questions/13244947/is-there-a-difference-between-new-and-regular-allocation)

e.g:

```go

v := &Vector{}

// VS

v := new(Vector)
``` 

- `new()` is a function that allocates memory, and zeros it out; that is every field (and the entire piece of memory for the structure) will be set to 0s.

---

##### new() vs make()
[ref](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)

- `new(T)` - Allocates memory and sets it to zero value for type `T`, i.e. `0` for int, `""` for string and `nil` for referenced types(slice, map, chan).
- Referenced types are just pointers to some underlying data structures, which won't be created by `new(T)`.
- So in case of slice `s := new(int[])`, the underlying array won't be created, thus `s` returns a pointer to nothing (nil).

- `make(T)` - Allocates memory for referenced data types (slice, map, chan), plus initializes their underlying data structures. 
- So in case of slice, the underlying `array` will be created with specified length and capacity.

---

- `new(T)` - Returns a pointer to type `T`  and a value of type `*T`. 
  - `new(T)` is equivalent to `&T{}`.
- `make(T)` - Returns an initialized value of type T. 
  - `make`'s return type of the same as the type of its argument, not a pointer of it.

