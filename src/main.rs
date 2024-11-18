mod game;

use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Startup;
use bevy::app::Update;
use bevy::color::Color;
use bevy::color::Luminance;
use bevy::prelude::in_state;
use bevy::prelude::AppExtStates;
use bevy::prelude::BuildChildren;
use bevy::prelude::Button;
use bevy::prelude::ButtonBundle;
use bevy::prelude::Camera2dBundle;
use bevy::prelude::Changed;
use bevy::prelude::Commands;
use bevy::prelude::Component;
use bevy::prelude::DespawnRecursiveExt;
use bevy::prelude::Entity;
use bevy::prelude::IntoSystemConfigs;
use bevy::prelude::NextState;
use bevy::prelude::NodeBundle;
use bevy::prelude::OnEnter;
use bevy::prelude::OnExit;
use bevy::prelude::Query;
use bevy::prelude::Res;
use bevy::prelude::ResMut;
use bevy::prelude::Resource;
use bevy::prelude::StateSet;
use bevy::prelude::States;
use bevy::prelude::SubStates;
use bevy::prelude::TextBundle;
use bevy::prelude::With;
use bevy::text::TextStyle;
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

#[derive(Debug, Clone, Copy, Default, Eq, PartialEq, Hash, States)]
enum GameState {
    #[default]
    Menu,
    InGame,
}

#[derive(SubStates, Clone, PartialEq, Eq, Hash, Debug, Default)]
#[source(GameState = GameState::InGame)]
enum GamePhase {
    #[default]
    SimonSays,
    PlayerSays,
}

fn startup(mut commands: Commands) {
    commands.spawn(Camera2dBundle::default());
}

#[derive(Resource)]
struct GameData {
    flex_box: Entity,
}

fn setup_game(mut commands: Commands) {
    // Add Simon Buttons
    let c1 = Color::hsl(360. * 1 as f32 / 4 as f32, 0.95, 0.7);
    let c2 = Color::hsl(360. * 2 as f32 / 4 as f32, 0.95, 0.7);
    let c3 = Color::hsl(360. * 3 as f32 / 4 as f32, 0.95, 0.7);
    let c4 = Color::hsl(360. * 4 as f32 / 4 as f32, 0.95, 0.7);

    let flex_box = commands
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
        })
        .id();
    commands.insert_resource(GameData { flex_box });
}
fn cleanup_game(mut commands: Commands, game_data: Res<GameData>) {
    commands.entity(game_data.flex_box).despawn_recursive();
}

#[derive(Component)]
struct GameButton {
    num: game::Button,
}

fn button_clicked(
    mut next_state: ResMut<NextState<GameState>>,
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
                if !is_correct {
                    next_state.set(GameState::Menu);
                }
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

#[derive(Resource)]
struct MenuData {
    button_entity: Entity,
}

fn setup_menu(mut commands: Commands) {
    let button_entity = commands
        .spawn(NodeBundle {
            style: Style {
                // center button
                width: Val::Percent(100.),
                height: Val::Percent(100.),
                justify_content: JustifyContent::Center,
                align_items: AlignItems::Center,
                ..default()
            },
            ..default()
        })
        .with_children(|parent| {
            parent
                .spawn(ButtonBundle {
                    style: Style {
                        width: Val::Px(150.),
                        height: Val::Px(65.),
                        // horizontally center child text
                        justify_content: JustifyContent::Center,
                        // vertically center child text
                        align_items: AlignItems::Center,
                        ..default()
                    },
                    background_color: Color::hsl(360. * 1 as f32 / 5 as f32, 0.95, 0.7).into(),
                    ..default()
                })
                .with_children(|parent| {
                    parent.spawn(TextBundle::from_section(
                        "Play",
                        TextStyle {
                            font_size: 40.0,
                            color: Color::BLACK.into(),
                            ..default()
                        },
                    ));
                });
        })
        .id();
    commands.insert_resource(MenuData { button_entity });
}

fn cleanup_menu(mut commands: Commands, menu_data: Res<MenuData>) {
    commands.entity(menu_data.button_entity).despawn_recursive();
}

fn menu(
    mut next_state: ResMut<NextState<GameState>>,
    mut interaction_query: Query<&Interaction, (Changed<Interaction>, With<Button>)>,
) {
    for interaction in &mut interaction_query {
        match *interaction {
            Interaction::Pressed => {
                next_state.set(GameState::InGame);
            }
            Interaction::Hovered => {}
            Interaction::None => {}
        }
    }
}

pub struct SimonSaysPlugin;

impl Plugin for SimonSaysPlugin {
    fn build(&self, app: &mut App) {
        app.init_state::<GameState>()
            .add_sub_state::<GamePhase>()
            .add_systems(Startup, startup)
            .add_systems(Update, menu.run_if(in_state(GameState::Menu)))
            .add_systems(OnEnter(GameState::InGame), setup_game)
            .add_systems(OnExit(GameState::InGame), cleanup_game)
            .add_systems(
                Update,
                button_clicked.run_if(in_state(GamePhase::PlayerSays)),
            )
            .add_systems(
                Update,
                (button_clicked, button_fade).run_if(in_state(GameState::InGame)),
            )
            .add_systems(OnEnter(GameState::Menu), setup_menu)
            .add_systems(OnExit(GameState::Menu), cleanup_menu)
            .init_resource::<game::Game>();
    }
}
