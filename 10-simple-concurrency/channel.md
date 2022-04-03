## Chan

### READ

Unbuffered, open

    -> Pause until something is written

Unbuffered, closed

    Return zero value (use comma ok to see if closed)

Buffered, open

    -> Pause if buffer is empty

Buffered, closed
    -> Return a remaining value in the buffer. 
    -> If the buffer is empty, return zero value (use comma ok to see if closed)

Nil
    -> Hang forever


### Write

Unbuffered, open

    -> Pause until something is read

Unbuffered, closed
    PANIC

Buffered, open

    -> Pause if buffer is full

Buffered, closed
    PANIC

Nil
    -> Hang forever


### Close

Unbuffered, open

    -> Works

Unbuffered, closed
    PANIC

Buffered, open

    -> Works, remaining values still there

Buffered, closed
    PANIC

Nil
    PANIC
