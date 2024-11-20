use bevy::prelude::StateSet;
use bevy::prelude::States;
use bevy::prelude::SubStates;

#[derive(Debug, Clone, Copy, Default, Eq, PartialEq, Hash, States)]
pub enum GameState {
    #[default]
    Menu,
    InGame,
}

#[derive(SubStates, Clone, PartialEq, Eq, Hash, Debug, Default)]
#[source(GameState = GameState::InGame)]
pub enum GamePhase {
    #[default]
    SimonSays,
    PlayerSays,
}
