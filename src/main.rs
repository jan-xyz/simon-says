mod game;
mod logic;
mod menu;
mod state;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::prelude::AppExtStates;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Commands;
use bevy::DefaultPlugins;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugins(SimonSaysPlugin)
        .run();
}

fn startup(mut commands: Commands) {
    commands.spawn(Camera2dBundle::default());
}

pub struct SimonSaysPlugin;

impl Plugin for SimonSaysPlugin {
    fn build(&self, app: &mut App) {
        app.init_state::<state::AppState>()
            .add_sub_state::<state::GamePhase>()
            .add_plugins(menu::MenuPlugin(state::AppState::Menu))
            .add_plugins(game::InGamePlugin(state::AppState::InGame))
            .add_systems(Startup, startup);
    }
}
