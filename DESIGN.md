# Design

## Premise

Saber will be designed to be an efficient CLI for DMing tabletop games and 
initially be system-agnostic but eventually support multiple systems.

## Features

- Manage campaigns and store history of commands as a set of command 
  transactions (journal).
- Roll dice in different string formats.
- Create random PCs/NPCs with skeleton backgrounds, personality, and 
  strengths/weaknesses.
- Create random locations and populations (e.g., cities, countries, regions) of 
  NPCs.
- Create random rumours, adventure hooks, and roll on any customizable random 
  table including [1]:
  - Factions
  - Noble Houses
  - Terrain
- Create random dungeons, populate them, and generate a graphical map output.
- Create connections between factions [2].
- Create histories for a setting using machine learning.
- Create adventure/episode/arc/story skeletons [3].
  - Randomly generate elements of these [4].
  - Output to PDF (LaTeX? Homebrewery-esque?).

## Package Structure

- `entities` stores all logic associated with entities (incl. NPCs, PCs) excl. 
  generation
- `encounters` stores all logic associated with encounters (between 
  `characters`)
- `rng` stores all random generation logic of different entities
- `stats` stores all system-specific logic
- `tools` stores all logic associated with utilities such as dice

[1] https://www.reddit.com/r/DnD/comments/452r6r/a_massive_and_growing_resource_of_random_tables/
[2] http://redbeardsravings.blogspot.ca/2011/08/faction-connections-in-megadungeon.html
[3] https://fate-srd.com/fate-core/defining-scenarios
[4] http://geekandsundry.com/how-to-write-the-best-dd-adventures-ever/