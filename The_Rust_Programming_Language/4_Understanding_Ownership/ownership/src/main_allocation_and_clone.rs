fn main() {
    let x = 5;
    let y = x; // known sizes and pushed to the stack
    println!("x :{x}, y: {y}");

    let s1 = String::from("hello");
    let s2 = s1; // ptr, len, capacity are on the stack, String is in the heap
    // ptr for s1 and s2 both point to the same thing
    // println!("{s1}, world!"); would have a compile error
    println!("{s2}, world!");

    let mut s = String::from("hello");
    s = String::from("ahoy");
    println!("{s}, world!"); // hello was dropped when s was reallocated

    let y1 = String::from("hello");
    let y2 = y1.clone();
    println!("y1 = {y1}, y2 = {y2}");
}
