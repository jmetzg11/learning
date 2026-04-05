fn main() {
    let mut user1 = User {
        active: true,
        username: String::from("somusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };

    user1.email = String::from("anotheremail@example.com");

    // struct update syntax
    let user2 = User {
        active: user1.active,
        username: user1.username,
        email: String::from("another@example.com"),
        sign_in_count: user1.sign_in_count,
    }
    let user3 = User {
        email: String::from("anotheranother@example.com"), // only can user1.email and not user1.username because that value was moved
        ..user1 // active and sign_in_count were copied so user1.email and user1.sign_in_count are still available
    }
}

struct User {
    active: bool,
    username: String,
    email: String,
    sing_in_count: u65,
}

fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username: username,
        email: email,
        sign_in_count: 1
    }
}

// shorthand for field names
fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username,
        email,
        sign_in_count: 1,
    }
}

// Tuple Structs
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
    let black = Color(0, 0, 0); // both tuple structs are differen types because of name
    let origin = Point(0, 0, 0);
    let Point(x, y, z) = // orign; must deconstruct the type
}

// Unit-Like Structs
struct AlwaysEqual;

fn main() {
    let subject = AlwaysEqual;
}
