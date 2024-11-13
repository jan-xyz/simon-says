use rand::{
    distributions::{Distribution, Standard},
    Rng,
};

#[derive(Default, Eq, PartialEq)]
pub enum Button {
    #[default]
    One,
    Two,
    Three,
    Four,
}

impl Distribution<Button> for Standard {
    fn sample<R: Rng + ?Sized>(&self, rng: &mut R) -> Button {
        // match rng.gen_range(0, 3) { // rand 0.5, 0.6, 0.7
        match rng.gen_range(0..=3) {
            // rand 0.8
            0 => Button::One,
            1 => Button::Two,
            2 => Button::Three,
            _ => Button::Four,
        }
    }
}

#[derive(Default)]
pub struct Game {
    sequence: Vec<Button>,
}

impl Game {
    pub fn new() -> Self {
        let b: Button = rand::random();
        Game { sequence: vec![b] }
    }

    pub fn start_game(mut self) {
        println!("start game!");
        let b: Button = rand::random();
        self.sequence = vec![b];
    }

    pub fn player_input(mut self, click: Button) -> bool {
        if self.sequence.last().unwrap() != &click {
            return false;
        }

        let b: Button = rand::random();
        self.sequence.push(b);
        return true;
    }
}
