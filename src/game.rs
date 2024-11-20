use std::time::Duration;

use crate::logic;
use crate::state;
use crate::state::AppState;
use crate::state::GamePhase;
use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Update;
use bevy::color::Color;
use bevy::color::Luminance;
use bevy::prelude::in_state;
use bevy::prelude::BuildChildren;
use bevy::prelude::ButtonBundle;
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
use bevy::prelude::With;
use bevy::time::Time;
use bevy::time::Timer;
use bevy::time::TimerMode;
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

pub struct InGamePlugin;

impl Plugin for InGamePlugin {
    fn build(&self, app: &mut App) {
        app.add_systems(OnEnter(AppState::InGame), setup_game)
            .add_systems(OnExit(AppState::InGame), cleanup_game)
            .add_systems(OnEnter(state::GamePhase::SimonSays), schedule_simon)
            .add_systems(Update, simon_says.run_if(in_state(GamePhase::SimonSays)))
            .add_systems(Update, player_says.run_if(in_state(GamePhase::PlayerSays)))
            .add_systems(
                Update,
                (button_fade, phase_change).run_if(in_state(AppState::InGame)),
            )
            .init_resource::<logic::Game>();
    }
}
#[derive(Resource)]
pub struct GameData {
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
                width: Val::Percent(100.),
                height: Val::Percent(100.),
                align_items: AlignItems::Center,
                flex_direction: FlexDirection::Column,
                row_gap: Val::Px(10.),
                ..default()
            },
            ..default()
        })
        .with_children(|builder| {
            builder
                .spawn(NodeBundle {
                    style: Style {
                        width: Val::Percent(100.),
                        height: Val::Percent(100.),
                        flex_direction: FlexDirection::Row,
                        align_items: AlignItems::End,
                        justify_content: JustifyContent::Center,
                        column_gap: Val::Px(10.),
                        ..Default::default()
                    },
                    ..Default::default()
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
                            button: logic::Button::One,
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
                            button: logic::Button::Two,
                        },
                    ));
                });
            builder
                .spawn(NodeBundle {
                    style: Style {
                        width: Val::Percent(100.),
                        height: Val::Percent(100.),
                        flex_direction: FlexDirection::Row,
                        align_items: AlignItems::Start,
                        justify_content: JustifyContent::Center,
                        column_gap: Val::Px(10.),
                        ..Default::default()
                    },
                    ..Default::default()
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
                            border_color: c3.into(),
                            ..default()
                        },
                        GameButton {
                            button: logic::Button::Three,
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
                            button: logic::Button::Four,
                        },
                    ));
                });
        })
        .id();
    commands.insert_resource(GameData { flex_box });
}

fn cleanup_game(mut commands: Commands, game_data: Res<GameData>) {
    commands.entity(game_data.flex_box).despawn_recursive();
}

#[derive(Component)]
struct GameButton {
    button: logic::Button,
}

fn player_says(
    mut next_state: ResMut<NextState<state::AppState>>,
    mut interaction_query: Query<
        (
            &Interaction,
            &mut BackgroundColor,
            &BorderColor,
            &GameButton,
        ),
        Changed<Interaction>,
    >,
    mut g: ResMut<logic::Game>,
    mut commands: Commands,
) {
    for (interaction, mut bg_color, border_color, button) in &mut interaction_query {
        match interaction {
            Interaction::Pressed => {
                *bg_color = border_color.0.into();
                let is_correct = g.player_input(&button.button);
                if !is_correct {
                    next_state.set(state::AppState::Menu);
                } else if g.current_index == 0 {
                    commands.spawn(PhaseTimer {
                        timer: Timer::new(Duration::from_secs(1), TimerMode::Once),
                        next_phase: GamePhase::SimonSays,
                    });
                }
            }
            _ => {}
        }
    }
}

fn button_fade(mut button_query: Query<&mut BackgroundColor, With<GameButton>>, time: Res<Time>) {
    for mut bg_color in &mut button_query {
        if bg_color.0.luminance() != 0. {
            *bg_color = bg_color.0.darker(1.5 * time.delta_seconds()).into()
        }
    }
}

fn schedule_simon(mut commands: Commands, g: Res<logic::Game>) {
    for (i, button) in g.sequence.iter().enumerate() {
        commands.spawn((SimonTimer {
            // create the non-repeating fuse timer
            timer: Timer::new(Duration::from_secs_f64(1. * i as f64), TimerMode::Once),
            button: button.clone(),
        },));
    }
}

#[derive(Component)]
struct SimonTimer {
    /// track when the bomb should explode (non-repeating timer)
    timer: Timer,
    button: logic::Button,
}

fn simon_says(
    mut q: Query<(Entity, &mut SimonTimer)>,
    time: Res<Time>,
    mut commands: Commands,
    mut button_query: Query<(&mut BackgroundColor, &BorderColor, &GameButton)>,
) {
    if q.iter().len() == 0 {
        commands.spawn(PhaseTimer {
            timer: Timer::new(Duration::from_secs_f64(0.1), TimerMode::Once),
            next_phase: GamePhase::PlayerSays,
        });
    }
    for (entity, mut simon_timer) in q.iter_mut() {
        simon_timer.timer.tick(time.delta());

        if simon_timer.timer.finished() {
            for (mut bg_color, &border_color, b) in &mut button_query {
                if b.button == simon_timer.button {
                    *bg_color = border_color.0.into();
                }
            }
            commands.entity(entity).despawn();
        }
    }
}

#[derive(Component)]
struct PhaseTimer {
    timer: Timer,
    next_phase: GamePhase,
}

fn phase_change(
    mut next_phase: ResMut<NextState<state::GamePhase>>,
    mut q: Query<(Entity, &mut PhaseTimer)>,
    mut commands: Commands,
    time: Res<Time>,
) {
    for (entity, mut phase_timer) in q.iter_mut() {
        phase_timer.timer.tick(time.delta());
        if phase_timer.timer.finished() {
            commands.entity(entity).despawn();
            next_phase.set(phase_timer.next_phase.clone());
        }
    }
}
