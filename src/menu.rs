use bevy::app::App;
use bevy::app::Plugin;
use bevy::app::Update;
use bevy::color::Color;
use bevy::prelude::in_state;
use bevy::prelude::BuildChildren;
use bevy::prelude::Button;
use bevy::prelude::ButtonBundle;
use bevy::prelude::Changed;
use bevy::prelude::Commands;
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
use bevy::prelude::TextBundle;
use bevy::prelude::With;
use bevy::text::TextStyle;
use bevy::ui::AlignItems;
use bevy::ui::Interaction;
use bevy::ui::JustifyContent;
use bevy::ui::Style;
use bevy::ui::Val;
use bevy::utils::default;

use crate::state::AppState;

pub struct MenuPlugin;

impl Plugin for MenuPlugin {
    fn build(&self, app: &mut App) {
        app.add_systems(Update, menu.run_if(in_state(AppState::Menu)))
            .add_systems(OnEnter(AppState::Menu), setup_menu)
            .add_systems(OnExit(AppState::Menu), cleanup_menu);
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
    mut next_state: ResMut<NextState<AppState>>,
    mut interaction_query: Query<&Interaction, (Changed<Interaction>, With<Button>)>,
) {
    for interaction in &mut interaction_query {
        match *interaction {
            Interaction::Pressed => {
                next_state.set(AppState::InGame);
            }
            Interaction::Hovered => {}
            Interaction::None => {}
        }
    }
}
