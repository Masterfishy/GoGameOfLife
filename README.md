# Gopher Life

This is a Golang implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) following [the great tutorial](https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl) by Kyle Banks.

Along with being an implementation of the Game of Life, it is also my first implementation of specific game architecture and game engine design. After reading Richard Lord's [article on Entity Component System (ECS) architectures for games](https://www.richardlord.net/blog/ecs/what-is-an-entity-framework.html), I modified Kyle Banks' implementation to follow the ECS design.

## Motivation

### The Problem

I have been making games for the past 5+ years. My programming background primarily being software development and university classes, I always approached game programming from a Object Oriented lense. Encapsulation, inheritance, polymorphism, our lovely world of OOP. For simple games made in 48 hours, this worked.

As my ambitions and games grew, I found myself running into the same issues related to horrible webs of dependencies. To place my player I needed my map manager. For my map manager I need the game manager. For the game manager I need my level manager. Just to experiment with my player I had to build my entire game. The webs were thick and my patience for messy code was running out. At this point, I realized it was time to learn a bit more about game programming.

### A Solution

I watched a fantastic [Game Developers Conference (GDC) talk](https://youtu.be/raQ3iHhE_Kk) on using Scriptable Objects in Unity. Scriptable objects, at the most fundamental, are serialized data classes. Exploring this world of scriptable objects, I discovered another paradigm of coding hidden in the shadow of OOP (from me): ECS.

Entity Component System (ECS) is an architecture composed primarily of the three parts of the name: entities, components, and systems. The idea is to create components that hold data and systems that act and transform this data. Now, this breaks many of the rules of OOP, like encapsulation, since data is not strictly associated with a class and typically accessible by multiple classes. This "liberation" of data, however, is extremely valuable in game programming.

When multiple game systems need to know the player's data, it is much more effective to serialize the data and store it in memory, than to have it only exist at runtime when the player exists in the game space. Now the player systems and other game systems are decoupled. No longer are there dependency webs, but instead, data. We love data!

### Game Architecture and Engines

Reading the [article by Richard Lord](https://www.richardlord.net/blog/ecs/what-is-an-entity-framework.html) I mentioned above, I found an opportunity to explore my interest in game engine design. While rather simple, ECS can be used to architecture a game engine where data is stored as components, systems operate and update the components, and entities group components.

With a simple game like Conway's Game of Life, I set off to try my hand at ECS.

### Why Go

Because Go is python with pointers, and I think that's radical.

## The Graphics

In the graphics folder, you can find code that communicates with OpenGL to render a window on Windows machines.

The file `shader.go` provides simple shaders and a function for compiling shaders with OpenGL.

The file `vao.go` provides definitions for vector array objects so that OpenGL can render simple points.

## The Engine Architecture

The engine is broken down into multiple levels: an engine, entities, nodes, systems, and components.

### Engine

The `engine.go` file contains the definition for the engine. The engine in the ECS model keeps track of the time, the entities, the systems, and the nodes. It updates the collection of active systems using the time. The engine is the main loop.

### Entities

The `entity.go` file contains the definition for the entity struct. In ECS, entities aren't actually necessary and can be implemented in many ways. Entities are primarily used to associate components with a specific thing in the game; like an enemy, player, level.

In the _Gopher Life_ engine, Entities hold components. This allows the engine to know which components are associated with an entity, it an entity is added or removed.

### Nodes

The `nodes.go` file contains the definition for the different node structs. Similar to the entities, nodes aren't necessary to ECS and can take different forms. The key here, is to group certain components with each other for use by a system. For example: both a rendering system and a movement system need position, so a render node and a movement node would hold pointers to a position component. But both systems don't necessarily need the same data, so nodes are system specific.

The _Gopher Life_ engine has two types of nodes: `RenderNodes` and `LivingNodes`.

The `RenderNodes` hold pointers to components necessary for rendering an object to the screen. The `RenderSystem` reads these nodes and the data associated to draw.

The `LivingNodes` holder pointers to components necessary for tracking the living state of a cell. The `LivingSystem` uses these nodes to update cells based on the [Rules of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules).

### Systems

The `systems.go` file contains the definitions for the two systems of the game: `RenderSystem` and the `LivingSystem`. Systems in ECS read and modify data to simulate our game. Calling the update function on the system moves the "simulation" a time step.

The `RenderSystem` creates a vector array object from data provided by the `RenderNodes` and draws them to the screen. It draws the living cells.

The `LivingSystem` determines the next state of all the cells based on the current state of the simulation contained in the `LivingNodes` and the [Rules of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules).

### Components

The `components.go` file contains all the data needed that a cell or game state or system might need.

The `PositionComponent` contains an x and y coordinate and a rotation.

The `DisplayComponent` has an array of points that are the vertices of a shape. It also contains an x, y, and rotation specific to rendering.

The `LivingComponent` has two bools determining the living state of a cell.
