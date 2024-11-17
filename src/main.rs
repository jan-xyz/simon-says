mod game;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::app::Update;
use bevy::color::Color;
use bevy::prelude::BuildChildren;
use bevy::prelude::ButtonBundle;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Changed;
use bevy::prelude::Commands;
use bevy::prelude::Component;
use bevy::prelude::NodeBundle;
use bevy::prelude::Query;
use bevy::ui::AlignItems;
use bevy::ui::BackgroundColor;
use bevy::ui::FlexDirection;
use bevy::ui::Interaction;
use bevy::ui::JustifyContent;
use bevy::ui::Style;
use bevy::ui::Val;
use bevy::utils::default;
use bevy::DefaultPlugins;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugins(SimonSaysPlugin)
        .run();
}

fn ui_system(mut commands: Commands) {
    // Camera
    commands.spawn(Camera2dBundle::default());

    // Add Simon Buttons
    let c1 = Color::hsl(360. * 1 as f32 / 4 as f32, 0.95, 0.7);
    let c2 = Color::hsl(360. * 2 as f32 / 4 as f32, 0.95, 0.7);
    let c3 = Color::hsl(360. * 3 as f32 / 4 as f32, 0.95, 0.7);
    let c4 = Color::hsl(360. * 4 as f32 / 4 as f32, 0.95, 0.7);

    let mut flex_box = commands.spawn(NodeBundle {
        style: Style {
            width: Val::Percent(100.0),
            height: Val::Percent(100.0),
            flex_direction: FlexDirection::Row,
            align_items: AlignItems::Center,
            justify_content: JustifyContent::Center,
            ..default()
        },
        ..default()
    });
    flex_box.with_children(|parent| {
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c1.into(),
                ..default()
            },
            GameButton { num: 1 },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c2.into(),
                ..default()
            },
            GameButton { num: 2 },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c3.into(),
                ..default()
            },
            GameButton { num: 3 },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    ..default()
                },
                background_color: c4.into(),
                ..default()
            },
            GameButton { num: 4 },
        ));
    });
}

#[derive(Component)]
struct GameButton {
    num: i64,
}

fn button_clicked(
    mut interaction_query: Query<
        (&Interaction, &BackgroundColor, &GameButton),
        Changed<Interaction>,
    >,
) {
    for (interaction, _color, button) in &mut interaction_query {
        match interaction {
            Interaction::Pressed => {
                println!("button pressed: {:?}", button.num)
            }
            _ => {}
        }
    }
}

pub struct SimonSaysPlugin;

impl Plugin for SimonSaysPlugin {
    fn build(&self, app: &mut App) {
        let g = game::Game::new();
        g.start_game();
        app.add_systems(Startup, ui_system)
            .add_systems(Update, button_clicked);
    }
}
