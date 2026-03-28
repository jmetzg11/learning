fn main() {
    // must declare type because a string can become many numeric types
    let guess: u32 = "42".parse().expect("Not a number!");
}
// INTEGER TYPES
// 8-bit	i8	u8
// 16-bit	i16	u16
// 32-bit	i32	u32
// 64-bit	i64	u64
// 128-bit	i128	u128
// Architecture-dependent	isize	usize

1_000_000 == 1000000

Number literals	Example
Decimal	98_222
Hex	0xff
Octal	0o77
Binary	0b1111_0000
Byte (u8 only)	b'A'

all floating points are signed
let x = 2.0; // f64
let y: f32 = 3.0;


// The Character Type
fn main() {
    let c = 'z';
    let z: char = 'ℤ'; // with explicit type annotation
    let heart_eyed_cat = '😻';
}

// The Tuple Type
fn main() {
    // type annotations are optional
    let tup: (i32, f64, u8) = (500, 6.4, 1);
}
// deconstructing a tuple
fn main() {
    let tup = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {y}"); -> 6.4

    let example = tup.0 -> example is 500
}

// The Array type
// fixed length and same type
fn main() {
    let a = [1, 2, 3, 4, 5];
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    let a = [3; 5]; -> let a = [3, 3, 3, 3, 3];
}
// Array Element Access
fn main() {
    let a = [1, 2, 3, 4, 5];

    let first = a[0];
    let second = a[1];
}
