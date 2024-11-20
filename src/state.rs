use bevy::prelude::StateSet;
use bevy::prelude::States;
use bevy::prelude::SubStates;

#[derive(Debug, Clone, Copy, Default, Eq, PartialEq, Hash, States)]
pub enum AppState {
    #[default]
    Menu,
    InGame,
}

#[derive(SubStates, Clone, PartialEq, Eq, Hash, Debug, Default)]
#[source(AppState = AppState::InGame)]
pub enum GamePhase {
    #[default]
    SimonSays,
    PlayerSays,
}
