## RingBuffer go

Ring/Circular buffer implementation in golang.

### Usage
```
import "github.com/NishanthSpShetty/goring"

buf, err := goring.New(10)

if err !=nil {
// failed to create a buffer for some reason
}

```

### Write and read from the buffer
```
//insert the value
err := buf.Offer(10)
if err != nil {
  //full may be?
}

//read from the value
val, err := buf.Poll()
if err != nil {
  // empty ?
}
```


### Other implementations
Rust implemetation [here](https://github.com/NishanthSpShetty/ringbuffer)

Zig implementaton  [here](https://github.com/NishanthSpShetty/zigring)
