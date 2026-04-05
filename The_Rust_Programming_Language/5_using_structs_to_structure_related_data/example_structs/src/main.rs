fn main() {
    let widht1 = 30;
    let heigh1 = 50;

    println!(
        "The are of the rectangle is {} square pixles.",
        area(widht1, heigh1)
    );
    let rect_tuple = (30, 50);
    println!(
        "The are of the rectangle is {} square pixles.",
        area_tuple(rect_tuple)
    );
    let rectangle = Rectangle{width: widht1, height: heigh1};
    println!(
        "The are of the rectangle is {} square pixles.",
        area_struct(&rectangle) // borrow the struct instad of taking ownership, this way we can use rectangle later
    );
    println!(
        "The are of the rectangle is {} square pixles.",
        rectangle.area()
    );

    if rectangle.width() {
        println!("The rectangle has a nonzero widht; it is {}", rectangle.width)
    }
    dbg!(rectangle.width, rectangle.height);
    dbg!(rectangle);

    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));

    let num = 3;
    let square = Rectangle::square(num);
    println!("The square of {} is {:?}",num, square);
}

fn area(width: u32, height: u32) -> u32 {
    width * height
}

fn area_tuple(t: (u32, u32)) -> u32 {
    t.0 * t.1
}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32
}

fn area_struct(rectangle: &Rectangle) -> u32 {
    println!("rectangle is {rectangle:?}");
    println!("rectangle is {rectangle:#?}");
    rectangle.width * rectangle.height
}


impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn width(&self) -> bool {
        self.width > 0
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }

    fn square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    }
}

// impl Rectangle {
//     fn width(&self) -> bool {
//         self.width > 0
//     }
// }
