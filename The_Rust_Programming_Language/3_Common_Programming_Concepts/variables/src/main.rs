// variable isn't mutable
// fn wont_compile() {
//     let x = 5;
//     println!("The value x is: {x}");
//     x = 6;
//     println!("The value x is: {x}");
// }

fn will_compile() {
    let mut x = 5;
    println!("The value x is: {x}");
    x = 6;
    println!("The value x is: {x}");
}

// constants must annotate the type
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;

fn main() {
    // wont_compile();
    will_compile();

    println!("\nConstants must not be declared during run time and cannot change {THREE_HOURS_IN_SECONDS}\n");

    let x = 5;
    let x = x + 1;
    {
        let x = x * 2;
        println!("The value of x in the innser scope is: {x}");
    }
    println!("The value of x is: {x} and shadows the original x")

    // okay
    let spaces = "   ";
    let spaces = spaces.len();
    // bad - cant change types, have to shadow/redeclare
    let mut spaces = "   ";
    spaces = spaces.len()

}

