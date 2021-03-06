# Saber

An efficient Go library with minimal CLI for dungeon mastering (DMing) tabletop games designed to be system-agnostic. Will be extended to multiple CLIs in other projects such as [for D&D 5E](https://github.com/AlexSafatli/airtable-dnd).

## Current Features

 - Character generation
 - World/region generation
 - Language generation
 - Family tree generation

## Planned Features

- Manage campaigns and store history of commands as a set of command 
  **transactions** (actions will be *journaled*, locally or remotely).
- Roll dice in different string formats.
- Create random PCs/NPCs with skeleton backgrounds, personality, and 
  strengths/weaknesses (**character generation** with *tagging*).
  - Complex characteristic and tag generation using editable tables found in files (see `tables`).
  - Will list relative attributes in a system-agnostic fashion (but use typical archetypical features such as Strength and Charisma).
- Create random locations and populations (e.g., cities, countries, regions) with placed 
  NPCs (**setting generation** with *tagging*).
- Create random rumours, adventure hooks, and roll on any customizable random 
  table including [1]:
  - Factions
  - Noble Houses
  - Terrain
- Create **random dungeons**, populate them, and generate a graphical map output.
- Create connections between locations, factions [2].
- Create histories for a setting using machine learning.
- Create adventure/episode/arc/story skeletons [3].
  - Randomly generate elements of these [4].
  - Output to PDF (LaTeX? Homebrewery-esque?).

## Package Structure

- `cmd` contains CLI structure and handlers
- `entities` stores all constructs and logic associated with entities (incl. NPCs, PCs) excl. 
  generation
- `encounters` stores all logic associated with encounters (between 
  `characters`)
- `rng` stores all logic associated with randomly choosing items
- `gen` stores all high-level random generation logic
- `stats` stores all system-specific logic
- `tables` contains random tables as text files

## Further Reading

### Related Articles

- https://www.reddit.com/r/Pathfinder_RPG/comments/2t8xwu/compilation_of_random_character_creation_tables/

### References

[1] https://www.reddit.com/r/DnD/comments/452r6r/a_massive_and_growing_resource_of_random_tables/

[2] http://redbeardsravings.blogspot.ca/2011/08/faction-connections-in-megadungeon.html

[3] https://fate-srd.com/fate-core/defining-scenarios

[4] http://geekandsundry.com/how-to-write-the-best-dd-adventures-ever/
