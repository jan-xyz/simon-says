mod game;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::app::Update;
use bevy::color::Color;
use bevy::input::ButtonInput;
use bevy::prelude::BuildChildren;
use bevy::prelude::ButtonBundle;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Commands;
use bevy::prelude::KeyCode;
use bevy::prelude::NodeBundle;
use bevy::prelude::Res;
use bevy::prelude::ResMut;
use bevy::sprite::Wireframe2dConfig;
use bevy::sprite::Wireframe2dPlugin;
use bevy::ui::AlignItems;
use bevy::ui::FlexDirection;
use bevy::ui::JustifyContent;
use bevy::ui::Style;
use bevy::ui::Val;
use bevy::utils::default;
use bevy::DefaultPlugins;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugins(SimonSaysPlugin)
        .add_plugins(Wireframe2dPlugin)
        .run();
}

fn setup(mut commands: Commands) {
    // Camera
    commands.spawn(Camera2dBundle::default());

    // Add Simon Buttons
    let c1 = Color::hsl(360. * 1 as f32 / 4 as f32, 0.95, 0.7);
    let c2 = Color::hsl(360. * 2 as f32 / 4 as f32, 0.95, 0.7);
    let c3 = Color::hsl(360. * 3 as f32 / 4 as f32, 0.95, 0.7);
    let c4 = Color::hsl(360. * 4 as f32 / 4 as f32, 0.95, 0.7);

    commands
        .spawn(NodeBundle {
            style: Style {
                width: Val::Percent(100.0),
                height: Val::Percent(100.0),
                flex_direction: FlexDirection::Row,
                align_items: AlignItems::Center,
                justify_content: JustifyContent::Center,
                ..default()
            },
            ..default()
        })
        .with_children(|parent| {
            parent.spawn(ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c1.into(),
                ..default()
            });
        })
        .with_children(|parent| {
            parent.spawn(ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c2.into(),
                ..default()
            });
        })
        .with_children(|parent| {
            parent.spawn(ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c3.into(),
                ..default()
            });
        })
        .with_children(|parent| {
            parent.spawn(ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c4.into(),
                ..default()
            });
        });
}

pub struct SimonSaysPlugin;

impl Plugin for SimonSaysPlugin {
    fn build(&self, app: &mut App) {
        let g = game::Game::new();
        g.start_game();
        app.add_systems(Startup, setup)
            .add_systems(Update, toggle_wireframe);
    }
}

fn toggle_wireframe(
    mut wireframe_config: ResMut<Wireframe2dConfig>,
    keyboard: Res<ButtonInput<KeyCode>>,
) {
    if keyboard.just_pressed(KeyCode::Space) {
        wireframe_config.global = !wireframe_config.global;
    }
}
