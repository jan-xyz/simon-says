mod game;
mod logic;
mod menu;
mod state;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::app::Update;
use bevy::prelude::in_state;
use bevy::prelude::AppExtStates;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Commands;
use bevy::prelude::IntoSystemConfigs;
use bevy::prelude::OnEnter;
use bevy::prelude::OnExit;
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
        app.init_state::<state::GameState>()
            .add_sub_state::<state::GamePhase>()
            .add_systems(Startup, startup)
            .add_systems(Update, menu::menu.run_if(in_state(state::GameState::Menu)))
            .add_systems(OnEnter(state::GameState::InGame), game::setup_game)
            .add_systems(OnExit(state::GameState::InGame), game::cleanup_game)
            .add_systems(
                Update,
                game::button_clicked.run_if(in_state(state::GamePhase::PlayerSays)),
            )
            .add_systems(
                Update,
                (game::button_clicked, game::button_fade)
                    .run_if(in_state(state::GameState::InGame)),
            )
            .add_systems(OnEnter(state::GameState::Menu), menu::setup_menu)
            .add_systems(OnExit(state::GameState::Menu), menu::cleanup_menu)
            .init_resource::<logic::Game>();
    }
}
