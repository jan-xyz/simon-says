use bevy::prelude::Resource;
use rand::{
    distributions::{Distribution, Standard},
    Rng,
};

#[derive(Debug, Default, Eq, PartialEq)]
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

#[derive(Resource)]
pub struct Game {
    pub sequence: Vec<Button>,
    current_index: usize,
}

impl Game {
    pub fn player_input(&mut self, click: &Button) -> bool {
        if &self.sequence[self.current_index] != click {
            self.current_index = 0;
            let b: Button = rand::random();
            self.sequence = vec![b];
            println!("{:?}", self.sequence);
            return false;
        }
        self.current_index += 1;

        if self.sequence.len() == self.current_index {
            let b: Button = rand::random();
            self.sequence.push(b);
            self.current_index = 0;
            println!("{:?}", self.sequence);
        }
        return true;
    }
}

// custom implementation for unusual values
impl Default for Game {
    fn default() -> Self {
        let b: Button = rand::random();
        let sequence = vec![b];
        println!("{:?}", sequence);
        let current_index = 0;
        Game {
            sequence,
            current_index,
        }
    }
}
