mod game;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::app::Update;
use bevy::color::Color;
use bevy::color::Luminance;
use bevy::prelude::BuildChildren;
use bevy::prelude::ButtonBundle;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Changed;
use bevy::prelude::Commands;
use bevy::prelude::Component;
use bevy::prelude::NodeBundle;
use bevy::prelude::Query;
use bevy::prelude::Res;
use bevy::prelude::ResMut;
use bevy::prelude::With;
use bevy::time::Time;
use bevy::ui::AlignItems;
use bevy::ui::BackgroundColor;
use bevy::ui::BorderColor;
use bevy::ui::FlexDirection;
use bevy::ui::Interaction;
use bevy::ui::JustifyContent;
use bevy::ui::Style;
use bevy::ui::UiRect;
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
                    border: UiRect::all(Val::Px(2.0)),
                    ..default()
                },
                background_color: Color::BLACK.into(),
                border_color: c1.into(),
                ..default()
            },
            GameButton {
                num: game::Button::One,
            },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    border: UiRect::all(Val::Px(2.0)),
                    ..default()
                },
                background_color: Color::BLACK.into(),
                border_color: c2.into(),
                ..default()
            },
            GameButton {
                num: game::Button::Two,
            },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    border: UiRect::all(Val::Px(2.0)),
                    ..default()
                },
                background_color: Color::BLACK.into(),
                border_color: c3.into(),
                ..default()
            },
            GameButton {
                num: game::Button::Three,
            },
        ));
        parent.spawn((
            ButtonBundle {
                style: Style {
                    width: Val::Px(100.),
                    height: Val::Px(100.),
                    border: UiRect::all(Val::Px(2.0)),
                    ..default()
                },
                background_color: Color::BLACK.into(),
                border_color: c4.into(),
                ..default()
            },
            GameButton {
                num: game::Button::Four,
            },
        ));
    });
}

#[derive(Component)]
struct GameButton {
    num: game::Button,
}

fn button_clicked(
    mut interaction_query: Query<
        (
            &Interaction,
            &mut BackgroundColor,
            &BorderColor,
            &GameButton,
        ),
        Changed<Interaction>,
    >,
    mut g: ResMut<game::Game>,
) {
    for (interaction, mut bg_color, border_color, button) in &mut interaction_query {
        match interaction {
            Interaction::Pressed => {
                *bg_color = border_color.0.into();
                let is_correct = g.player_input(&button.num);
                println!("{}", is_correct);
            }
            _ => {}
        }
    }
}

fn button_fade(mut button_query: Query<&mut BackgroundColor, With<GameButton>>, time: Res<Time>) {
    for mut bg_color in &mut button_query {
        *bg_color = bg_color.0.darker(1.5 * time.delta_seconds()).into()
    }
}

pub struct SimonSaysPlugin;

impl Plugin for SimonSaysPlugin {
    fn build(&self, app: &mut App) {
        app.add_systems(Startup, ui_system)
            .add_systems(Update, (button_clicked, button_fade))
            .init_resource::<game::Game>();
    }
}
