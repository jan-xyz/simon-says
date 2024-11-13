use bevy::app::{App, Plugin, Startup, Update};
use bevy::color::Color;
use bevy::input::ButtonInput;
use bevy::prelude::BuildChildren;
use bevy::prelude::ButtonBundle;
use bevy::prelude::KeyCode;
use bevy::prelude::NodeBundle;
use bevy::prelude::{Camera2dBundle, Commands, Component, Res, ResMut, Resource, TextBundle};
use bevy::sprite::{Wireframe2dConfig, Wireframe2dPlugin};
use bevy::text::TextStyle;
use bevy::time::{Time, Timer, TimerMode};
use bevy::ui::AlignItems;
use bevy::ui::FlexDirection;
use bevy::ui::JustifyContent;
use bevy::ui::Style;
use bevy::ui::Val;
use bevy::utils::default;
use bevy::DefaultPlugins;

const TEXT_COLOR: Color = Color::srgb(0.5, 0.5, 1.0);
const FONT_SIZE: f32 = 40.0;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugins(HelloPlugin)
        .add_plugins(Wireframe2dPlugin)
        .run();
}

fn setup(mut commands: Commands) {
    // Camera
    commands.spawn(Camera2dBundle::default());

    // Add Text
    commands.spawn((
        HelloWorldUi,
        TextBundle::from_section(
            "Hello World",
            TextStyle {
                font_size: FONT_SIZE,
                color: TEXT_COLOR,
                ..default()
            },
        ),
    ));

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

#[derive(Component)]
struct HelloWorldUi;

#[derive(Resource)]
struct GreetTimer(Timer);

pub struct HelloPlugin;

impl Plugin for HelloPlugin {
    fn build(&self, app: &mut App) {
        app.insert_resource(GreetTimer(Timer::from_seconds(2.0, TimerMode::Repeating)))
            .add_systems(Startup, setup)
            .add_systems(Update, (game_tick, toggle_wireframe));
    }
}

fn game_tick(time: Res<Time>, mut timer: ResMut<GreetTimer>) {
    if timer.0.tick(time.delta()).just_finished() {
        println!("tick!");
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
