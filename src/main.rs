use bevy::app::{App, Plugin, Startup, Update};
use bevy::asset::Assets;
use bevy::color::Color;
use bevy::input::ButtonInput;
use bevy::prelude::KeyCode;
use bevy::prelude::Transform;
use bevy::prelude::{
    Camera2dBundle, Commands, Component, Mesh, Rectangle, Res, ResMut, Resource, TextBundle,
};
use bevy::sprite::ColorMaterial;
use bevy::sprite::{MaterialMesh2dBundle, Mesh2dHandle, Wireframe2dConfig, Wireframe2dPlugin};
use bevy::text::{TextSection, TextStyle};
use bevy::time::{Time, Timer, TimerMode};
use bevy::ui::{PositionType, Style, Val};
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

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
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
    let b1 = Mesh2dHandle(meshes.add(Rectangle::new(100.0, 100.0)));
    let c1 = Color::hsl(360. * 1 as f32 / 4 as f32, 0.95, 0.7);
    let xy1 = (-120., -120.);

    let b2 = Mesh2dHandle(meshes.add(Rectangle::new(100.0, 100.0)));
    let c2 = Color::hsl(360. * 2 as f32 / 4 as f32, 0.95, 0.7);
    let xy2 = (120., -120.);
    let b3 = Mesh2dHandle(meshes.add(Rectangle::new(100.0, 100.0)));
    let c3 = Color::hsl(360. * 3 as f32 / 4 as f32, 0.95, 0.7);
    let xy3 = (-120., 120.);
    let b4 = Mesh2dHandle(meshes.add(Rectangle::new(100.0, 100.0)));
    let c4 = Color::hsl(360. * 4 as f32 / 4 as f32, 0.95, 0.7);
    let xy4 = (120., 120.);

    commands.spawn(MaterialMesh2dBundle {
        mesh: b1,
        material: materials.add(c1),
        transform: Transform::from_xyz(xy1.0, xy1.1, 0.),
        ..default()
    });
    commands.spawn(MaterialMesh2dBundle {
        mesh: b2,
        material: materials.add(c2),
        transform: Transform::from_xyz(xy2.0, xy2.1, 0.),
        ..default()
    });
    commands.spawn(MaterialMesh2dBundle {
        mesh: b3,
        material: materials.add(c3),
        transform: Transform::from_xyz(xy3.0, xy3.1, 0.),
        ..default()
    });
    commands.spawn(MaterialMesh2dBundle {
        mesh: b4,
        material: materials.add(c4),
        transform: Transform::from_xyz(xy4.0, xy4.1, 0.),
        ..default()
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
