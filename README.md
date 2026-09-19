The server is currently under WIP, so please be patient!
I'm currently learning the SporeModAPI SDK to understand how the game works, so after it I can create a simple mod.
Due to the Spore game does not have any multiplayer support, I need to create custom gamemodes, like the cell gamemode etc in the server side, and also need to re-create it in the game too, due to I need to sync everything in server -> client.
I'm planning to add fully multiplayer support with custom game / base game modes.
Its possible? I'll figure out.
Due to I dont know C++ development, I'm looking for C++ developers to create a custom server extension to the game, with a custom network layer, after its done, I can continue it.

Heres what the server going to support:
Custom creatures
Custom gamemodes with custom missions
NPC creatures around the game (if enabled)
Sync to other players

Heres what I done so far:
I have sent the starting cell creature to the server, and I parsed it in the server side, and created a custom object.
See in: internal/models/creature/rigblock.go
Each part is tagged with '+' and new parts (IDs) sepperated with '#'

Please note that, I'm doing it in my free time, and also need to deep drive into the reverse engineered SporeModAPI SDK to understand the capabilites, etc...

The server is currently being ported from one of my custom tcp socket server: https://github.com/BXn4/CafeServerGo
The client will be ported to a custom Smartfox like extension: https://smartfoxserver.com/

Theres a test client in the server: test/client.go to test some of the client behaviour.

Status: Work In Progess!
