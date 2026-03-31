fn main() {
    // let s = "hello"; // literal is not mutable
    // stripe literal on the heap
    let mut s = String::from("hello"); // String is mutable

    s.push_str(", world!");

    println!("{s}");
} // new schope, s is not valid
// popped from the stack when a string literal
// memory return to allocator when String type - Rust call `drop` automatically for us

string literals are know at compile and hard coded in the binary
need String type when when don't know the size for the allocator - memory is request at run time and returned when we're done with the String

